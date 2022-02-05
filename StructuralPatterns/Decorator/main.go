package main

func main() {
	// gets notifier
	notifier := NewNotifier()

	// basic notifier
	basicNotifier := NewBaseNotifierDecorator(notifier)
	basicNotifier.send("Basic Message")

	// sms
	smsNotifier := NewSMSDecorator(*basicNotifier)
	smsNotifier.send("SMS Message")

	// email
	emailNotifier := NewEmailDecorator(*basicNotifier)
	emailNotifier.send("Email Message")

}
