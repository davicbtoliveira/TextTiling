package tiling

import (
	"regexp"
	"strings"
)

type Tokenizer struct {
	stopwords map[string]struct{}
	w        int
}

func NewTokenizer(stopwords map[string]struct{}, w int) *Tokenizer {
	sw := make(map[string]struct{})
	for word := range stopwords {
		sw[strings.ToLower(word)] = struct{}{}
	}
	stopwords = sw
	return &Tokenizer{
		stopwords: stopwords,
		w:        w,
	}
}

func (tz *Tokenizer) Tokenize(text string) []string {
	re := regexp.MustCompile(`[\p{L}\p{N}]+`)
	matches := re.FindAllString(strings.ToLower(text), -1)

	tokens := make([]string, 0, len(matches))
	for _, match := range matches {
		if _, isStopword := tz.stopwords[match]; !isStopword {
			tokens = append(tokens, match)
		}
	}
	return tokens
}

func (tz *Tokenizer) Pseudosentences(text string) []Pseudosentence {
	tokens := tz.Tokenize(text)
	if len(tokens) == 0 {
		return nil
	}

	positions := tz.tokenPositions(text)

	numPS := (len(tokens) + tz.w - 1) / tz.w
	result := make([]Pseudosentence, numPS)

	for i := 0; i < numPS; i++ {
		start := i * tz.w
		end := start + tz.w
		if end > len(tokens) {
			end = len(tokens)
		}

		psTokens := tokens[start:end]
		psPositions := positions[start:end]

		result[i] = Pseudosentence{
			Index:      i,
			Tokens:    psTokens,
			Positions: psPositions,
		}
	}

	return result
}

func (tz *Tokenizer) tokenPositions(text string) []int {
	re := regexp.MustCompile(`[\p{L}\p{N}]+`)
	matches := re.FindAllStringIndex(strings.ToLower(text), -1)

	positions := make([]int, len(matches))
	for i, match := range matches {
		positions[i] = match[0]
	}
	return positions
}