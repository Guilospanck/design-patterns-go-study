package main

func main() {
	stationManager := NewStationManager()

	passengerTrain := NewPassengerTrain(stationManager)
	freightTrain := NewFreightTrain(stationManager)

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()
}
