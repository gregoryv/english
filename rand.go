package english

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandWord(words []Word) Word {
	return words[rand.Intn(len(words))]
}
