package main

import "fmt"

type AuthenticationMediator struct{}

func (mediator *AuthenticationMediator) Notify(sender IComponent, event string) {
	fmt.Printf("Type of sender is %T and event is %s\n", sender, event)
}

func NewAuthenticationMediator() *AuthenticationMediator {
	return &AuthenticationMediator{}
}
