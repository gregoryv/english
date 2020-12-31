package main

import (
	"strings"

	"github.com/gregoryv/english"
)

func randomText(lang *english.Language) string {
	sentences := make([]string, 5)
	for i := range sentences {
		sentences[i] = lang.RandSentence(5).String()
	}
	return strings.Join(sentences, " ")
}

func longestWord(lang *english.Language) int {
	var l, wordLen int
	for _, word := range lang.Words() {
		wordLen = len(word.String())
		if wordLen > l {
			l = wordLen
		}
	}
	return l
}
