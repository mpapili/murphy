package occ

import (
	"math/rand"

	"github.com/mpapili/murphy/internal/proto"
)

// GenerateOccurrence builds a random occurrence for a game (skeleton rules).
func GenerateOccurrence(r *rand.Rand, gameID string, tick uint64) proto.Occurrence {
	types := AllTypes()
	t := types[r.Intn(len(types))]
	return proto.Occurrence{
		Type: string(t),
		Location: proto.OccurrenceLoc{
			GameID:  gameID,
			ChunkX:  int32(r.Intn(5) - 2),
			ChunkY:  0,
			ChunkZ:  int32(r.Intn(5) - 2),
			OffsetX: int32(r.Intn(16)),
			OffsetZ: int32(r.Intn(16)),
		},
		Magnitude: float32(r.Float64()),
		Duration:  uint32(30 + r.Intn(90)),
		StartTick: tick,
	}
}
