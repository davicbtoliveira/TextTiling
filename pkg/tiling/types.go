package tiling

type Segment struct {
	Index     int    `json:"index"`
	Text     string `json:"text"`
	StartPos int    `json:"start_pos"`
	EndPos   int    `json:"end_pos"`
	Size     int    `json:"size"`
}

type Pseudosentence struct {
	Index      int
	Tokens    []string
	Positions []int
}

type Result struct {
	Segments        []Segment       `json:"segments"`
	Pseudosentences []Pseudosentence `json:"pseudosentences"`
	RawScores      []float64     `json:"raw_scores"`
	SmoothedScores []float64     `json:"smoothed_scores"`
	DepthScores   []float64     `json:"depth_scores"`
	Boundaries    []bool       `json:"boundaries"`
}

type SegmentWithGap struct {
	Segment   Segment `json:"segment"`
	GapBefore bool    `json:"gap_before"`
	GapAfter  bool    `json:"gap_after"`
}