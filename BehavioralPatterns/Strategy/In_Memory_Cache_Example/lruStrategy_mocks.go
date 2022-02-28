package main

type LRUStrategyMock struct{}

var LRUStrategyMockEvictCalled bool = false

func (strategy *LRUStrategyMock) Evict(ICache) {
	LRUStrategyMockEvictCalled = true
}

func NewLRUStrategyMock() *LRUStrategyMock {
	return &LRUStrategyMock{}
}
