package proto

import "testing"

func TestDefaultPalette(t *testing.T) {
	palette := DefaultPalette()
	if len(palette) != 6 {
		t.Fatalf("DefaultPalette() len = %d, want 6", len(palette))
	}

	checks := map[string]struct {
		solid, liquid, transparent, collidable bool
	}{
		"air":   {solid: false, liquid: false, transparent: true, collidable: false},
		"rock":  {solid: true, liquid: false, transparent: false, collidable: true},
		"water": {solid: false, liquid: true, transparent: true, collidable: false},
	}
	got := make(map[string]BrickDef, len(palette))
	for _, d := range palette {
		got[d.Name] = d
	}
	for name, want := range checks {
		g, ok := got[name]
		if !ok {
			t.Errorf("missing brick %q", name)
			continue
		}
		if g.Solid != want.solid || g.Liquid != want.liquid || g.Transparent != want.transparent || g.Collidable != want.collidable {
			t.Errorf("brick %q = %+v, want solid=%v liquid=%v transparent=%v collidable=%v",
				name, g, want.solid, want.liquid, want.transparent, want.collidable)
		}
	}
}
