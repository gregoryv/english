package main

import (
	"flag"
	"fmt"

	"github.com/gregoryv/english"
)

func main() {
	r := flag.String("r", "", "random: word, verb, noun or adjective(adj)")
	flag.Parse()

	lang := english.NewLanguage()

	switch {
	case *r != "":
		var words func() []english.Word
		switch *r {
		case "word":
			words = lang.Words
		case "verb":
			words = lang.Verbs
		case "noun":
			words = lang.Nouns
		case "adjective", "adj":
			words = lang.Adjectives
		}
		fmt.Println(english.RandWord(words()))
	default:
		fmt.Println(lang)
	}
}
