package embed

import "errors"

type Embedder interface {
	Embed(text string) ([]float64, error)
	CosineSimilarity(a, b []float64) float64
}

type SimilarityScorer interface {
	Compute(texts []string) ([]float64, error)
}

var ErrEmbedderNotSet = errors.New("embedder not set")