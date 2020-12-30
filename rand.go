package english

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandWord() Word {
	w := Words()
	return w[rand.Intn(len(w))]
}

func Words() []Word {
	verbs := Verbs()
	nouns := Nouns()
	words := make([]Word, len(verbs)+len(nouns))
	var i int
	for _, word := range verbs {
		words[i] = Word{Verb, word}
		i++
	}
	for _, word := range nouns {
		words[i] = Word{Noun, word}
		i++
	}
	return words
}
