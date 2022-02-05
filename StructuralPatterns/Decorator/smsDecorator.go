package main

import "fmt"

type SMSDecorator struct {
	base INotifier
}

func (sms *SMSDecorator) send(message string) {
	fmt.Println("Sending SMS notification...")
	sms.base.send(message)
}

func NewSMSDecorator(base INotifier) *SMSDecorator {
	return &SMSDecorator{
		base: base,
	}
}
