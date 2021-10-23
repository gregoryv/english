package english

import "testing"

func Test_Class_Random(t *testing.T) {
	var class Class
	for ; class < finalClass; class++ {
		v := class.Random()
		if v == "" {
			t.Errorf("no random word for %v", v)
		}
	}
}
