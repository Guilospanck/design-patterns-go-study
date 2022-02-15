package main

import "fmt"

type EmailOTP struct {
	OTP
}

func (email *EmailOTP) generateRandomOTP(length int) string {
	fmt.Println("[EMAIL] Generating random OTP...")
	randomOTP := "456487945"

	return randomOTP[:length]
}

func (email *EmailOTP) getMessage(otp string) string {
	fmt.Printf("[EMAIL] Getting message from OTP %s\n", otp)
	message := "[EMAIL] This is an OTP message"

	return message
}

func (email *EmailOTP) sendNotification(message string) error {
	fmt.Printf("[EMAIL] Notification sent with message %s\n", message)
	return nil
}

func (email *EmailOTP) publishMetrics() {
	fmt.Println("[EMAIL] Publishing metrics...")
}

func NewEmailOTP() *EmailOTP {
	return &EmailOTP{}
}
