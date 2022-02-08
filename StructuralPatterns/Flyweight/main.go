package main

func main() {
	// gets TreeFactory
	treeFactory := NewTreeFactory()

	// gets new Forest
	forest := NewForest(treeFactory)

	// plant some trees
	forest.plantTree(0, 0, "Tree 1", "Green", "Clean")
	forest.plantTree(1, 1, "Tree 2", "White", "Dirty")
	forest.plantTree(2, 2, "Tree 3", "Red", "Clean")
	forest.plantTree(3, 3, "Tree 1", "Green", "Clean") // reusing
	forest.plantTree(4, 4, "Tree 5", "Blue", "Clean")
	forest.plantTree(5, 5, "Tree 6", "Pink", "Dirty")
	forest.plantTree(6, 6, "Tree 1", "Green", "Clean") // reusing
	forest.plantTree(7, 7, "Tree 8", "Yellow", "Dirty")
	forest.plantTree(8, 8, "Tree 9", "Aqua", "Clean")
	forest.plantTree(9, 9, "Tree 1", "Green", "Clean") // reusing
	forest.plantTree(10, 10, "Tree 11", "Purple", "Clean")
	forest.plantTree(11, 11, "Tree 12", "Orange", "Dirty")

	// draw to canvases
	forest.drawTrees("canvas1")

}
