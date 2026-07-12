package occ

import (
	"math/rand"
	"testing"
)

func TestGenerateOccurrence(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	occ := GenerateOccurrence(r, "g1", 10)

	if occ.Location.GameID != "g1" {
		t.Errorf("GameID = %q, want %q", occ.Location.GameID, "g1")
	}
	if occ.StartTick != 10 {
		t.Errorf("StartTick = %d, want 10", occ.StartTick)
	}
	if occ.Type == "" {
		t.Error("Type should not be empty")
	}
	if occ.Duration < 30 || occ.Duration > 119 {
		t.Errorf("Duration = %d, want [30, 120)", occ.Duration)
	}
	if occ.Magnitude < 0 || occ.Magnitude > 1 {
		t.Errorf("Magnitude = %f, want [0, 1]", occ.Magnitude)
	}
}

func TestGenerateOccurrenceTypeInCatalog(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	occ := GenerateOccurrence(r, "g1", 1)

	found := false
	for _, tt := range AllTypes() {
		if string(tt) == occ.Type {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Generated type %q not in catalog", occ.Type)
	}
}

func TestGenerateOccurrenceOffsetRange(t *testing.T) {
	r := rand.New(rand.NewSource(7))
	for i := 0; i < 20; i++ {
		occ := GenerateOccurrence(r, "g1", uint64(i))
		if occ.Location.OffsetX < 0 || occ.Location.OffsetX >= 16 {
			t.Errorf("OffsetX = %d, want [0, 16)", occ.Location.OffsetX)
		}
		if occ.Location.OffsetZ < 0 || occ.Location.OffsetZ >= 16 {
			t.Errorf("OffsetZ = %d, want [0, 16)", occ.Location.OffsetZ)
		}
	}
}
