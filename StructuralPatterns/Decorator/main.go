package main

func main() {
	// gets notifier
	notifier := NewNotifier()

	// basic notifier
	basicNotifier := NewBaseNotifierDecorator(notifier)
	basicNotifier.send("Basic Message")

	// sms
	smsNotifier := NewSMSDecorator(basicNotifier)
	smsNotifier.send("SMS Message")

	// email
	emailNotifier := NewEmailDecorator(basicNotifier)
	emailNotifier.send("Email Message")

	// email and sms
	emailWithSMSNotifier := NewEmailDecorator(smsNotifier)
	emailWithSMSNotifier.send("Email And SMS Message")

}
