package tests

import (
	"base/BehavioralPatterns/Mediator/StationManagerExample/implementations"
	"base/BehavioralPatterns/Mediator/StationManagerExample/interfaces"
	"base/BehavioralPatterns/Mediator/StationManagerExample/mocks"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type PassengerTrainTestSuite struct {
	suite.Suite

	train    *implementations.PassengerTrain
	mediator interfaces.IMediator
}

func (s *PassengerTrainTestSuite) SetupSuite() {
	s.mediator = mocks.NewStationManagerMock()
	s.train = implementations.NewPassengerTrain(s.mediator)
}

func (s *PassengerTrainTestSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), nil)
}

func TestPassengerTrainTestSuiteInit(t *testing.T) {
	suite.Run(t, new(PassengerTrainTestSuite))
}

func (s *PassengerTrainTestSuite) TestArrive() {
	// arrange

	// act
	s.train.Arrive()

	// assert
	require.Equal(s.T(), true, mocks.StationManagerCanArriveCalled)
}

func (s *PassengerTrainTestSuite) TestDepart() {
	// arrange

	// act
	s.train.Depart()

	// assert
	require.Equal(s.T(), true, mocks.StationManagerNotifyAboutDepartureCalled)
}

func (s *PassengerTrainTestSuite) TestPermitArrival() {
	// arrange

	// act
	s.train.PermitArrival()

	// assert
	require.Equal(s.T(), true, mocks.StationManagerCanArriveCalled)
}
