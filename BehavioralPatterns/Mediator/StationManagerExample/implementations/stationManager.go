package implementations

import "base/BehavioralPatterns/Mediator/StationManagerExample/interfaces"

// Concrete mediator
type StationManager struct {
	IsPlatformFree bool
	TrainQueue     []interfaces.ITrain
}

func (stationManager *StationManager) CanArrive(train interfaces.ITrain) bool {
	if stationManager.IsPlatformFree {
		stationManager.IsPlatformFree = false
		return true
	}
	stationManager.TrainQueue = append(stationManager.TrainQueue, train)
	return false
}

func (stationManager *StationManager) NotifyAboutDeparture() {
	stationManager.IsPlatformFree = true

	if len(stationManager.TrainQueue) > 0 {
		stationManager.TrainQueue = stationManager.TrainQueue[1:]

		if len(stationManager.TrainQueue) > 0 {
			firstTrainInQueue := stationManager.TrainQueue[0]
			firstTrainInQueue.PermitArrival()
		}
	}
}

func NewStationManager() *StationManager {
	return &StationManager{}
}
