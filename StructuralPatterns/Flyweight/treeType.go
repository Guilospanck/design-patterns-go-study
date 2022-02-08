package main

import "fmt"

/*
  Whole Tree object:
    tree = {
      name
      color
      texture
      x
      y
    }
*/

// Intrinsic values (FLYWEIGHT)
type TreeType struct {
	name, color, texture string
}

func (t *TreeType) draw(canvas string, x int, y int) {
	fmt.Printf(
		"Drawing the tree %s, with color %s and texture %s on the canvas %s at the coordinates (%d, %d)\n==========================\n",
		t.name, t.color, t.texture, canvas, x, y)
}

func NewTreeType(name, color, texture string) *TreeType {
	return &TreeType{
		name:    name,
		color:   color,
		texture: texture,
	}
}
