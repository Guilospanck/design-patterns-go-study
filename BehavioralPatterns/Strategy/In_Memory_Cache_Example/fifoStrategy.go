package main

import "fmt"

type FIFOStrategy struct{}

func (strategy *FIFOStrategy) Evict(ICache) {
	fmt.Println("Removing using FIFO strategy...")
}

func NewFIFOStrategy() *FIFOStrategy {
	return &FIFOStrategy{}
}
