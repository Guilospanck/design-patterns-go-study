package main

// Extrinsic values
type Tree struct {
	x, y int
	*TreeType
}

func (t *Tree) draw(canvas string) {
	t.TreeType.draw(canvas, t.x, t.y)
}

func NewTree(x, y int, treeType TreeType) *Tree {
	return &Tree{
		x:        x,
		y:        y,
		TreeType: &treeType,
	}
}
