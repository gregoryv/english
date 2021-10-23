package english

import (
	"fmt"
	"log"
)

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

	finalClass // for testing ranges
)

// Random returns a random word of the given class.
func (me Class) Random() string {
	fn, found := classMap[me]
	if !found {
		err := fmt.Errorf("unknown class: %v", me)
		log.Println(err)
		return ""
	}
	return fn()
}

var classMap = map[Class]func() string{
	ClassNoun:      func() string { return randomField(_NounWords) },
	ClassVerb:      func() string { return randomField(_VerbWords) },
	ClassAdjective: func() string { return randomField(_AdjectiveWords) },

	ClassHowAdverb:        func() string { return randomField(_HowAdverbWords) },
	ClassWhenAdverb:       func() string { return randomField(_WhenAdverbWords) },
	ClassWhereAdverb:      func() string { return randomField(_WhereAdverbWords) },
	ClassWhatExtentAdverb: func() string { return randomField(_WhatExtentAdverbWords) },

	ClassPersonalPronoun:              func() string { return randomField(_PersonalPronounWords) },
	ClassPossessivePronoun:            func() string { return randomField(_PossessivePronounWords) },
	ClassIndependentPossessivePronoun: func() string { return randomField(_IndependentPossessivePronounWords) },
	ClassObjectPronoun:                func() string { return randomField(_ObjectPronounWords) },
	ClassIndefinitePronoun:            func() string { return randomField(_IndefinitePronounWords) },
	ClassReflexivePronoun:             func() string { return randomField(_ReflexivePronounWords) },
	ClassDemonstrativePronoun:         func() string { return randomField(_DemonstrativePronounWords) },
	ClassInterrogativePronoun:         func() string { return randomField(_InterrogativePronounWords) },
	ClassRelativePronoun:              func() string { return randomField(_RelativePronounWords) },
	ClassArchaicPronoun:               func() string { return randomField(_ArchaicPronounWords) },

	ClassQuestion:    func() string { return randomField(_QuestionWords) },
	ClassConjunction: func() string { return randomField(_ConjunctionWords) },
}
