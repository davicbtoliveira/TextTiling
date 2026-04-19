package tiling

import (
	"math"

	"github.com/texttiling/texttiling/pkg/config"
)

type SimilarityComputer struct {
	k     int
	method config.SimilarityMethod
}

func NewSimilarityComputer(k int, method config.SimilarityMethod) *SimilarityComputer {
	return &SimilarityComputer{
		k:     k,
		method: method,
	}
}

func (sc *SimilarityComputer) ComputeScores(pseudosentences []Pseudosentence) []float64 {
	if len(pseudosentences) < 2 {
		return nil
	}

	numGaps := len(pseudosentences) - 1
	scores := make([]float64, numGaps)

	for gap := 0; gap < numGaps; gap++ {
		scores[gap] = sc.computeGapScore(pseudosentences, gap)
	}

	return scores
}

func (sc *SimilarityComputer) computeGapScore(ps []Pseudosentence, gap int) float64 {
	switch sc.method {
	case config.VocabularyIntroduction:
		return sc.vocabularyIntroductionScore(ps, gap)
	default:
		return sc.blockComparisonScore(ps, gap)
	}
}

func (sc *SimilarityComputer) blockComparisonScore(ps []Pseudosentence, gap int) float64 {
	windowSize := sc.k
	if gap < sc.k-1 {
		windowSize = gap + 1
	} else if gap > len(ps)-sc.k {
		windowSize = len(ps) - gap
	}

	leftTokens := collectTokens(ps, gap-windowSize+1, gap+1)
	rightTokens := collectTokens(ps, gap+1, gap+windowSize+1)

	dotProduct := 0.0
	normLeft := 0.0
	normRight := 0.0

	for term, freqLeft := range leftTokens {
		freqRight := rightTokens[term]
		dotProduct += float64(freqLeft) * float64(freqRight)
		normLeft += float64(freqLeft) * float64(freqLeft)
		normRight += float64(freqRight) * float64(freqRight)
	}

	if normLeft == 0 || normRight == 0 {
		return 0.0
	}

	return dotProduct / (math.Sqrt(normLeft) * math.Sqrt(normRight))
}

func (sc *SimilarityComputer) vocabularyIntroductionScore(ps []Pseudosentence, gap int) float64 {
	leftTokens := collectTokenSet(ps, 0, gap+1)
	rightTokens := collectTokenSet(ps, gap+1, len(ps))

	newTerms := 0
	rightUnique := 0
	for term := range rightTokens {
		rightUnique++
		if !leftTokens[term] {
			newTerms++
		}
	}

	w := len(ps[0].Tokens)
	if w == 0 {
		w = 1
	}

	return float64(newTerms*2) / float64(2*w)
}

func collectTokens(ps []Pseudosentence, start int, end int) map[string]int {
	result := make(map[string]int)
	for i := start; i < end && i < len(ps); i++ {
		if i < 0 {
			continue
		}
		for _, token := range ps[i].Tokens {
			result[token]++
		}
	}
	return result
}

func collectTokenSet(ps []Pseudosentence, start int, end int) map[string]bool {
	result := make(map[string]bool)
	for i := start; i < end && i < len(ps); i++ {
		if i < 0 {
			continue
		}
		for _, token := range ps[i].Tokens {
			result[token] = true
		}
	}
	return result
}