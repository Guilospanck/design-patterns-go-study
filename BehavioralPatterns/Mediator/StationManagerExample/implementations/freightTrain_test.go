package implementations

import (
	"base/BehavioralPatterns/Mediator/StationManagerExample/mocks"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

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
