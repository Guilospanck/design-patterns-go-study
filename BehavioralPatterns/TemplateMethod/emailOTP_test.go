package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type EmailOTPSuite struct {
	suite.Suite

	mock  *EmailOTPMock
	email *EmailOTP
}

func (s *EmailOTPSuite) SetupSuite() {
	s.mock = NewEmailOTPMock()

	message := s.mock.getEmailMessage()
	randomOTP := s.mock.getRandomOTP()

	s.email = NewEmailOTP(randomOTP, message)
}

func (s *EmailOTPSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), nil)
}

func TestEmailOTPInit(t *testing.T) {
	suite.Run(t, new(EmailOTPSuite))
}

func (s *EmailOTPSuite) TestGenerateRandomOTP() {
	// arrange
	length := s.mock.getEmailOTPLength()
	randomOTPMock := s.mock.getRandomOTP()[:length]

	// act
	randomOTP := s.email.generateRandomOTP(length)

	// assert
	require.Equal(s.T(), randomOTP, randomOTPMock)
}

func (s *EmailOTPSuite) TestGetMessage() {
	// arrange
	length := s.mock.getEmailOTPLength()
	randomOTPMock := s.mock.getRandomOTP()[:length]
	messageMock := s.mock.getEmailMessage()

	// act
	message := s.email.getMessage(randomOTPMock)

	// assert
	require.Equal(s.T(), message, messageMock)
}

func (s *EmailOTPSuite) TestSendNotification() {
	// arrange
	messageMock := s.mock.getEmailMessage()

	// act
	err := s.email.sendNotification(messageMock)

	// assert
	require.NoError(s.T(), err)
}
