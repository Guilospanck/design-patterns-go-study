package main

type ISubscriber interface {
	Update(Data)
	GetID() string
}
