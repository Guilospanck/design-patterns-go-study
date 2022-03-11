package mocks

import (
	interfaces "base/BehavioralPatterns/Mediator/StationManagerExample/interfaces"
)

type StationManagerMock struct {
	IsPlatformFree bool
	TrainQueue     []interfaces.ITrain
}

func NewStationManagerMock() *StationManagerMock {
	return &StationManagerMock{
		IsPlatformFree: false,
		TrainQueue:     make([]interfaces.ITrain, 0),
	}
}
