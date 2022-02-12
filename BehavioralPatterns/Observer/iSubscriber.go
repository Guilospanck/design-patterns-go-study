package main

type ISubscriber interface {
	Update(string)
	GetID() string
}
