package emit

import (
	"testing"

	"github.com/mpapili/murphy/internal/proto"
)

func TestPublishNilConn(t *testing.T) {
	err := Publish(nil, "g1", proto.Occurrence{})
	if err == nil {
		t.Error("Publish with nil conn should return error")
	}
}
