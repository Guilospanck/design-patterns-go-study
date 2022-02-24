package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SMSOTPSuite struct {
	suite.Suite

	mock   *SMSOTPMock
	smsOTP *SMSOTP
}

func (s *SMSOTPSuite) SetupSuite() {
	s.mock = NewSMSOTPMock()
	s.smsOTP = NewSMSOTP(s.mock.getSMSMessage(), s.mock.getRandomOTP())
}

func (s *SMSOTPSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), nil)
}

func TestSMSOTPInit(t *testing.T) {
	suite.Run(t, new(SMSOTPSuite))
}

func (s *SMSOTPSuite) TestGenerateRandomOTP() {
	// arrange
	length := s.mock.getSMSOTPLength()
	mockRandomOTP := s.mock.getRandomOTP()[:length]

	// act
	randomOTP := s.smsOTP.generateRandomOTP(length)

	// assert
	require.Equal(s.T(), randomOTP, mockRandomOTP)
}

func (s *SMSOTPSuite) TestGetMessage() {
	// arrange
	length := s.mock.getSMSOTPLength()
	mockRandomOTP := s.mock.getRandomOTP()[:length]
	mockMessage := s.mock.getSMSMessage()

	// act
	message := s.smsOTP.getMessage(mockRandomOTP)

	// assert
	require.Equal(s.T(), message, mockMessage)
}

func (s *SMSOTPSuite) TestSendNotification() {
	// arrange
	messageMock := s.mock.getSMSMessage()

	// act
	err := s.smsOTP.sendNotification(messageMock)

	// assert
	require.NoError(s.T(), err)
}
