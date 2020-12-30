package english

import "strings"

func Nouns() []string {
	return strings.Split(strings.TrimSpace(nouns), "\n")
}

const nouns = `
bicycle
boat
car
computer
cup
glas
house
keyboard
mouse
radio
tv
`
