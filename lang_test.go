package english

import (
	"strings"
	"testing"
)

func Test_Language_RandWord(t *testing.T) {
	lang := NewLanguage()
	w := RandWord(lang.Words())
	if w.String() == "" {
		t.Error(w)
	}
	switch {
	case w.Is(Verb):
	case w.Is(Noun):
	case w.Is(Adjective):
	case w.Is(Adverb):
	case w.Is(Preposition):
	default:
		t.Error("is:", w.Mode.String())
	}
	g := w.Mode
	if g < 0 {
		t.Error("negative group")
	}
}

func Test_Language(t *testing.T) {
	lang := NewLanguage()
	for _, word := range lang.Words() {
		w := strings.TrimSpace(word.String())
		if w == "" {
			t.Error("found empty word")
		}
	}
	if len(lang.Verbs()) == 0 {
		t.Error("no verbs")
	}
	if len(lang.Nouns()) == 0 {
		t.Error("no nouns")
	}
	if len(lang.Adjectives()) == 0 {
		t.Error("no adjectives")
	}
}

func Test_Language_String(t *testing.T) {
	lang := NewLanguage()
	if lang.String() == "" {
		t.Error("empty string")
	}
}
