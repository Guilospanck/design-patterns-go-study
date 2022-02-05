package main

import "fmt"

type EmailDecorator struct {
	base INotifier
}

func (email *EmailDecorator) send(message string) {
	fmt.Println("Sending email notification...")
	email.base.send(message)
}

func NewEmailDecorator(base INotifier) *EmailDecorator {
	return &EmailDecorator{
		base: base,
	}
}
