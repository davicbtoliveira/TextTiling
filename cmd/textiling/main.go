package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/texttiling/texttiling/pkg/config"
	"github.com/texttiling/texttiling/pkg/tiling"
)

func main() {
	opts := parseFlags()

	tt, err := tiling.New(opts...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating TextTiling: %v\n", err)
		os.Exit(1)
	}

	var text string
	if flag.NArg() > 0 {
		data, err := os.ReadFile(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}
		text = string(data)
	} else {
		fmt.Fprintln(os.Stderr, "Reading from stdin not implemented. Provide a file argument.")
		os.Exit(1)
	}

	segments, err := tt.Segment(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error segmenting: %v\n", err)
		os.Exit(1)
	}

	printSegments(segments)
}

func parseFlags() []config.Option {
	w := flag.Int("w", 20, "Pseudosentence size")
	k := flag.Int("k", 10, "Block size")
	method := flag.String("method", "block", "Similarity method (block|vocab)")
	policy := flag.String("policy", "hc", "Cutoff policy (hc|lc)")
	debug := flag.Bool("debug", false, "Show debug information")

	flag.Parse()

	opts := []config.Option{
		config.WithPseudosentenceSize(*w),
		config.WithBlockSize(*k),
	}

	if *method == "vocab" {
		opts = append(opts, config.WithSimilarityMethod(config.VocabularyIntroduction))
	}

	if *policy == "lc" {
		opts = append(opts, config.WithCutoffPolicy(config.LC))
	}

	if *debug {
		opts = append(opts, config.WithDemoMode(true))
	}

	return opts
}

func printSegments(segments []tiling.Segment) {
	for i, seg := range segments {
		text := seg.Text
		if len(text) > 100 {
			text = text[:100] + "..."
		}
		text = strings.ReplaceAll(text, "\n", "\\n")
		fmt.Printf("Segment %d: [%d chars]\n%s\n\n", i, len(seg.Text), text)
	}
	fmt.Printf("Total segments: %d\n", len(segments))
}