package tiling

import (
	"testing"

	"github.com/texttiling/texttiling/pkg/config"
)

func TestNew_Default(t *testing.T) {
	tt, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if tt == nil {
		t.Fatal("New() returned nil")
	}
}

func TestNew_CustomOptions(t *testing.T) {
	tt, err := New(
		config.WithPseudosentenceSize(15),
		config.WithBlockSize(8),
	)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if tt == nil {
		t.Fatal("New() returned nil")
	}
}

func TestSegment_Empty(t *testing.T) {
	tt, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	segments, err := tt.Segment("")
	if err != nil {
		t.Fatalf("Segment() error = %v", err)
	}
	if len(segments) != 0 {
		t.Errorf("Segment() on empty text = %d segments, want 0", len(segments))
	}
}

func TestSegment_SingleSegment(t *testing.T) {
	tt, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	text := "This is a short paragraph with some words."
	segments, err := tt.Segment(text)
	if err != nil {
		t.Fatalf("Segment() error = %v", err)
	}
	if len(segments) != 1 {
		t.Errorf("Segment() on short text = %d segments, want 1", len(segments))
	}
}

func TestSegment_BasicDocument(t *testing.T) {
	tt, err := New(config.WithCutoffPolicy(config.LC))
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	text := `Artificial intelligence is transforming our world.
Machine learning enables computers to learn from data.
Deep learning uses neural networks with many layers.
Natural language processing helps machines understand text.
Sentiment analysis determines the emotion in text.
Text summarization creates short summaries.
Computer vision allows machines to see images.
Object detection identifies objects in photos.
Face recognition identifies human faces.`

	segments, err := tt.Segment(text)
	if err != nil {
		t.Fatalf("Segment() error = %v", err)
	}
	if len(segments) != 1 {
		t.Errorf("Segment() = %d segments, want 1", len(segments))
	}
}

func TestSegment_WithDebug(t *testing.T) {
	tt, err := New(config.WithCutoffPolicy(config.LC))
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	text := `First topic discusses machine learning.
Machine learning is a subset of artificial intelligence.
It enables systems to learn from data.
Second topic covers deep learning.
Deep learning uses neural networks.
Convolutional networks are used for image recognition.
Third topic is about reinforcement learning.
Q-learning is a popular algorithm.
Policy gradients help with control tasks.`

	result, err := tt.SegmentWithDebug(text)
	if err != nil {
		t.Fatalf("SegmentWithDebug() error = %v", err)
	}
	if len(result.Segments) != 1 {
		t.Errorf("SegmentWithDebug() segments = %d, want 1", len(result.Segments))
	}
	if len(result.DepthScores) == 0 {
		t.Error("SegmentWithDebug() DepthScores should not be empty")
	}
}

func TestSegment_WithGaps(t *testing.T) {
	tt, err := New(config.WithCutoffPolicy(config.LC))
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	text := `This is topic one.
More about topic one.
Still topic one.
More details here.
Another sentence.

This is topic two.
More about topic two.
Different content here.`

	segments, err := tt.SegmentWithGaps(text)
	if err != nil {
		t.Fatalf("SegmentWithGaps() error = %v", err)
	}
	if len(segments) < 1 {
		t.Skip("Not enough segments to test gaps")
	}

	for i, seg := range segments {
		if i == 0 && seg.GapBefore {
			t.Error("First segment should not have gap before")
		}
		if i == len(segments)-1 && seg.GapAfter {
			t.Error("Last segment should not have gap after")
		}
		if !seg.GapBefore && !seg.GapAfter && len(segments) == 1 {
			t.Logf("Single segment has no gaps - OK")
		}
	}
}

func TestSegment_ConfigOptions(t *testing.T) {
	tests := []struct {
		name    string
		opts   []config.Option
		wantOK bool
	}{
		{
			name: "default",
			opts: []config.Option{},
		},
		{
			name: "block size 5",
			opts: []config.Option{config.WithBlockSize(5)},
		},
		{
			name: "LC cutoff",
			opts: []config.Option{config.WithCutoffPolicy(config.LC)},
		},
		{
			name: "custom language - german",
			opts: []config.Option{config.WithLanguage("de")},
		},
		{
			name: "all options",
			opts: []config.Option{
				config.WithPseudosentenceSize(10),
				config.WithBlockSize(5),
				config.WithCutoffPolicy(config.LC),
				config.WithSmoothing(3, 2),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.opts...)
			if err != nil {
				t.Errorf("New(%v) error = %v", tt.name, err)
			}
		})
	}
}