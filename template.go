package english

import "fmt"

type Template []interface{}

func (me Template) Build() []string {
	res := make([]string, len(me))
	for i, v := range me {
		switch v := v.(type) {
		case Class:
			res[i] = v.Random()

		default:
			res[i] = fmt.Sprintf("%v", v)
		}
	}
	return res
}
