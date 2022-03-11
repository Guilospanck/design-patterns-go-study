package mocks

import (
	interfaces "base/BehavioralPatterns/Mediator/StationManagerExample/interfaces"
)

type StationManagerMock struct {
}

var (
	StationManagerCanArriveCalled            = false
	StationManagerNotifyAboutDepartureCalled = false
)

func (stationManager *StationManagerMock) CanArrive(train interfaces.ITrain) bool {
	StationManagerCanArriveCalled = true

	return true
}

func (stationManager *StationManagerMock) NotifyAboutDeparture() {
	StationManagerNotifyAboutDepartureCalled = true
}

func NewStationManagerMock() *StationManagerMock {
	return &StationManagerMock{}
}
