package main

import "fmt"

type INotifier interface {
	send(message string)
}

type Notifier struct{}

func (n *Notifier) send(message string) {
	fmt.Printf("Message %s sent!", message)
}

func NewNotifier() *Notifier {
	return &Notifier{}
}
