package main

type OTPMock struct {
	iOTP IOTP
}

func (mock *OTPMock) generateAndSendOTP(otpLength int) error {
	return nil
}

func (mock *OTPMock) saveOTPCache(otp string) {

}

func NewOTPMock(iotp IOTP) *OTPMock {
	return &OTPMock{
		iOTP: iotp,
	}
}
