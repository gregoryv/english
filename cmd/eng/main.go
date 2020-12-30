package main

import (
	"flag"
	"fmt"

	"github.com/gregoryv/english"
)

func main() {
	r := flag.String("r", "", "random [word]")
	flag.Parse()

	lang := english.NewLanguage()

	switch *r {
	case "word":
		fmt.Println(english.RandWord(lang.Words()))
	default:
		fmt.Println(lang)
	}
}
