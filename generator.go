package english

import (
	"math/rand"
	"strings"
)

func NewGenerator() *Generator {
	return &Generator{}
}

type Generator struct{}

// RandWord returns a random word.
func (me *Generator) RandWord() string {
	return randWord(Words())
}

func randWord(words []string) string {
	return words[rand.Intn(len(words))]
}

// RandSentence returns a random sentence of n words.
func (me *Generator) RandSentence(n int) Sentence {
	s := make([]string, n)
	for i := range s {
		s[i] = me.RandWord()
	}
	return s
}

type Sentence []string

// String returns a string with first word capitalized and ending with
// a dot.
func (me Sentence) String() string {
	var sb strings.Builder
	for i, word := range me {
		if i == 0 {
			sb.WriteString(strings.Title(word))
		} else {
			sb.WriteString(word)
		}
		if i < len(me)-1 {
			sb.WriteByte(' ')
		}
	}
	sb.WriteByte('.')
	return sb.String()
}
