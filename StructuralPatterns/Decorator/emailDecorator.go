package main

import "fmt"

type EmailDecorator struct {
	base *BaseNotifierDecorator
}

func (email *EmailDecorator) send(message string) {
	fmt.Println("Sending email notification...")
	email.base.send(message)
}

func NewEmailDecorator(base BaseNotifierDecorator) *EmailDecorator {
	return &EmailDecorator{
		base: &base,
	}
}
