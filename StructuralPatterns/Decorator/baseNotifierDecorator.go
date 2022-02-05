package main

type BaseNotifierDecorator struct {
	Notifier INotifier
}

func NewBaseNotifierDecorator(notifier INotifier) *BaseNotifierDecorator {
	return &BaseNotifierDecorator{
		Notifier: notifier,
	}
}
