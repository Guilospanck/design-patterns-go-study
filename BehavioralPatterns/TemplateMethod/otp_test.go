package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type OTPSuite struct {
	suite.Suite

	mock      *OTPMock
	emailMock *EmailOTPMock

	otp  *OTP
	iOTP IOTP
}

func (s *OTPSuite) SetupSuite() {
	s.emailMock = NewEmailOTPMock()
	s.iOTP = s.emailMock
	s.mock = NewOTPMock(s.iOTP)

	s.otp = NewOTP(s.iOTP)
}

func (s *OTPSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), nil)
}

func TestOTPInit(t *testing.T) {
	suite.Run(t, new(OTPSuite))
}

func (s *OTPSuite) TestGenerateAndSendOTP() {
	// arrange
	lengthMock := s.emailMock.getEmailOTPLength()

	// act
	err := s.otp.generateAndSendOTP(lengthMock)

	// assert
	require.NoError(s.T(), err)
}
