package english

import "strings"

func Prepositions() []string {
	return strings.Split(SinglePrepositionWords, "\n")
}

const SinglePrepositionWords = `about
beside
near
to
above
between
of
towards
across
beyond
off
under
after
by
on
underneath
against
despite
onto
unlike
along
down
opposite
until
among
during
out
up
around
except
outside
upon
as
for
over
via
at
from
past
with
before
in
round
within
behind
inside
since
without
below
into
than
beneath
like
through`