package main

type BaseNotifierDecorator struct {
	Notifier INotifier
}

func (b *BaseNotifierDecorator) send(message string) {
	b.Notifier.send(message)
}

func NewBaseNotifierDecorator(notifier INotifier) *BaseNotifierDecorator {
	return &BaseNotifierDecorator{
		Notifier: notifier,
	}
}
