package implementations

import (
	"base/BehavioralPatterns/Command/SendDataExample/domain"
	"fmt"
)

type Receiver struct{}

func (r *Receiver) ReceiveData(data domain.Data) {
	fmt.Printf("Received: %+v\n", data)
}

func NewReceiver() *Receiver {
	return &Receiver{}
}
