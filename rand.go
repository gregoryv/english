package english

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
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
	return RandomWords(1)[0]
}

// RandomWords returns a slice of random english words.
func RandomWords(n int) []string {
	res := make([]string, n)
	words := Words()
	for i := range res {
		res[i] = words[rand.Intn(len(words))]
	}
	return res
}
