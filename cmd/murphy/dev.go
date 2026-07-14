package main

import (
	"encoding/json"
	"net/http"

	"github.com/mpapili/murphy/internal/bus"
	"github.com/mpapili/murphy/internal/emit"
	"github.com/mpapili/murphy/internal/proto"
	"github.com/nats-io/nats.go"
)

// devOccurrenceBody is optional JSON for POST /dev/occurrence.
type devOccurrenceBody struct {
	Type      string               `json:"type"`
	Location  *proto.OccurrenceLoc `json:"location,omitempty"`
	Magnitude *float32             `json:"magnitude,omitempty"`
	Duration  *uint32              `json:"duration,omitempty"`
	StartTick *uint64              `json:"startTick,omitempty"`
}

func defaultSpawnOccurrence(gameID string) proto.Occurrence {
	return proto.Occurrence{
		Type: "spawn",
		Location: proto.OccurrenceLoc{
			GameID:  gameID,
			ChunkX:  0,
			ChunkY:  0,
			ChunkZ:  0,
			OffsetX: 8,
			OffsetZ: 8,
		},
		Magnitude: 1,
		Duration:  60,
		StartTick: 0,
	}
}

// handleDevOccurrence publishes a specified (or default spawn) occurrence on NATS.
// Local-dev control plane, mirrors budowac-server /dev/spawn-goat.
func handleDevOccurrence(nc *nats.Conn, gameID string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST required", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		occ := defaultSpawnOccurrence(gameID)
		if r.Body != nil {
			var body devOccurrenceBody
			dec := json.NewDecoder(r.Body)
			if err := dec.Decode(&body); err == nil {
				if body.Type != "" {
					occ.Type = body.Type
				}
				if body.Location != nil {
					occ.Location = *body.Location
					if occ.Location.GameID == "" {
						occ.Location.GameID = gameID
					}
				}
				if body.Magnitude != nil {
					occ.Magnitude = *body.Magnitude
				}
				if body.Duration != nil {
					occ.Duration = *body.Duration
				}
				if body.StartTick != nil {
					occ.StartTick = *body.StartTick
				}
			}
			// empty body / bad JSON still uses defaults — force-spawn path
		}

		if err := emit.Publish(nc, gameID, occ); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"ok": false, "error": err.Error(),
			})
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"ok":         true,
			"subject":    bus.MurphyOccurrence(gameID),
			"occurrence": occ,
		})
	}
}
