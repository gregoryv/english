package english

type Word struct {
	group Group
	v     string
}

func (me Word) Group() Group   { return me.group }
func (me Word) String() string { return me.v }

//go:generate stringer -type=Group
type Group int

const (
	Verb Group = iota
	Noun
	Adjective
	Injection
)
