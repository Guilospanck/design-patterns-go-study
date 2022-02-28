package main

import (
	"fmt"
	"strconv"
	"strings"
)

type CacheMocks struct {
	evictStrategy IEvictStrategy
	storage       map[string]string
	maxCapacity   int
}

func (c *CacheMocks) SetEvictStrategy(evictStrategy IEvictStrategy) {
	c.evictStrategy = evictStrategy
}

func (c *CacheMocks) SetMaxCapacity(maxCapacity int) {
	c.maxCapacity = maxCapacity
}

func (c *CacheMocks) Add(key, value string) {
	if len(c.storage) >= c.maxCapacity {
		c.Evict(key)
	}
	c.storage[key] = value
}

var PrintCacheCalled bool = false

func (c *CacheMocks) PrintCache() {
	PrintCacheCalled = true
}

func (c *CacheMocks) Evict(key string) {
	c.evictStrategy.Evict(c)

	index := strings.Split(key, "key")
	indexInt, _ := strconv.Atoi(index[1])
	indexToBeDeleted := indexInt - 2
	keyToBeDeleted := fmt.Sprintf("key%d", indexToBeDeleted)
	delete(c.storage, keyToBeDeleted)
}

func NewCacheMocks(evictStrategy IEvictStrategy) *CacheMocks {
	return &CacheMocks{
		evictStrategy: evictStrategy,
		storage:       make(map[string]string),
		maxCapacity:   2,
	}
}

var MapMocks = map[string]string{
	"key1": "value1",
	"key2": "value2",
	"key3": "value3",
	"key4": "value4",
	"key5": "value5",
}
