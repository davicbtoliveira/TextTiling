package text

import (
	"testing"
)

func TestLoadStopwords_Default(t *testing.T) {
	words, err := LoadStopwords("", "en")
	if err != nil {
		t.Fatalf("LoadStopwords() error = %v", err)
	}
	if len(words) == 0 {
		t.Fatal("LoadStopwords() returned empty map")
	}
	if _, ok := words["the"]; !ok {
		t.Error("Default stopwords should include 'the'")
	}
}

func TestLoadStopwords_German(t *testing.T) {
	words, err := LoadStopwords("", "de")
	if err != nil {
		t.Fatalf("LoadStopwords() error = %v", err)
	}
	if len(words) == 0 {
		t.Fatal("LoadStopwords() returned empty map")
	}
	if _, ok := words["der"]; !ok {
		t.Error("German stopwords should include 'der'")
	}
}

func TestLoadStopwords_French(t *testing.T) {
	words, err := LoadStopwords("", "fr")
	if err != nil {
		t.Fatalf("LoadStopwords() error = %v", err)
	}
	if len(words) == 0 {
		t.Fatal("LoadStopwords() returned empty map")
	}
	if _, ok := words["le"]; !ok {
		t.Error("French stopwords should include 'le'")
	}
}

func TestLoadStopwords_Spanish(t *testing.T) {
	words, err := LoadStopwords("", "es")
	if err != nil {
		t.Fatalf("LoadStopwords() error = %v", err)
	}
	if len(words) == 0 {
		t.Fatal("LoadStopwords() returned empty map")
	}
	if _, ok := words["el"]; !ok {
		t.Error("Spanish stopwords should include 'el'")
	}
}

func TestLoadStopwords_CustomFile(t *testing.T) {
	t.Skip("Custom file test requires os.WriteFile - skipping")
}