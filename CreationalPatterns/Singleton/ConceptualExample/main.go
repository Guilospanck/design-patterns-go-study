package main

import (
	"base/CreationalPatterns/Singleton/ConceptualExample/implementations"
	"fmt"
)

func main() {
	for i := 0; i < 30; i++ {
		go implementations.GetInstance()
	}

	fmt.Scanln()
}
