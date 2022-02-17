package main

import "fmt"

// concrete component
type FreightTrain struct {
	Mediator IMediator
}

func (freightTrain *FreightTrain) Arrive() {
	if !freightTrain.Mediator.CanArrive(freightTrain) {
		fmt.Println("[FREIGHT TRAIN]: Arrival blocked. Waiting.")
		return
	}
	fmt.Println("[FREIGHT TRAIN]: Arrived.")
}

func (freightTrain *FreightTrain) Depart() {
	fmt.Println("[FREIGHT TRAIN]: Leaving...")
	freightTrain.Mediator.NotifyAboutDeparture()
}

func (freightTrain *FreightTrain) PermitArrival() {
	fmt.Println("[FREIGHT TRAIN]: Arrival permitted. Arriving...")
	freightTrain.Arrive()
}

func NewFreightTrain(mediator IMediator) *FreightTrain {
	return &FreightTrain{
		Mediator: mediator,
	}
}
