package english

import "strings"

func Verbs() []string {
	return strings.Split(strings.TrimSpace(verbs), "\n")
}

const verbs = `
close
open
run
say
speak
talk
walk
`
