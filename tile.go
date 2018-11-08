//go:generate go run src/types_generator.go
//go:generate go fmt

package tiledefs

// Kind corresponds to the possible kinds of ground tiles in the
// game.  These enumerators are automatically generated from a text
// file that defines all of the possiblities.
type Kind int

// Property is a bitmask of properties that can be applied to tiles.
// Each tile can have either no properties (Property = 0), or multiple
// properties.
type Property int

// Tile is a ground tile on the world map.
type Tile struct {
	Symbol   string
	Kind     Kind
	Property Property
}

func (t Tile) String() string {
	return t.Symbol
}

func DecodeSymbol(s string) (Tile, bool) {
	t := Tile{
		Symbol: s,
	}
	kind, ok := SymbolToKind[s]
	if !ok {
		return t, false
	}
	prop, ok2 := DefaultProperty[kind]
	if !ok2 {
		return t, false
	}
	t.Kind = kind
	t.Property = prop
	return t, true
}
