package english

import (
	"strings"
)

// Sentence returns a string starting with capitalized word and ending
// with the given byte.
func Sentence(words []string, end byte) string {
	var sb strings.Builder
	for i, word := range words {
		if i == 0 {
			sb.WriteString(strings.Title(word))
		} else {
			sb.WriteString(word)
		}
		if i < len(words)-1 {
			sb.WriteByte(' ')
		}
	}
	sb.WriteByte(end)
	return sb.String()
}
