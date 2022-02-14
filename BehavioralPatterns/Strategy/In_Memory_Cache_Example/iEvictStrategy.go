package main

type IEvictStrategy interface {
	Evict(*Cache)
}
