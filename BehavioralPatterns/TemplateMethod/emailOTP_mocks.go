package main

type EmailOTPMock struct {
	OTPMock
}

func (sms *EmailOTPMock) generateRandomOTP(length int) string {
	return ""
}

func (sms *EmailOTPMock) getMessage(otp string) string {
	return ""
}

func (sms *EmailOTPMock) sendNotification(message string) error {
	return nil
}

func (sms *EmailOTPMock) publishMetrics() {

}

// mocks
func (sms *EmailOTPMock) getEmailOTPLength() int {
	return 2
}

func (sms *EmailOTPMock) getEmailMessage() string {
	return "[EMAIL] This is an OTP message"
}

func (sms *EmailOTPMock) getRandomOTP() string {
	return "456487945"
}

func NewEmailOTPMock() *EmailOTPMock {
	return &EmailOTPMock{}
}
