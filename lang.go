package english

import (
	"fmt"
)

func NewLanguage() *Language {
	verbs := Verbs()
	nouns := Nouns()
	adj := Adjectives()
	words := make([]Word, len(verbs)+len(nouns)+len(adj))
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
	adjStart := i
	for _, word := range adj {
		words[i] = Word{Adjective, word}
		i++
	}
	lang.adj = words[adjStart:i]
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

func (me *Language) String() string {
	return fmt.Sprintf("%v verbs, %v nouns, %v adjectives, %v words",
		len(me.verbs),
		len(me.nouns),
		len(me.adj),
		len(me.words),
	)
}
