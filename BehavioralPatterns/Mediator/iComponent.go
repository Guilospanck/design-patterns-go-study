package main

type IComponent interface {
	Click()
	KeyPress()
}

type Component struct {
	Mediator IMediator
}

func (component *Component) Click() {
	component.Mediator.Notify(NewComponent(), "click")
}

func (component *Component) KeyPress() {
	component.Mediator.Notify(NewComponent(), "keypress")
}

func NewComponent() *Component {
	mediator := NewAuthenticationMediator()
	component := &Component{}
	component.Mediator = mediator

	return component
}
