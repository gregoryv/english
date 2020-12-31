package main

import (
	"strings"

	"github.com/gregoryv/english"
)

func randomText() string {
	sentences := make([]string, 5)
	for i := range sentences {
		s := make(english.Sentence, 5)
		for i := range s {
			s[i] = english.RandomWord()
		}
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
