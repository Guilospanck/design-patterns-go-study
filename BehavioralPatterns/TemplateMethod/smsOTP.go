package main

import "fmt"

type SMSOTP struct {
	OTP
}

func (sms *SMSOTP) generateRandomOTP(length int) string {
	fmt.Println("[SMS] Generating random OTP...")
	randomOTP := "1234556789"

	return randomOTP[:length]
}

func (sms *SMSOTP) getMessage(otp string) string {
	fmt.Printf("[SMS] Getting message from OTP %s\n", otp)
	message := "[SMS] This is an OTP message"

	return message
}

func (sms *SMSOTP) sendNotification(message string) error {
	fmt.Printf("[SMS] Notification sent with message %s\n", message)
	return nil
}

func (sms *SMSOTP) publishMetrics() {
	fmt.Println("[SMS] Publishing metrics...")
}

func NewSMSOTP() *SMSOTP {
	return &SMSOTP{}
}
