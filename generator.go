package english

import (
	"math/rand"
	"strings"
)

func NewGenerator(lang *Language) *Generator {
	return &Generator{
		lang,
	}
}

type Generator struct {
	lang *Language
}

// RandWord returns a random word.
func (me *Generator) RandWord() Word {
	return randWord(me.lang.Words())
}

func randWord(words []Word) Word {
	return words[rand.Intn(len(words))]
}

// RandSentence returns a random sentence of n words.
func (me *Generator) RandSentence(n int) Sentence {
	s := make([]Word, n)
	for i := range s {
		s[i] = me.RandWord()
	}
	return s
}

type Sentence []Word

// String returns a string with first word capitalized and ending with
// a dot.
func (me Sentence) String() string {
	var sb strings.Builder
	for i, word := range me {
		if i == 0 {
			sb.WriteString(strings.Title(word.String()))
		} else {
			sb.WriteString(word.String())
		}
		if i < len(me)-1 {
			sb.WriteByte(' ')
		}
	}
	sb.WriteByte('.')
	return sb.String()
}
