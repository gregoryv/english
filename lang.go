package english

import (
	"fmt"
)

func NewLanguage() *Language {
	verbs := Verbs()
	nouns := Nouns()
	adj := Adjectives()
	lang := &Language{}

	lang.verbs = lang.AddWords(Verb, verbs)
	lang.nouns = lang.AddWords(Noun, nouns)
	lang.adj = lang.AddWords(Adjective, adj)
	lang.Generator = NewGenerator(lang)
	return lang
}

type Language struct {
	verbs []Word
	nouns []Word
	adj   []Word
	words []Word
	*Generator
}

// AddWords adds the slice of words with the given mode. Returns slice
// of the added words.
func (me *Language) AddWords(m Mode, inwords []string) []Word {
	from := len(me.words)
	for _, word := range inwords {
		me.words = append(me.words, Word{m, word})
	}
	return me.words[from : from+len(inwords)]
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
