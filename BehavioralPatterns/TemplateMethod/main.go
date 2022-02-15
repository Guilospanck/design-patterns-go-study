package main

func main() {
	// sms otp
	smsOTP := NewSMSOTP()

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
