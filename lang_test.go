package english

import "testing"

func Test_Language_RandWord(t *testing.T) {
	lang := NewLanguage()
	w := RandWord(lang.Words())
	if w.String() == "" {
		t.Errorf("%s: %s", w.Group(), w)
	}
	g := w.Group()
	if g < 0 {
		t.Error("negative group")
	}
}

func Test_Language(t *testing.T) {
	lang := NewLanguage()
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

func Test_Group(t *testing.T) {
	g := Group(-1)
	if g.String() == "" {
		t.Error("empty group name: try go generate .")
	}

	g = Verb
	if g.String() == "" {
		t.Error("empty group name: try go generate .")
	}
}