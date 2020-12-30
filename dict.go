package english

import "math/rand"

func NewDict() *Dict {
	return &Dict{
		verbs: Verbs(),
		nouns: Nouns(),
		adj:   Adjectives(),
	}
}

type Dict struct {
	verbs []string
	nouns []string
	adj   []string
}

// Verbs
func (me *Dict) Verbs() []string      { return me.verbs }
func (me *Dict) Nouns() []string      { return me.nouns }
func (me *Dict) Adjectives() []string { return me.adj }

func (me *Dict) RandWord() Word {
	w := me.Words()
	return w[rand.Intn(len(w))]
}

func (me *Dict) Words() []Word {
	verbs := me.verbs
	nouns := me.nouns
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
