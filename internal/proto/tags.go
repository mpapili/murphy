// Package proto holds deliberately-duplicated wire types.
// Authoritative contract: budowac-space/initial-repo-spec.md
package proto

// MessageTag is the binary frame / logical message discriminator.
// Binary layouts are TBD; tag enums and field names are the contract.
type MessageTag uint8

const (
	TagHello MessageTag = iota
	TagAuth
	TagChunkRequest
	TagChunkData
	TagChunkDiff
	TagPlayerInput
	TagSnapshot
	TagEntityState
	TagEdit
	TagOccurrence
	TagStateSummary
	TagJoin
	TagLeave
)

func (t MessageTag) String() string {
	switch t {
	case TagHello:
		return "Hello"
	case TagAuth:
		return "Auth"
	case TagChunkRequest:
		return "ChunkRequest"
	case TagChunkData:
		return "ChunkData"
	case TagChunkDiff:
		return "ChunkDiff"
	case TagPlayerInput:
		return "PlayerInput"
	case TagSnapshot:
		return "Snapshot"
	case TagEntityState:
		return "EntityState"
	case TagEdit:
		return "Edit"
	case TagOccurrence:
		return "Occurrence"
	case TagStateSummary:
		return "StateSummary"
	case TagJoin:
		return "Join"
	case TagLeave:
		return "Leave"
	default:
		return "Unknown"
	}
}
