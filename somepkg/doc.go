//go:generate go run -tags='generator' types_generator.go
//go:generate go fmt

// Package somepkg contains type definitions for the world's Tiles.
// The initial configuration of the world map will also be found here.
// Most of the code here is automatically generated from plain-text
// files.
//
// When a game world is generated, each tile has certain default
// properties, which can be found here.
package somepkg

// Kind corresponds to the possible kinds of ground tiles in the
// game.  These enumerators are automatically generated from a text
// file that defines all of the possiblities.
type Kind int

// Property is a bitmask of properties that can be applied to tiles.
// Each tile can have either no properties (Property = 0), or multiple
// properties.
type Property int
