package english

import "testing"

func Test_word_count(t *testing.T) {
	got := len(Words())
	if got != WordCount {
		t.Errorf("WordCount != %v", got)
	}
}
