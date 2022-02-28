package main

import "fmt"

type LFUStrategy struct{}

func (strategy *LFUStrategy) Evict(ICache) {
	fmt.Println("Removing using LFU strategy...")
}

func NewLFUStrategy() *LFUStrategy {
	return &LFUStrategy{}
}
