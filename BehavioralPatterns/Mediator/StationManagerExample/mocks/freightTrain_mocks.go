package mocks

import (
	interfaces "base/BehavioralPatterns/Mediator/StationManagerExample/interfaces"
)

var (
	FreightTrainArrived = false
	FreightTrainWaiting = false

	FreightTrainArriveFunctionCalled = false
	FreightTrainDepartFunctionCalled = false
)

type FreightTrainMocks struct {
	Mediator interfaces.IMediator
}

func (freightTrain *FreightTrainMocks) Arrive() {
	FreightTrainArriveFunctionCalled = true

	if !freightTrain.Mediator.CanArrive(freightTrain) {
		FreightTrainArrived = false
		FreightTrainWaiting = true
	}
	FreightTrainArrived = true
	FreightTrainWaiting = false
}

func (freightTrain *FreightTrainMocks) Depart() {
	FreightTrainDepartFunctionCalled = true
	freightTrain.Mediator.NotifyAboutDeparture()
}

func (freightTrain *FreightTrainMocks) PermitArrival() {
	freightTrain.Arrive()
}

func NewFreightTrainMocks(mediator interfaces.IMediator) *FreightTrainMocks {
	return &FreightTrainMocks{
		Mediator: mediator,
	}
}
