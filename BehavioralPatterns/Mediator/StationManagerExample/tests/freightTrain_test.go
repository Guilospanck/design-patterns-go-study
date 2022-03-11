package tests

import (
	"base/BehavioralPatterns/Mediator/StationManagerExample/implementations"
	"base/BehavioralPatterns/Mediator/StationManagerExample/interfaces"
	"base/BehavioralPatterns/Mediator/StationManagerExample/mocks"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FreightTrainTestSuite struct {
	suite.Suite

	train    *implementations.FreightTrain
	mediator interfaces.IMediator
}

func (s *FreightTrainTestSuite) SetupSuite() {
	s.mediator = mocks.NewStationManagerMock()
	s.train = implementations.NewFreightTrain(s.mediator)
}

func (s *FreightTrainTestSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), nil)
}

func TestFreightTrainTestSuiteInit(t *testing.T) {
	suite.Run(t, new(FreightTrainTestSuite))
}

func (s *FreightTrainTestSuite) TestArrive() {
	// arrange

	// act
	s.train.Arrive()

	// assert
	require.Equal(s.T(), true, mocks.StationManagerCanArriveCalled)
}

func (s *FreightTrainTestSuite) TestDepart() {
	// arrange

	// act
	s.train.Depart()

	// assert
	require.Equal(s.T(), true, mocks.StationManagerNotifyAboutDepartureCalled)
}

func (s *FreightTrainTestSuite) TestPermitArrival() {
	// arrange

	// act
	s.train.PermitArrival()

	// assert
	require.Equal(s.T(), true, mocks.StationManagerCanArriveCalled)
}
