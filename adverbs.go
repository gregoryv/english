package english

import "strings"

// Adverbs are words that describe verbs other adverbs and adjetives.

func Adverbs() []string {
	all := make([]string, 0)
	for _, words := range []string{
		adverbsHow,
		adverbsWhen,
		adverbsWhere,
		adverbsWhatExtent,
	} {
		all = append(all, strings.Split(strings.TrimSpace(words), "\n")...)
	}
	return all
}

const adverbsHow = `
absentmindedly
adoringly
awkwardly
beatifully
briskly
brutally
carefully
cheerfully
competitively
eagerly
effortlessly
extravagantly
girlishly
gracefully
grimly
happily
halfheartedly
hungrily
lazily
lifelessly
loyally
quickly
quitely
quizzically
really
recklessly
remorsefully
ruthlessly
savagely
sloppily
so
stylishly
unabashedly
unevenly
urgently
well
wishfully
worriedly
`

const adverbsWhen = `
after
afterwards
annually
before
daily
never
now
soon
still
then
today
tomorrow
weekly
when
yesterday
`

const adverbsWhere = `
abroad
anywhere
away
down
everywhere
here
home
in
inside
out
outside
somewhere
there
underground
upstairs
`

const adverbsWhatExtent = `
extremely
not
quite
rather
really
terribly
too
very
`
