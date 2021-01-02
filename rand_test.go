package english

import "testing"

func Test_RandomStatement(t *testing.T) {
	ok := func(min, max int) {
		t.Helper()
		q := RandomStatement(min, max)
		if len(q) > max {
			t.Errorf("max exceeded: %q", Sentence(q, '.'))
		}
		if len(q) < min {
			t.Errorf("min exceeded: %q", Sentence(q, '.'))
		}
	}
	ok(1, 5)
	ok(2, 100)
	ok(4, 4)
}

func Test_RandomQuestion(t *testing.T) {
	max := 4
	q := RandomQuestion(1, max)
	if len(q) > max {
		t.Errorf("max exceeded: %q", Sentence(q, '?'))
	}
}

func Test_RandomWord(t *testing.T) {
	a := RandomWord()
	b := RandomWord()
	c := RandomWord()

	if a == b && b == c {
		t.Error("poor randomness")
	}
}

func Test_RandomWords(t *testing.T) {
	r := RandomWords(3)

	if r[0] == r[1] && r[1] == r[2] {
		t.Error("poor randomness")
	}
}

// ----------------------------------------

func Benchmark_RandomWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomWord()
	}
}

func Benchmark_10_RandomWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomWords(10)
	}
}
