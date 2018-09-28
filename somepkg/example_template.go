// +build ignore

package somepkg

// TileType corresponds to the possible kinds of ground tiles in the
// game.  These enumerators are automatically generated from a text
// file that defines all of the possiblities.  Empty tiles are
// automatically included.
type TileType int

const (
	empty TileType = iota
	{{range .TileTypeNameList}}
	{{.}}
	{{end}}
)
