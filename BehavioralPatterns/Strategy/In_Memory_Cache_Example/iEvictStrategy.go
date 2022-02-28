package main

type IEvictStrategy interface {
	Evict(ICache)
}
