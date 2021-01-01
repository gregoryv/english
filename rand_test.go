package english

import "testing"

func Benchmark_RandomWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomWord()
	}
}

func Benchmark_RandomWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomWords(10)
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
