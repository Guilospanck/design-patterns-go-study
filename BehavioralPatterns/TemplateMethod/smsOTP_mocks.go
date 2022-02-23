package main

type SMSOTPMock struct {
	OTP
}

func (sms *SMSOTPMock) generateRandomOTP(length int) string {
	return ""
}

func (sms *SMSOTPMock) getMessage(otp string) string {
	return ""
}

func (sms *SMSOTPMock) sendNotification(message string) error {
	return nil
}

func (sms *SMSOTPMock) publishMetrics() {

}

// mocks
func (sms *SMSOTPMock) getSMSOTPLength() int {
	return 4
}

func (sms *SMSOTPMock) getSMSMessage() string {
	return "[SMS] This is an OTP message"
}

func (sms *SMSOTPMock) getRandomOTP() string {
	return "1234556789"
}

func NewSMSOTPMock() *SMSOTPMock {
	return &SMSOTPMock{}
}
