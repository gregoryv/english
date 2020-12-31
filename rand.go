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

func RandWord(words []Word) Word {
	return words[rand.Intn(len(words))]
}
