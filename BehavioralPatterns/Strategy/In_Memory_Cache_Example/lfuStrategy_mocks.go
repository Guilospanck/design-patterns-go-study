package main

type LFUStrategyMock struct{}

var LFUStrategyMockEvictCalled bool = false

func (strategy *LFUStrategyMock) Evict(ICache) {
	LFUStrategyMockEvictCalled = true
}

func NewLFUStrategyMock() *LFUStrategyMock {
	return &LFUStrategyMock{}
}
