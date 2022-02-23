package main

import "fmt"

type SMSOTP struct {
	OTP
	randomOTP string
	message   string
}

func (sms *SMSOTP) generateRandomOTP(length int) string {
	fmt.Println("[SMS] Generating random OTP...")

	return sms.randomOTP[:length]
}

func (sms *SMSOTP) getMessage(otp string) string {
	fmt.Printf("[SMS] Getting message from OTP %s\n", otp)

	return sms.message
}

func (sms *SMSOTP) sendNotification(message string) error {
	fmt.Printf("[SMS] Notification sent with message %s\n", message)
	return nil
}

func (sms *SMSOTP) publishMetrics() {
	fmt.Println("[SMS] Publishing metrics...")
}

func NewSMSOTP(message, randomOTP string) *SMSOTP {
	return &SMSOTP{
		randomOTP: randomOTP,
		message:   message,
	}
}
