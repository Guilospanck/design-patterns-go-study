package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CacheSuite struct {
	suite.Suite

	mock     *CacheMocks
	cache    *Cache
	strategy IEvictStrategy
}

func (s *CacheSuite) SetupSuite() {
	s.strategy = NewFIFOStrategyMocks()
	s.mock = NewCacheMocks(s.strategy)
	s.cache = NewCache(s.strategy)
}

func (s *CacheSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), nil)
}

func TestCacheInit(t *testing.T) {
	suite.Run(t, new(CacheSuite))
}

func (s *CacheSuite) TestSetEvictStrategy() {
	// arrange
	s.strategy = NewLFUStrategyMock()

	// act
	s.cache.SetEvictStrategy(s.strategy)

	// assert
	require.Equal(s.T(), s.strategy, s.cache.evictStrategy)
}

func (s *CacheSuite) TestSetMaxCapacity() {
	// arrange
	capacityMock := 5

	// act
	s.cache.SetMaxCapacity(capacityMock)

	// assert
	require.Equal(s.T(), capacityMock, s.cache.maxCapacity)
}

func (s *CacheSuite) TestAdd() {
	// arrange
	keyMock := "key1"
	valueMock := "value1"

	// act
	s.cache.Add(keyMock, valueMock)

	// assert
	require.Equal(s.T(), 1, len(s.cache.storage))
	require.Equal(s.T(), false, FifoStrategyEvictCalled)
}

func (s *CacheSuite) TestMultipleAdd() {
	// arrange
	key1 := "key1"
	key2 := "key2"

	// act
	s.cache.Add(key1, MapMocks[key1])
	s.cache.Add(key2, MapMocks[key2])

	// assert
	require.Equal(s.T(), 2, len(s.cache.storage))
	require.Equal(s.T(), false, FifoStrategyEvictCalled)
}

func (s *CacheSuite) TestMultipleAddWithEvict() {
	// arrange
	key1 := "key1"
	key2 := "key2"
	key3 := "key3"
	maxCapacity := s.cache.maxCapacity

	// act
	s.cache.Add(key1, MapMocks[key1])
	s.cache.Add(key2, MapMocks[key2])
	s.cache.Add(key3, MapMocks[key3])

	// assert
	require.Equal(s.T(), maxCapacity, len(s.cache.storage))
	require.Equal(s.T(), true, FifoStrategyEvictCalled)
}

func (s *CacheSuite) TestPrintCache() {
	// arrange

	// act
	s.cache.PrintCache()

	// assert
	require.NoError(s.T(), nil)
}
