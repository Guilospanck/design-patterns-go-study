package main

import "fmt"

// Decides whether to re-use existing flyweight of to create a new object
type TreeFactory struct {
	treeTypes []*TreeType
}

func (factory *TreeFactory) getTreeType(name, color, texture string) *TreeType {

	treeType, err := factory._doesTreeTypeExist(name, color, texture)
	if err == nil {
		fmt.Println("Returning a cached flyweight...")
		return treeType
	}

	fmt.Println("Creating a new flyweight...")
	newTreeType := NewTreeType(name, color, texture)
	factory.treeTypes = append(factory.treeTypes, newTreeType)

	return newTreeType
}

func (factory *TreeFactory) _doesTreeTypeExist(name, color, texture string) (*TreeType, error) {
	for _, treeType := range factory.treeTypes {
		if treeType.name == name && treeType.color == color && treeType.texture == texture {
			return treeType, nil
		}
	}

	return nil, fmt.Errorf("there is no treeType")
}

func NewTreeFactory() *TreeFactory {
	return &TreeFactory{}
}
