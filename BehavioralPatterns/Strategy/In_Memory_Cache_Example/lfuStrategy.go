package main

import "fmt"

type LFUStrategy struct{}

func (strategy *LFUStrategy) Evict(*Cache) {
	fmt.Println("Removing using LFU strategy...")
}

func NewLFUStrategy() *LFUStrategy {
	return &LFUStrategy{}
}
