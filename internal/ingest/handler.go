package ingest

import (
	"encoding/json"
	"log"

	"github.com/mpapili/murphy/internal/bus"
	"github.com/mpapili/murphy/internal/proto"
	"github.com/nats-io/nats.go"
)

// SubscribeState listens for server state summaries to bias future occurrences.
func SubscribeState(nc *nats.Conn, gameID string) {
	if nc == nil {
		return
	}
	_, err := nc.Subscribe(bus.ServerState(gameID), func(msg *nats.Msg) {
		var s proto.StateSummary
		if err := json.Unmarshal(msg.Data, &s); err != nil {
			log.Printf("ingest: bad state: %v", err)
			return
		}
		log.Printf("ingest: state players=%d weather=%s tick=%d", s.PlayerCount, s.ActiveWeather, s.Tick)
		// Expandable: update scheduler weighting from state.
	})
	if err != nil {
		log.Printf("ingest: sub: %v", err)
	}
}
