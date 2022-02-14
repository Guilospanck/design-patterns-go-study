package main

import "fmt"

type LRUStrategy struct{}

func (strategy *LRUStrategy) Evict(*Cache) {
	fmt.Println("Removing using LRU strategy...")
}

func NewLRUStrategy() *LRUStrategy {
	return &LRUStrategy{}
}
