package main

import "fmt"

type CustomerSubscriber struct {
	id string
}

func (sub *CustomerSubscriber) Update(data string) {
	fmt.Printf("[Subscriber] Sending email to %s %s\n", sub.id, data)
}

func (sub *CustomerSubscriber) GetID() string {
	return sub.id
}

func NewCustomerSubscriber(id string) *CustomerSubscriber {
	return &CustomerSubscriber{
		id: id,
	}
}
