package occ

import "testing"

func TestAllTypes(t *testing.T) {
	types := AllTypes()
	if len(types) != 4 {
		t.Fatalf("AllTypes() len = %d, want 4", len(types))
	}

	expected := []OccurrenceType{TypeWeather, TypeAnomaly, TypeSpawn, TypeTimed}
	for i, want := range expected {
		if types[i] != want {
			t.Errorf("AllTypes()[%d] = %q, want %q", i, types[i], want)
		}
	}
}
