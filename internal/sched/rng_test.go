package sched

import "testing"

func TestNewSeededDeterministic(t *testing.T) {
	r1 := NewSeeded(42)
	r2 := NewSeeded(42)

	for i := 0; i < 5; i++ {
		a := r1.Intn(1000)
		b := r2.Intn(1000)
		if a != b {
			t.Errorf("iteration %d: %d != %d (same seed should produce same sequence)", i, a, b)
		}
	}
}

func TestNewSeededDifferent(t *testing.T) {
	r1 := NewSeeded(1)
	r2 := NewSeeded(2)

	a := r1.Intn(1000)
	b := r2.Intn(1000)
	if a == b {
		t.Errorf("different seeds produced same first value %d", a)
	}
}
