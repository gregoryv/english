package english

import "testing"

func Test_Generator(t *testing.T) {
	lang := NewLanguage()

	s := lang.RandSentence(5).String()
	if s[len(s)-1] != '.' {
		t.Error("should end with .: ", s)
	}

	w := lang.RandWord()
	if w.String() == "" {
		t.Error(w)
	}
}
