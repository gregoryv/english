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

// RandomSentence returns a sentence of random english words, trying
// to be as correct as possible. TODO implement grammar rules.
func RandomSentence(min, max int) []string {
	res := make([]string, max)
	for i := range res {
		res[i] = _words[rand.Intn(WordCount)]
	}
	return res
}

// RandomQuestion returns a question of random english words, trying
// to be as correct as possible. TODO implement grammar rules.
func RandomQuestion(min, max int) []string {
	return []string{
		randomField(_QuestionWords),
		RandomWord(),
		RandomWord(),
	}
}

func randomField(txt string) string {
	fields := strings.Fields(txt)
	return fields[rand.Intn(len(fields))]

}
