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
