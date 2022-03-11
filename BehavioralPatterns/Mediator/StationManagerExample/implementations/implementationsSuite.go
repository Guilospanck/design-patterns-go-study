package implementations

import (
	"base/BehavioralPatterns/Mediator/StationManagerExample/interfaces"
	"base/BehavioralPatterns/Mediator/StationManagerExample/mocks"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FreightTrainTestSuite struct {
	suite.Suite

	train    *FreightTrain
	mediator interfaces.IMediator
}

func (s *FreightTrainTestSuite) SetupSuite() {
	s.mediator = mocks.NewStationManagerMock()
	s.train = NewFreightTrain(s.mediator)
}

func (s *FreightTrainTestSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), nil)
}

type PassengerTrainTestSuite struct {
	suite.Suite

	train    *PassengerTrain
	mediator interfaces.IMediator
}

func (s *PassengerTrainTestSuite) SetupSuite() {
	s.mediator = mocks.NewStationManagerMock()
	s.train = NewPassengerTrain(s.mediator)
}

func (s *PassengerTrainTestSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), nil)
}

type StationManagerTestSuite struct {
	suite.Suite

	station        *StationManager
	freightTrain   interfaces.ITrain
	passengerTrain interfaces.ITrain
}

func (s *StationManagerTestSuite) SetupSuite() {
	s.station = NewStationManager()
	s.freightTrain = mocks.NewFreightTrainMocks(s.station)
	s.passengerTrain = mocks.NewPassengerTrainMocks(s.station)
}

func (s *StationManagerTestSuite) AfterTest(_, _ string) {
	s.SetupSuite()
	require.NoError(s.T(), nil)
}
