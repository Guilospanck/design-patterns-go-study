package main

import "fmt"

type FIFOStrategy struct{}

func (strategy *FIFOStrategy) Evict(*Cache) {
	fmt.Println("Removing using FIFO strategy...")
}

func NewFIFOStrategy() *FIFOStrategy {
	return &FIFOStrategy{}
}
