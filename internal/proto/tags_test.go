package proto

import "testing"

func TestMessageTagString(t *testing.T) {
	cases := []struct {
		tag    MessageTag
		expect string
	}{
		{TagHello, "Hello"},
		{TagAuth, "Auth"},
		{TagChunkRequest, "ChunkRequest"},
		{TagChunkData, "ChunkData"},
		{TagChunkDiff, "ChunkDiff"},
		{TagPlayerInput, "PlayerInput"},
		{TagSnapshot, "Snapshot"},
		{TagEntityState, "EntityState"},
		{TagEdit, "Edit"},
		{TagOccurrence, "Occurrence"},
		{TagStateSummary, "StateSummary"},
		{TagJoin, "Join"},
		{TagLeave, "Leave"},
		{MessageTag(255), "Unknown"},
	}
	for _, c := range cases {
		if got := c.tag.String(); got != c.expect {
			t.Errorf("MessageTag(%d).String() = %q, want %q", c.tag, got, c.expect)
		}
	}
}
