package english

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"strings"
	"time"
)

func init() {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		rand.Seed(time.Now().Unix())
		return
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

// RandomWord returns a random english word.
func RandomWord() string {
	return _words[rand.Intn(WordCount)]
}

// RandomWords returns a slice of random english words.
func RandomWords(n int) []string {
	res := make([]string, n)
	for i := range res {
		res[i] = _words[rand.Intn(WordCount)]
	}
	return res
}

// RandomStatement returns a sentence of random english words, trying
// to be as correct as possible. TODO implement grammar rules.
func RandomStatement(min, max int) []string {
	if min == 0 {
		panic("min must be > 0")
	}
	size := rand.Intn(max-min) + min
	res := make([]string, size)
	for i := range res {
		res[i] = _words[rand.Intn(WordCount)]
	}
	return res
}

// RandomQuestion returns a question of random english words, trying
// to be as correct as possible. TODO implement grammar rules.
func RandomQuestion(min, max int) []string {
	if min == 0 {
		panic("min must be > 0")
	}
	size := rand.Intn(max-min) + min
	res := make([]string, size)
	if len(res) >= 1 {
		res[0] = randomField(_QuestionWords)
	}
	if len(res) >= 2 {
		res[1] = randomField(_VerbWords)
	}

	if len(res) >= 3 {
		for i := range res[2:] {
			res[i+2] = RandomWord()
		}
	}
	return res
}

func randomField(txt string) string {
	fields := strings.Fields(txt)
	return fields[rand.Intn(len(fields))]

}
