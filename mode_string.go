// Code generated by "stringer -type=Mode"; DO NOT EDIT.

package english

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Verb-2147483648]
	_ = x[Noun-1073741824]
	_ = x[Adjective-536870912]
	_ = x[Injection-268435456]
	_ = x[Adverb-134217728]
	_ = x[Preposition-67108864]
}

const (
	_Mode_name_0 = "Preposition"
	_Mode_name_1 = "Adverb"
	_Mode_name_2 = "Injection"
	_Mode_name_3 = "Adjective"
	_Mode_name_4 = "Noun"
	_Mode_name_5 = "Verb"
)

func (i Mode) String() string {
	switch {
	case i == 67108864:
		return _Mode_name_0
	case i == 134217728:
		return _Mode_name_1
	case i == 268435456:
		return _Mode_name_2
	case i == 536870912:
		return _Mode_name_3
	case i == 1073741824:
		return _Mode_name_4
	case i == 2147483648:
		return _Mode_name_5
	default:
		return "Mode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
