package main

func main() {
	message := "[SMS] This is an OTP message"
	randomOTP := "1234556789"

	// sms otp
	smsOTP := NewSMSOTP(message, randomOTP)

	// email otp
	emailOTP := NewEmailOTP()

	// get otp
	otp := NewOTP(smsOTP)

	// generate sms otp using skeleton
	otp.generateAndSendOTP(4)

	// get otp with email
	otp = NewOTP(emailOTP)

	// generate sms otp using skeleton
	otp.generateAndSendOTP(5)

}
