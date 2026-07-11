package proto

// BrickType is the canonical palette. Drift here is the most dangerous;
// keep reconciled with initial-repo-spec.md.
type BrickType uint8

const (
	BrickAir BrickType = iota
	BrickRock
	BrickSoil
	BrickGrass
	BrickWood
	BrickWater
)

// BrickDef describes physical properties for a brick type.
type BrickDef struct {
	Type        BrickType `json:"type"`
	Name        string    `json:"name"`
	Solid       bool      `json:"solid"`
	Liquid      bool      `json:"liquid"`
	Transparent bool      `json:"transparent"`
	Collidable  bool      `json:"collidable"`
}

// DefaultPalette is the starter brick set (expandable).
func DefaultPalette() []BrickDef {
	return []BrickDef{
		{Type: BrickAir, Name: "air", Solid: false, Liquid: false, Transparent: true, Collidable: false},
		{Type: BrickRock, Name: "rock", Solid: true, Liquid: false, Transparent: false, Collidable: true},
		{Type: BrickSoil, Name: "soil", Solid: true, Liquid: false, Transparent: false, Collidable: true},
		{Type: BrickGrass, Name: "grass", Solid: true, Liquid: false, Transparent: false, Collidable: true},
		{Type: BrickWood, Name: "wood", Solid: true, Liquid: false, Transparent: false, Collidable: true},
		{Type: BrickWater, Name: "water", Solid: false, Liquid: true, Transparent: true, Collidable: false},
	}
}
