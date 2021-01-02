package english

import "fmt"

type Class int

const (
	ClassNoun Class = iota
	ClassVerb
	ClassAdjective

	ClassHowAdverb
	ClassWhenAdverb
	ClassWhereAdverb
	ClassWhatExtentAdverb

	ClassPersonalPronoun
	ClassPossessivePronoun
	ClassIndependentPossessivePronoun
	ClassObjectPronoun
	ClassIndefinitePronoun
	ClassReflexivePronoun
	ClassDemonstrativePronoun
	ClassInterrogativePronoun
	ClassRelativePronoun
	ClassArchaicPronoun

	ClassQuestion
	ClassConjunction
)

// Random
func (me Class) Random() string {
	switch me {
	case ClassNoun:
		return randomField(_NounWords)
	case ClassQuestion:
		return randomField(_QuestionWords)
	case ClassVerb:
		return randomField(_VerbWords)
	default:
		panic(fmt.Errorf("unknown class: %v", me))
	}
}
