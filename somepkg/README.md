
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
Sep 28 01:52:06
~~~



TileType Names
----------------------------------------------------------------------

The template generator is written in go, and is run when the go
generate command is called from another file.  The templates for the
generated code are found within template_generator.go.

First, the tile type definitions are processed, and tile_types.go is
generated.  This file contains enumerator constants that can be used
throughout the code base.  It is just a list of tile type names.


<details><summary>Open List of Tile Type Names</summary>blockquote>

grass \
dirt \
bush \
tree \
water \


</blockquote></details>




Tile Types Definitions
----------------------------------------------------------------------


|--------------------|--------------------|
|       Symbol       |        Name        |
|--------------------|--------------------|
| , | grass |
| _ | dirt |
| o | bush |
| t | tree |
| ~ | water |


 |


