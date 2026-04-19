package config

type SimilarityMethod int

const (
	BlockComparison SimilarityMethod = iota
	VocabularyIntroduction
)

type CutoffPolicy int

const (
	HC CutoffPolicy = iota
	LC
)

type Options struct {
	W                  int
	K                  int
	SimilarityMethod    SimilarityMethod
	SmoothingWidth     int
	SmoothingRounds    int
	CutoffPolicy       CutoffPolicy
	StopwordsFile      string
	Language          string
	DemoMode          bool
	EnableStemming     bool
	MinSegmentSize     int
	MaxSegmentSize    int
}

func Default() *Options {
	return &Options{
		W:               20,
		K:               10,
		SimilarityMethod: BlockComparison,
		SmoothingWidth:   2,
		SmoothingRounds: 1,
		CutoffPolicy:   HC,
		Language:     "en",
		MinSegmentSize: 50,
		MaxSegmentSize: 5000,
	}
}

type Option func(*Options)

func WithPseudosentenceSize(w int) Option {
	return func(o *Options) {
		if w > 0 {
			o.W = w
		}
	}
}

func WithBlockSize(k int) Option {
	return func(o *Options) {
		if k > 0 {
			o.K = k
		}
	}
}

func WithLanguage(lang string) Option {
	return func(o *Options) {
		o.Language = lang
	}
}

func WithSimilarityMethod(method SimilarityMethod) Option {
	return func(o *Options) {
		o.SimilarityMethod = method
	}
}

func WithCutoffPolicy(policy CutoffPolicy) Option {
	return func(o *Options) {
		o.CutoffPolicy = policy
	}
}

func WithSmoothing(width, rounds int) Option {
	return func(o *Options) {
		if width > 0 {
			o.SmoothingWidth = width
		}
		if rounds > 0 {
			o.SmoothingRounds = rounds
		}
	}
}

func WithDemoMode(enabled bool) Option {
	return func(o *Options) {
		o.DemoMode = enabled
	}
}

func WithStemming(enabled bool) Option {
	return func(o *Options) {
		o.EnableStemming = enabled
	}
}

func WithMinSegmentSize(size int) Option {
	return func(o *Options) {
		if size > 0 {
			o.MinSegmentSize = size
		}
	}
}

func WithMaxSegmentSize(size int) Option {
	return func(o *Options) {
		if size > 0 {
			o.MaxSegmentSize = size
		}
	}
}

func WithStopwordsFile(path string) Option {
	return func(o *Options) {
		o.StopwordsFile = path
	}
}