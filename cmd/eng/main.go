package main

import (
	"flag"
	"fmt"

	"github.com/gregoryv/english"
)

func main() {
	r := flag.String("r", "", "random [word]")
	flag.Parse()

	switch *r {
	case "word":
		fmt.Println(english.RandWord())
	default:
		verbs := english.Verbs()
		nouns := english.Nouns()
		adj := english.Adjectives()

		fmt.Println(len(verbs), "verbs")
		fmt.Println(len(nouns), "nouns")
		fmt.Println(len(adj), "adjectives")
		fmt.Println(len(english.Words()), "words")
	}
}
