package main

// with Go 1.18 we can change it to work with generics
// type T User

// Iterator
type IIterator interface {
	hasNext() bool
	getNext() string
}
