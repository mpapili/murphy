package sched

import "math/rand"

// NewSeeded returns an RNG for a game id. Expandable: persist seed in Postgres.
func NewSeeded(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}
