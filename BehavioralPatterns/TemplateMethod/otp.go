package main

import "fmt"

type IOTP interface {
	generateRandomOTP(length int) string
	saveOTPCache(otp string)
	getMessage(otp string) string
	sendNotification(message string) error
	publishMetrics()
}

// embedding a interface into a struct is Go way of creating a "abstract class"
type OTP struct {
	iOTP IOTP
}

// skeleton method
func (o *OTP) generateAndSendOTP(otpLength int) error {
	otp := o.iOTP.generateRandomOTP(otpLength)
	o.iOTP.saveOTPCache(otp) // default method
	message := o.iOTP.getMessage(otp)
	err := o.iOTP.sendNotification(message)
	if err != nil {
		return err
	}
	o.iOTP.publishMetrics()
	return nil
}

// default method
func (o *OTP) saveOTPCache(otp string) {
	fmt.Printf("[DEFAULT METHOD] Saving %s to cache\n", otp)
}

func NewOTP(iotp IOTP) *OTP {
	return &OTP{
		iOTP: iotp,
	}
}
