package tiling

import (
	"math"

	"github.com/texttiling/texttiling/pkg/config"
)

type BoundaryDetector struct {
	smoothingWidth int
	smoothingRounds int
	cutoffPolicy    config.CutoffPolicy
	minSize        int
	maxSize        int
}

func NewBoundaryDetector(opts *config.Options) *BoundaryDetector {
	return &BoundaryDetector{
		smoothingWidth: opts.SmoothingWidth,
		smoothingRounds: opts.SmoothingRounds,
		cutoffPolicy:    opts.CutoffPolicy,
		minSize:        opts.MinSegmentSize,
		maxSize:        opts.MaxSegmentSize,
	}
}

func (bd *BoundaryDetector) Smooth(scores []float64) []float64 {
	if len(scores) == 0 {
		return scores
	}

	result := make([]float64, len(scores))
	copy(result, scores)

	width := bd.smoothingWidth
	if width > len(scores)/2 {
		width = len(scores) / 2
	}
	if width < 1 {
		width = 1
	}

	for round := 0; round < bd.smoothingRounds; round++ {
		smoothed := make([]float64, len(result))
		for i := 0; i < len(result); i++ {
			start := i - width
			end := i + width + 1
			if start < 0 {
				start = 0
			}
			if end > len(result) {
				end = len(result)
			}

			sum := 0.0
			count := 0
			for j := start; j < end; j++ {
				sum += result[j]
				count++
			}
			smoothed[i] = sum / float64(count)
		}
		result = smoothed
	}

	return result
}

func (bd *BoundaryDetector) ComputeDepth(scores []float64) []float64 {
	if len(scores) < 3 {
		return make([]float64, len(scores))
	}

	depth := make([]float64, len(scores))

	for i := 1; i < len(scores)-1; i++ {
		leftDiff := scores[i-1] - scores[i]
		rightDiff := scores[i+1] - scores[i]

		if leftDiff > rightDiff {
			depth[i] = leftDiff
		} else {
			depth[i] = rightDiff
		}
	}

	if len(scores) > 1 {
		lastIdx := len(scores) - 1
		secondLastIdx := len(scores) - 2
		depth[0] = scores[1] - scores[0]
		depth[lastIdx] = scores[secondLastIdx] - scores[lastIdx]
	}

	return depth
}

func (bd *BoundaryDetector) DetectBoundaries(depthScores []float64) []bool {
	if len(depthScores) == 0 {
		return nil
	}

	threshold := bd.computeThreshold(depthScores)

	boundaries := make([]bool, len(depthScores))
	for i, score := range depthScores {
		if score > threshold {
			boundaries[i] = true
		}
	}

	boundaries = bd.enforceSizeConstraints(boundaries)

	if boundaries[0] {
		boundaries[0] = false
	}

	return boundaries
}

func (bd *BoundaryDetector) computeThreshold(depthScores []float64) float64 {
	sum := 0.0
	for _, score := range depthScores {
		sum += score
	}
	mean := sum / float64(len(depthScores))

	sumSq := 0.0
	for _, score := range depthScores {
		diff := score - mean
		sumSq += diff * diff
	}
	stdDev := math.Sqrt(sumSq / float64(len(depthScores)))

	switch bd.cutoffPolicy {
	case config.LC:
		return mean - 0.5*stdDev
	default:
		return mean + 0.5*stdDev
	}
}

func (bd *BoundaryDetector) enforceSizeConstraints(boundaries []bool) []bool {
	result := make([]bool, len(boundaries))
	copy(result, boundaries)

	minGaps := bd.minSize / 20
	if minGaps < 1 {
		minGaps = 1
	}

	i := 0
	for i < len(result) {
		if !result[i] {
			i++
			continue
		}

		count := 0
		for j := i; j < len(result) && j < i+minGaps && result[j]; j++ {
			count++
		}
		if count >= minGaps {
			for j := i - count + 1; j < i; j++ {
				if j > 0 {
					result[j] = false
				}
			}
		}

		i++
	}

	return result
}