package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Context
type Cache struct {
	evictStrategy IEvictStrategy
	storage       map[string]string
	maxCapacity   int
}

func (c *Cache) SetEvictStrategy(evictStrategy IEvictStrategy) {
	c.evictStrategy = evictStrategy
}

func (c *Cache) SetMaxCapacity(maxCapacity int) {
	c.maxCapacity = maxCapacity
}

func (c *Cache) Add(key, value string) {
	if len(c.storage) >= c.maxCapacity {
		c.Evict(key)
	}
	c.storage[key] = value
}

func (c *Cache) PrintCache() {
	for key, value := range c.storage {
		fmt.Printf("[%s]: [%s]\n", key, value)
	}
	fmt.Println("================================")
}

func (c *Cache) Evict(key string) {
	c.evictStrategy.Evict(c)

	index := strings.Split(key, "key")
	indexInt, _ := strconv.Atoi(index[1])
	indexToBeDeleted := indexInt - 2
	keyToBeDeleted := fmt.Sprintf("key%d", indexToBeDeleted)
	delete(c.storage, keyToBeDeleted)
}

func NewCache(evictStrategy IEvictStrategy) *Cache {
	return &Cache{
		evictStrategy: evictStrategy,
		storage:       make(map[string]string),
		maxCapacity:   2,
	}
}
