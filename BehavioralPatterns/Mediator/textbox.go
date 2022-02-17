package main

type Textbox struct {
	Component IComponent
}

func NewTextbox(component IComponent) *Textbox {
	return &Textbox{
		Component: component,
	}
}
