package english

import (
	"fmt"
	"strings"
)

func NewLanguage() *Language {
	lang := &Language{}

	lang.verbs = lang.AddWords(Verb, VerbWords)
	lang.adverbs = lang.addWords(Adverb, Adverbs())
	lang.nouns = lang.AddWords(Noun, NounWords)
	lang.adj = lang.AddWords(Adjective, AdjectiveWords)
	lang.prepositions = lang.addWords(Preposition, Prepositions())

	lang.Generator = NewGenerator()
	return lang
}

type Language struct {
	verbs        []Word
	adverbs      []Word
	nouns        []Word
	adj          []Word
	prepositions []Word
	words        []Word
	*Generator
}

// AddWords adds the words from the text with the given
// mode. Returns slice of the added words.
func (me *Language) AddWords(m Mode, text string) []Word {
	return me.addWords(m, strings.Fields(text))
}

func (me *Language) addWords(m Mode, inwords []string) []Word {
	from := len(me.words)
	for _, word := range inwords {
		me.words = append(me.words, Word{m, word})
	}
	return me.words[from : from+len(inwords)]
}

func (me *Language) Verbs() []Word        { return me.verbs }
func (me *Language) Adverbs() []Word      { return me.adverbs }
func (me *Language) Nouns() []Word        { return me.nouns }
func (me *Language) Adjectives() []Word   { return me.adj }
func (me *Language) Prepositions() []Word { return me.prepositions }
func (me *Language) Words() []Word        { return me.words }

func (me *Language) String() string {
	return fmt.Sprintf("%v verbs, %v adverbs, %v nouns, %v adjectives, %v prepositions, %v words",
		len(me.verbs),
		len(me.adverbs),
		len(me.nouns),
		len(me.adj),
		len(me.prepositions),
		len(me.words),
	)
}
