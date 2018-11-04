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
	Kind     Kind
	Property Property
}
