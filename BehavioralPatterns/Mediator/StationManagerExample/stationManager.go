package main

// Concrete mediator
type StationManager struct {
	IsPlatformFree bool
	TrainQueue     []ITrain
}

func (stationManager *StationManager) CanArrive(train ITrain) bool {
	if stationManager.IsPlatformFree {
		stationManager.IsPlatformFree = false
		return true
	}
	stationManager.TrainQueue = append(stationManager.TrainQueue, train)
	return false
}

func (stationManager *StationManager) NotifyAboutDeparture() {
	if !stationManager.IsPlatformFree {
		stationManager.IsPlatformFree = true
	}
	if len(stationManager.TrainQueue) > 0 {
		firstTrainInQueue := stationManager.TrainQueue[0]
		stationManager.TrainQueue = stationManager.TrainQueue[1:]
		firstTrainInQueue.PermitArrival()
	}
}

func NewStationManager() *StationManager {
	return &StationManager{}
}
