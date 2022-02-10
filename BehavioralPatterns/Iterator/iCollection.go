package main

// Collection
type ICollection interface {
	createIterator() IIterator
}
