package main

type FIFOStrategyMocks struct{}

var FifoStrategyEvictCalled bool = false

func (strategy *FIFOStrategyMocks) Evict(ICache) {
	FifoStrategyEvictCalled = true
}

func NewFIFOStrategyMocks() *FIFOStrategyMocks {
	return &FIFOStrategyMocks{}
}
