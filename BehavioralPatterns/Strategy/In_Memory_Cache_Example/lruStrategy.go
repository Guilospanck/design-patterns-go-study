package main

import "fmt"

type LRUStrategy struct{}

func (strategy *LRUStrategy) Evict(ICache) {
	fmt.Println("Removing using LRU strategy...")
}

func NewLRUStrategy() *LRUStrategy {
	return &LRUStrategy{}
}
