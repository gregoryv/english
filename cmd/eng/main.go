package main

import (
	"flag"
	"fmt"

	"github.com/gregoryv/english"
)

func main() {
	r := flag.String("r", "", "random [word]")
	flag.Parse()

	eng := english.NewDict()

	switch *r {
	case "word":
		fmt.Println(eng.RandWord())
	default:
		verbs := eng.Verbs()
		nouns := eng.Nouns()
		adj := eng.Adjectives()

		fmt.Println(len(verbs), "verbs")
		fmt.Println(len(nouns), "nouns")
		fmt.Println(len(adj), "adjectives")
		fmt.Println(len(eng.Words()), "words")
	}
}
