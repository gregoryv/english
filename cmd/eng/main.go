// Command eng prints english words.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/english"
)

func main() {
	var (
		cli  = cmdline.NewParser(os.Args...)
		help = cli.Flag("-h, --help")
		p    = cli.Option("-p, --print", "Print all words").Bool()
		r    = cli.Option("-r, --random", "Print one random word").Int(0)
		w    = os.Stderr
	)

	switch {
	case help:
		cli.WriteUsageTo(w)
		os.Exit(0)

	case !cli.Ok():
		fmt.Fprintln(w, cli.Error())
		fmt.Fprintln(w, "Try --help for more information")
		os.Exit(1)
	}

	switch {
	case p:
		fmt.Println(strings.Join(english.Words(), " "))

	case r > 0:
		fmt.Println(strings.Join(english.RandomWords(r), " "))

	default:
		fmt.Println(len(english.Words()), "words")
	}
}
