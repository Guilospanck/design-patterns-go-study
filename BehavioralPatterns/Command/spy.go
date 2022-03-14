package main

import "github.com/stretchr/testify/mock"

type SendDataCommandSpy struct {
	mock.Mock
}

func (spy *SendDataCommandSpy) Execute() {
	spy.Called()
}

func NewSendDataCommandSpy() *SendDataCommandSpy {
	return new(SendDataCommandSpy)
}

//
type ReceiverSpy struct {
	mock.Mock
}

func (spy *ReceiverSpy) ReceiveData(data Data) {
	spy.Called(data)
}

func NewReceiverSpy() *ReceiverSpy {
	return new(ReceiverSpy)
}
