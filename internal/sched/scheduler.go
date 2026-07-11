package sched

import (
	"context"
	"log"
	"time"

	"github.com/mpapili/murphy/internal/emit"
	"github.com/mpapili/murphy/internal/occ"
	"github.com/nats-io/nats.go"
)

// Run periodically emits occurrences while interval > 0.
func Run(ctx context.Context, nc *nats.Conn, gameID string, intervalSec int, seed int64) {
	if intervalSec <= 0 {
		log.Printf("sched: occurrence interval disabled")
		return
	}
	r := NewSeeded(seed)
	t := time.NewTicker(time.Duration(intervalSec) * time.Second)
	defer t.Stop()
	var tick uint64
	log.Printf("sched: emitting every %ds for game=%s", intervalSec, gameID)
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			tick++
			o := occ.GenerateOccurrence(r, gameID, tick)
			if err := emit.Publish(nc, gameID, o); err != nil {
				log.Printf("sched: emit: %v", err)
			}
		}
	}
}
