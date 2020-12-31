package english

import (
	"strings"
)

type Sentence []string

// String returns a string with first word capitalized and ending with
// a dot.
func (me Sentence) String() string {
	var sb strings.Builder
	for i, word := range me {
		if i == 0 {
			sb.WriteString(strings.Title(word))
		} else {
			sb.WriteString(word)
		}
		if i < len(me)-1 {
			sb.WriteByte(' ')
		}
	}
	sb.WriteByte('.')
	return sb.String()
}
