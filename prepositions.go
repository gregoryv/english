package english

import "strings"

func Prepositions() []string {
	return strings.Split(strings.TrimSpace(prepositions), "\n")
}

const prepositions = `
a
aboard
about
above
abreast
absent
across
after against
along
aloft
alongside
amid
among
anti
apropos
around
as
aslant
astride
at
atop
bar
barring
before
behind
below
`

// https://en.wikipedia.org/wiki/List_of_English_prepositions#cite_ref-CGEL618_107-38
