package main

// Iterator
type IIterator interface {
	hasNext() bool
	getNext() *User
}
