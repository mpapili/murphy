package emit

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mpapili/murphy/internal/bus"
	"github.com/mpapili/murphy/internal/proto"
	"github.com/nats-io/nats.go"
)

// Publish sends an Occurrence on the murphy subject for gameID.
func Publish(nc *nats.Conn, gameID string, o proto.Occurrence) error {
	if nc == nil {
		return fmt.Errorf("nats not connected")
	}
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}
	subj := bus.MurphyOccurrence(gameID)
	if err := nc.Publish(subj, b); err != nil {
		return err
	}
	log.Printf("emit: %s type=%s", subj, o.Type)
	return nil
}
