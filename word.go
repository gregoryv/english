package english

import "fmt"

type Word struct {
	Mode
	v string
}

func (me Word) String() string { return fmt.Sprintf("%s: %s", me.Mode, me.v) }

//go:generate stringer -type=Mode
type Mode uint32

const (
	Verb Mode = 1 << (32 - 1 - iota)
	Noun
	Adjective
	Injection
)

func (me Mode) Is(v Mode) bool { return (me & v) == v }
