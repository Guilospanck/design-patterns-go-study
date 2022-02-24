package main

import "fmt"

type EmailOTP struct {
	OTP
	randomOTP, message string
}

func (email *EmailOTP) generateRandomOTP(length int) string {
	fmt.Println("[EMAIL] Generating random OTP...")

	return email.randomOTP[:length]
}

func (email *EmailOTP) getMessage(otp string) string {
	fmt.Printf("[EMAIL] Getting message from OTP %s\n", otp)

	return email.message
}

func (email *EmailOTP) sendNotification(message string) error {
	fmt.Printf("[EMAIL] Notification sent with message %s\n", message)
	return nil
}

func (email *EmailOTP) publishMetrics() {
	fmt.Println("[EMAIL] Publishing metrics...")
}

func NewEmailOTP(randomOTP, message string) *EmailOTP {
	return &EmailOTP{
		randomOTP: randomOTP,
		message:   message,
	}
}
