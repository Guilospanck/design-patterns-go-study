package main

import (
	implementations "base/BehavioralPatterns/Mediator/StationManagerExample/implementations"
)

func main() {
	stationManager := implementations.NewStationManager()

	passengerTrain := implementations.NewPassengerTrain(stationManager)
	freightTrain := implementations.NewFreightTrain(stationManager)

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()
}
