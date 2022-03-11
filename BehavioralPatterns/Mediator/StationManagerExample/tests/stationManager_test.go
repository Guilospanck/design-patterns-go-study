package tests

import (
	"base/BehavioralPatterns/Mediator/StationManagerExample/implementations"
	"base/BehavioralPatterns/Mediator/StationManagerExample/interfaces"
	"base/BehavioralPatterns/Mediator/StationManagerExample/mocks"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type StationManagerTestSuite struct {
	suite.Suite

	station        *implementations.StationManager
	freightTrain   interfaces.ITrain
	passengerTrain interfaces.ITrain
}

func (s *StationManagerTestSuite) SetupSuite() {
	s.station = implementations.NewStationManager()
	s.freightTrain = mocks.NewFreightTrainMocks(s.station)
	s.passengerTrain = mocks.NewPassengerTrainMocks(s.station)
}

func (s *StationManagerTestSuite) AfterTest(_, _ string) {
	s.SetupSuite()
	require.NoError(s.T(), nil)
}

func TestStationManagerTestSuiteInit(t *testing.T) {
	suite.Run(t, new(StationManagerTestSuite))
}

func (s *StationManagerTestSuite) TestCanArrive() {
	// arrange

	// act and assert
	result := s.station.CanArrive(s.freightTrain)

	require.Equal(s.T(), 1, len(s.station.TrainQueue))
	require.Equal(s.T(), false, result)
	require.Equal(s.T(), false, s.station.IsPlatformFree)

	result = s.station.CanArrive(s.freightTrain)

	require.Equal(s.T(), 2, len(s.station.TrainQueue))
	require.Equal(s.T(), false, result)
	require.Equal(s.T(), false, s.station.IsPlatformFree)
}

func (s *StationManagerTestSuite) TestNotifyAboutDeparture() {
	// arrange

	// act and assert
	s.station.NotifyAboutDeparture()

	require.Equal(s.T(), true, s.station.IsPlatformFree)
	require.Equal(s.T(), 0, len(s.station.TrainQueue))

	s.station.TrainQueue = append(s.station.TrainQueue, s.freightTrain)
	require.Equal(s.T(), 1, len(s.station.TrainQueue))
	s.station.TrainQueue = append(s.station.TrainQueue, s.passengerTrain)
	require.Equal(s.T(), 2, len(s.station.TrainQueue))

	s.station.NotifyAboutDeparture()
	require.Equal(s.T(), false, s.station.IsPlatformFree) // passenger train will call CanArrive and then "IsPlatformFree" will be false
	require.Equal(s.T(), 1, len(s.station.TrainQueue))

	require.Equal(s.T(), true, mocks.PassengerTrainArriveFunctionCalled)
	require.Equal(s.T(), true, mocks.PassengerTrainArrived)
	require.Equal(s.T(), false, mocks.PassengerTrainWaiting)
}
