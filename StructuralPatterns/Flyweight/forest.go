package main

type Forest struct {
	trees   []*Tree
	factory *TreeFactory
}

func (f *Forest) plantTree(x, y int, name, color, texture string) {
	treeType := f.factory.getTreeType(name, color, texture)
	tree := NewTree(x, y, *treeType)
	f.trees = append(f.trees, tree)
}

func (f *Forest) drawTrees(canvas string) {
	for _, tree := range f.trees {
		tree.draw(canvas)
	}
}

func NewForest(treeFactory *TreeFactory) *Forest {
	return &Forest{
		factory: treeFactory,
	}
}
