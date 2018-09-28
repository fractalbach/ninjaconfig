
Generated Tile Types
======================================================================

This directory is an experiment with code generation.  The goal is to
contain the game world within simple text files, and to have the
source code generate itself from that.


The Handwritten Files:

1. template_generator.go
2. tile_types.txt
3. grid.txt



This README was automatically generated.

Time Generated |
---------------|
Fri Sep 28 10:09:27 UTC 2018 |


Tile Types Definitions
----------------------------------------------------------------------

This defintion table is generated from tile_types.txt.  Each symbol
corresponds to the Name, which becomes a constant in types.go.  Since
the following symbols have been defined, they can be used in grid.txt
to create game worlds!

Symbol | Name 
-------|------
, | grass
_ | dirt
o | bush
t | tree
~ | water


