package bus

import "testing"

func TestServerState(t *testing.T) {
	got := ServerState("g1")
	want := "budowac.server.g1.state"
	if got != want {
		t.Errorf("ServerState() = %q, want %q", got, want)
	}
}

func TestMurphyOccurrence(t *testing.T) {
	got := MurphyOccurrence("my-game")
	want := "budowac.murphy.my-game.occurrence"
	if got != want {
		t.Errorf("MurphyOccurrence() = %q, want %q", got, want)
	}
}
