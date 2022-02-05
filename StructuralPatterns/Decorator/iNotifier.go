package main

import "fmt"

type INotifier interface {
	send(message string)
}

type Notifier struct{}

func (n *Notifier) send(message string) {
	fmt.Printf("Message '%s' sent!\n===============\n", message)
}

func NewNotifier() INotifier {
	return &Notifier{}
}
