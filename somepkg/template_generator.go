// +build generator

package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

// TileGen is a structure used to automatically generate the TileType
// definitions and the GameWorldGrid arrays.  This is used by the
// template to produce Go source code.
type TileGen struct {

	// Timestamp is the time at which the templates were
	// generated. (which is when the go generate command is
	// called)
	Timestamp string

	// Definitions contains the definitions that give meaning the
	// grid.txt file.  The definitions are used to generate a
	// 2-dimensional array of tiles.
	Definitions map[string]string
}

// Readme is the template for the generated README.md file.
const Readme = `
Generated Tile Types
======================================================================

This directory is an experiment with code generation.  The goal is to
contain the game world within simple text files, and to have the
source code generate itself from that.


The Handwritten Files:

1. template_generator.go
2. tile_types.txt
3. grid.txt



~~~
This README is automatically generated.

Time Generated:
{{.Timestamp}}
~~~



TileType Names
----------------------------------------------------------------------

The template generator is written in go, and is run when the go
generate command is called from another file.  The templates for the
generated code are found within template_generator.go.

First, the tile type definitions are processed, and tile_types.go is
generated.  This file contains enumerator constants that can be used
throughout the code base.  It is just a list of tile type names.


<details><summary>Open List of Tile Type Names</summary><blockquote>

{{range $key := .Definitions}}{{$key}} \
{{end}}
</blockquote></details>




Tile Types Definitions
----------------------------------------------------------------------


|       Symbol       |        Name        |
|--------------------|--------------------|
{{range $symbol, $name := .Definitions}}| {{$symbol}} | {{$name}} |
{{end}}

`

// CodeTemplate is the template for the automatically generated go
// source code.
const CodeTemplate = `
package somepkg

// TileType corresponds to the possible kinds of ground tiles in the
// game.  These enumerators are automatically generated from a text
// file that defines all of the possiblities.  Empty tiles are
// automatically included.
type TileType int

const (
	empty TileType = iota
	{{range $_, $name := .Definitions}}{{$name}}
	{{end}}
)
`

const (
	filenameTileTypes  = "tile_types.txt"
	filenameGrid       = "grid.txt"
	filenameCodeOutput = "types.go"
	filenameReadme     = "README.md"
)

var (
	tilegen = &TileGen{}
)

// reads the file into a string, then calls the parser
func processTileTypesFile() {
	b, err := ioutil.ReadFile(filenameTileTypes)
	if err != nil {
		log.Fatal(err)
	}
	tilegen.Definitions = parseTileDefs(string(b))
}

// The parser that scans line-by-line, converting each valid line into
func parseTileDefs(s string) map[string]string {

	scanner := bufio.NewScanner(strings.NewReader(s))
	m := make(map[string]string)
	line := ""

	for scanner.Scan() {
		line = scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		arr := strings.SplitN(line, " ", 2)

		if len(arr) != 2 {
			continue
		}

		m[arr[0]] = arr[1]
	}
	return m
}

func makeTheFile(templ, filename string) {
	t := template.Must(template.New("templ").Parse(templ))
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	t.Execute(file, tilegen)
}

func main() {
	tilegen.Timestamp = time.Now().Format(time.Stamp)
	processTileTypesFile()
	makeTheFile(Readme, filenameReadme)
	makeTheFile(CodeTemplate, filenameCodeOutput)

}
