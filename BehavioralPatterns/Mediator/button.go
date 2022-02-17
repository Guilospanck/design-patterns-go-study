package main

type Button struct {
	Component IComponent
}

func NewButton(component IComponent) *Button {
	return &Button{
		Component: component,
	}
}
