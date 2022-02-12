package main

type IReceiver interface {
	ReceiveData(data Data)
}
