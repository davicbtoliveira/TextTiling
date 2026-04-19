package tiling

import (
	"errors"

	"github.com/texttiling/texttiling/pkg/config"
	"github.com/texttiling/texttiling/pkg/text"
)

type TextTiling struct {
	opts      *config.Options
	stopwords map[string]struct{}
	tokenizer *Tokenizer
	similarity *SimilarityComputer
	boundary *BoundaryDetector
}

func New(opts ...config.Option) (*TextTiling, error) {
	cfg := config.Default()
	for _, opt := range opts {
		opt(cfg)
	}

	stopwords, err := text.LoadStopwords(cfg.StopwordsFile, cfg.Language)
	if err != nil {
		return nil, err
	}

	tk := NewTokenizer(stopwords, cfg.W)
	sc := NewSimilarityComputer(cfg.K, cfg.SimilarityMethod)
	bd := NewBoundaryDetector(cfg)

	return &TextTiling{
		opts:      cfg,
		stopwords: stopwords,
		tokenizer: tk,
		similarity: sc,
		boundary: bd,
	}, nil
}

func (t *TextTiling) Segment(text string) ([]Segment, error) {
	if text == "" {
		return []Segment{}, nil
	}

	pseudosentences := t.tokenizer.Pseudosentences(text)
	if len(pseudosentences) < 2 {
		return []Segment{{Index: 0, Text: text, StartPos: 0, EndPos: len(text), Size: len(t.tokenizer.Tokenize(text))}}, nil
	}

	scores := t.similarity.ComputeScores(pseudosentences)
	smoothed := t.boundary.Smooth(scores)
	depthScores := t.boundary.ComputeDepth(smoothed)
	boundaries := t.boundary.DetectBoundaries(depthScores)

	return t.createSegments(text, boundaries, pseudosentences), nil
}

func (t *TextTiling) SegmentWithDebug(text string) (*Result, error) {
	if text == "" {
		return &Result{
			Segments:        []Segment{},
			Pseudosentences: []Pseudosentence{},
			RawScores:      []float64{},
			SmoothedScores:  []float64{},
			DepthScores:   []float64{},
			Boundaries:    []bool{},
		}, nil
	}

	pseudosentences := t.tokenizer.Pseudosentences(text)
	if len(pseudosentences) < 2 {
		return &Result{
			Segments: []Segment{{Index: 0, Text: text, StartPos: 0, EndPos: len(text), Size: len(t.tokenizer.Tokenize(text))}},
			Pseudosentences: pseudosentences,
			RawScores:      []float64{},
			SmoothedScores:  []float64{},
			DepthScores:   []float64{},
			Boundaries:    []bool{false},
		}, nil
	}

	rawScores := t.similarity.ComputeScores(pseudosentences)
	smoothed := t.boundary.Smooth(rawScores)
	depthScores := t.boundary.ComputeDepth(smoothed)
	boundaries := t.boundary.DetectBoundaries(depthScores)
	segments := t.createSegments(text, boundaries, pseudosentences)

	return &Result{
		Segments:        segments,
		Pseudosentences: pseudosentences,
		RawScores:      rawScores,
		SmoothedScores:  smoothed,
		DepthScores:   depthScores,
		Boundaries:    boundaries,
	}, nil
}

func (t *TextTiling) SegmentWithGaps(text string) ([]SegmentWithGap, error) {
	segments, err := t.Segment(text)
	if err != nil {
		return nil, err
	}

	result := make([]SegmentWithGap, len(segments))
	for i, seg := range segments {
		result[i] = SegmentWithGap{
			Segment:   seg,
			GapBefore: i > 0,
			GapAfter:  i < len(segments)-1,
		}
	}

	return result, nil
}

func (t *TextTiling) createSegments(text string, boundaries []bool, pseudosentences []Pseudosentence) []Segment {
	if len(boundaries) == 0 {
		return []Segment{{Index: 0, Text: text, StartPos: 0, EndPos: len(text), Size: len(t.tokenizer.Tokenize(text))}}
	}

	var segmentIndices []int
	segmentIndices = append(segmentIndices, 0)

	for i, isBoundary := range boundaries {
		if isBoundary {
			segmentIndices = append(segmentIndices, i+1)
		}
	}
	segmentIndices = append(segmentIndices, len(pseudosentences))

	segments := make([]Segment, 0, len(segmentIndices)-1)
	for i := 0; i < len(segmentIndices)-1; i++ {
		startIdx := segmentIndices[i]
		endIdx := segmentIndices[i+1]

		if startIdx >= len(pseudosentences) {
			continue
		}
		if endIdx > len(pseudosentences) {
			endIdx = len(pseudosentences)
		}

		startPos := 0
		if startIdx > 0 && startIdx < len(pseudosentences) {
			startPos = pseudosentences[startIdx].Positions[0]
		}

		endPos := len(text)
		if endIdx > 0 && endIdx <= len(pseudosentences) && len(pseudosentences[endIdx-1].Positions) > 0 {
			lastPos := pseudosentences[endIdx-1].Positions[len(pseudosentences[endIdx-1].Positions)-1]
			for _, token := range pseudosentences[endIdx-1].Tokens {
				if len(token) > 0 {
					endPos = lastPos + len(token)
					break
				}
			}
			if endPos < startPos {
				endPos = len(text)
			}
		}

		if startPos >= len(text) {
			startPos = 0
		}
		if endPos > len(text) {
			endPos = len(text)
		}
		if endPos < startPos {
			endPos = startPos
		}

		segmentText := text[startPos:endPos]
		if i < len(segmentIndices)-1 && endIdx < len(pseudosentences) {
			lastPS := pseudosentences[endIdx-1]
			lastTokenEnd := 0
			if len(lastPS.Positions) > 0 && len(lastPS.Tokens) > 0 {
				lastTokenEnd = lastPS.Positions[len(lastPS.Positions)-1] + len(lastPS.Tokens[len(lastPS.Tokens)-1])
				if lastTokenEnd > startPos {
					extra := text[lastTokenEnd:]
					lineEnd := len(extra)
					for j, r := range extra {
						if r == '\n' {
							lineEnd = j
							break
						}
					}
					if lineEnd < len(extra) {
						segmentText = text[startPos:lastTokenEnd+lineEnd]
					}
				}
			}
		}

		totalTokens := 0
		for j := startIdx; j < endIdx && j < len(pseudosentences); j++ {
			totalTokens += len(pseudosentences[j].Tokens)
		}

		segments = append(segments, Segment{
			Index:     i,
			Text:     segmentText,
			StartPos: startPos,
			EndPos:   startPos + len(segmentText),
			Size:     totalTokens,
		})
	}

	return segments
}

var ErrEmptyText = errors.New("empty text provided")