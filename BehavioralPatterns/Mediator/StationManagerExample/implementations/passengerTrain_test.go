package implementations

import (
	"base/BehavioralPatterns/Mediator/StationManagerExample/mocks"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

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
