package english

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// init rand Seed properly
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
	size, err := randSize(min, max)
	if err != nil {
		panic(err)
	}
	res := make([]string, size)
	for i := range res {
		res[i] = _words[rand.Intn(WordCount)]
	}
	return res
}

func randSize(min, max int) (int, error) {
	if min == 0 {
		return 0, fmt.Errorf("min must be > 0")
	}
	if min > max {
		return 0, fmt.Errorf("min > max")
	}
	if min == max {
		return max, nil
	}
	return rand.Intn(max-min) + min, nil
}

// RandomQuestion returns a question of random english words, trying
// to be as correct as possible. TODO implement grammar rules.
func RandomQuestion(min, max int) []string {
	size, err := randSize(min, max)
	if err != nil {
		panic(err)
	}
	res := make(Template, size)
	if len(res) >= 1 {
		res[0] = ClassQuestion
	}
	if len(res) >= 2 {
		res[1] = ClassVerb
	}

	if len(res) >= 3 {
		for i := range res[2:] {
			res[i+2] = ClassNoun
		}
	}
	return res.Build()
}

func randomField(txt string) string {
	fields := strings.Fields(txt)
	return fields[rand.Intn(len(fields))]

}
