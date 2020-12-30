package english

import "math/rand"

func NewLanguage() *Language {
	verbs := Verbs()
	nouns := Nouns()
	words := make([]Word, len(verbs)+len(nouns))
	lang := &Language{words: words}
	var i int
	for _, word := range verbs {
		words[i] = Word{Verb, word}
		i++
	}
	lang.verbs = words[:i]
	nounStart := i
	for _, word := range nouns {
		words[i] = Word{Noun, word}
		i++
	}
	lang.nouns = words[nounStart:i]
	return lang
}

type Language struct {
	verbs []Word
	nouns []Word
	adj   []Word
	words []Word
}

func (me *Language) Verbs() []Word      { return me.verbs }
func (me *Language) Nouns() []Word      { return me.nouns }
func (me *Language) Adjectives() []Word { return me.adj }
func (me *Language) Words() []Word      { return me.words }

func (me *Language) RandWord() Word {
	w := me.Words()
	return w[rand.Intn(len(w))]
}
