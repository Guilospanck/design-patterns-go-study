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
