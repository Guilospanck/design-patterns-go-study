package main

import "fmt"

// Concrete component
type PassengerTrain struct {
	Mediator IMediator
}

func (passengerTrain *PassengerTrain) Arrive() {
	if !passengerTrain.Mediator.CanArrive(passengerTrain) {
		fmt.Println("[PASSENGER TRAIN]: Arrival blocked. Waiting.")
		return
	}
	fmt.Println("[PASSENGER TRAIN]: Arrived.")
}

func (passengerTrain *PassengerTrain) Depart() {
	fmt.Println("[PASSENGER TRAIN]: Leaving...")
	passengerTrain.Mediator.NotifyAboutDeparture()
}

func (passengerTrain *PassengerTrain) PermitArrival() {
	fmt.Println("[PASSENGER TRAIN]: Arrival permitted. Arriving...")
	passengerTrain.Arrive()
}

func NewPassengerTrain(mediator IMediator) *PassengerTrain {
	return &PassengerTrain{
		Mediator: mediator,
	}
}
