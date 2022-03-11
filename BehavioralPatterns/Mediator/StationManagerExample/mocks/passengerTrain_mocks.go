package mocks

import (
	interfaces "base/BehavioralPatterns/Mediator/StationManagerExample/interfaces"
)

var (
	PassengerTrainArrived = false
	PassengerTrainWaiting = false

	PassengerTrainArriveFunctionCalled = false
	PassengerTrainDepartFunctionCalled = false
)

type PassengerTrainMocks struct {
	Mediator interfaces.IMediator
}

func (passengerTrain *PassengerTrainMocks) Arrive() {
	PassengerTrainArriveFunctionCalled = true

	if !passengerTrain.Mediator.CanArrive(passengerTrain) {
		PassengerTrainArrived = false
		PassengerTrainWaiting = true
	}
	PassengerTrainArrived = true
	PassengerTrainWaiting = false
}

func (passengerTrain *PassengerTrainMocks) Depart() {
	PassengerTrainDepartFunctionCalled = true
	passengerTrain.Mediator.NotifyAboutDeparture()
}

func (passengerTrain *PassengerTrainMocks) PermitArrival() {
	passengerTrain.Arrive()
}

func NewPassengerTrainMocks(mediator interfaces.IMediator) *PassengerTrainMocks {
	return &PassengerTrainMocks{
		Mediator: mediator,
	}
}
