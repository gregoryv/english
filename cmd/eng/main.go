package main

import (
	"fmt"
	"os"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/english"
)

func main() {
	var (
		cli  = cmdline.NewParser(os.Args...)
		help = cli.Flag("-h, --help")
		p    = cli.Option(
			"-p, --print",
			"word, verb, noun or adjective(adj)",
		).String("")

		r = cli.Option("-r", "random: word, verb, noun or adjective(adj)").String("")
		w = os.Stderr
	)

	switch {
	case help:
		cli.WriteUsageTo(w)

	case !cli.Ok():
		fmt.Fprintln(w, cli.Error())
		fmt.Fprintln(w, "Try --help for more information")

	}

	lang := english.NewLanguage()
	switch {
	case p != "":
		for _, word := range wordsBy(lang, p) {
			fmt.Println(word)
		}

	case r != "":
		word := english.RandWord(wordsBy(lang, r))
		fmt.Println(word)
	default:
		fmt.Println(lang)
	}
}

func wordsBy(lang *english.Language, m string) []english.Word {
	switch m {
	case "word":
		return lang.Words()
	case "verb":
		return lang.Verbs()
	case "noun":
		return lang.Nouns()
	case "adjective", "adj":
		return lang.Adjectives()
	}
	fmt.Println("Unknown:", m)
	os.Exit(1)
	return nil
}
