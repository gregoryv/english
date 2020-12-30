package english

import "testing"

func Test_Dict_RandWord(t *testing.T) {
	dict := NewDict()
	w := dict.RandWord()
	if w.String() == "" {
		t.Errorf("%s: %s", w.Group(), w)
	}
	g := w.Group()
	if g < 0 {
		t.Error("negative group")
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
