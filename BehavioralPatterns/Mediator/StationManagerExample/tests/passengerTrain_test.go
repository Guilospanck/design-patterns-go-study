package tests

// import (
// 	"base/BehavioralPatterns/Mediator/StationManagerExample/implementations"
// 	"base/BehavioralPatterns/Mediator/StationManagerExample/interfaces"
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// 	"github.com/stretchr/testify/suite"
// )

// type PassengerTrainTestSuite struct {
// 	suite.Suite

// 	train    *implementations.PassengerTrain
// 	mediator interfaces.IMediator
// }

// func (s *PassengerTrainTestSuite) SetupSuite() {

// 	s.train = implementations.NewPassengerTrain()
// }

// func (s *PassengerTrainTestSuite) AfterTest(_, _ string) {
// 	require.NoError(s.T(), nil)
// }

// func TestPassengerTrainTestSuiteInit(t *testing.T) {
// 	suite.Run(t, new(PassengerTrainTestSuite))
// }

// func (passengerTrain *PassengerTrainTestSuite) Arrive() {
// 	if !passengerTrain.Mediator.CanArrive(passengerTrain) {
// 		fmt.Println("[PASSENGER TRAIN]: Arrival blocked. Waiting.")
// 		return
// 	}
// 	fmt.Println("[PASSENGER TRAIN]: Arrived.")
// }

// func (passengerTrain *PassengerTrainTestSuite) Depart() {
// 	fmt.Println("[PASSENGER TRAIN]: Leaving...")
// 	passengerTrain.Mediator.NotifyAboutDeparture()
// }

// func (passengerTrain *PassengerTrainTestSuite) PermitArrival() {
// 	fmt.Println("[PASSENGER TRAIN]: Arrival permitted. Arriving...")
// 	passengerTrain.Arrive()
// }
