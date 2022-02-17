package main

type IMediator interface {
	Notify(sender IComponent, event string)
}
