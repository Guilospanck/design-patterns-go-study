package main

type SMSDecorator struct {
	BaseNotifierDecorator
}

func NewSMSDecorator() *SMSDecorator {
	return &SMSDecorator{}
}
