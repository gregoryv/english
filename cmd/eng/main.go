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
			"words",
		).Bool()

		r = cli.Option("-r, --random", "word").Bool()
		w = os.Stderr
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
		fmt.Println(english.Words())

	case r:
		fmt.Println(english.RandomWord())

	default:
		fmt.Println(len(english.Words()), "words")
	}
}
