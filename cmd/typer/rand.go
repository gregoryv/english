package main

import (
	"strings"

	"github.com/gregoryv/english"
)

func randomText() string {
	sentences := make([]string, 5)
	for i := range sentences {
		s := english.Sentence(english.RandomWords(5))
		sentences[i] = s.String()
	}
	return strings.Join(sentences, " ")
}

func longestWord() int {
	var max int
	for _, word := range english.Words() {
		if len(word) > max {
			max = len(word)
		}
	}
	return max
}
