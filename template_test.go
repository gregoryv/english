package english

import (
	"fmt"
	"math/rand"
)

func ExampleTemplate_Build() {
	rand.Seed(0)
	tpl := Template{
		ClassHowAdverb,
		ClassVerb,
		ClassNoun,
		"rocks",
		3,
		"times",
	}
	fmt.Println(tpl.Build())
	// output:
	// [lazily solve head rocks 3 times]
}
