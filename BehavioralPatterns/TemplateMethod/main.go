package main

func main() {
	// sms otp
	message := "[SMS] This is an OTP message"
	randomOTP := "1234556789"
	smsOTP := NewSMSOTP(message, randomOTP)

	// email otp
	emailMessage := "[EMAIL] This is an OTP message"
	emailRandomOTP := "456487945"
	emailOTP := NewEmailOTP(emailRandomOTP, emailMessage)

	// get otp
	otp := NewOTP(smsOTP)

	// generate sms otp using skeleton
	otp.generateAndSendOTP(4)

	// get otp with email
	otp = NewOTP(emailOTP)

	// generate sms otp using skeleton
	otp.generateAndSendOTP(5)

}
