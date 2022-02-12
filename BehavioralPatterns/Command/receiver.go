package main

import "fmt"

type Receiver struct{}

func (r *Receiver) ReceiveData(data Data) {
	fmt.Printf("Received: %+v\n", data)
}

func NewReceiver() *Receiver {
	return &Receiver{}
}
