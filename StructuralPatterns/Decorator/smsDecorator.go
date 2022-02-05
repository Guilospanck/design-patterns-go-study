package main

import "fmt"

type SMSDecorator struct {
	base *BaseNotifierDecorator
}

func (sms *SMSDecorator) send(message string) {
	fmt.Println("Sending SMS notification...")
	sms.base.send(message)
}

func NewSMSDecorator(base BaseNotifierDecorator) *SMSDecorator {
	return &SMSDecorator{
		base: &base,
	}
}
