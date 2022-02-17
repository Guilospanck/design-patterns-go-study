package main

type Checkbox struct {
	Component IComponent
}

func (checkbox *Checkbox) Check() {
	component, ok := checkbox.Component.(*Component)
	if ok {
		component.Mediator.Notify(NewCheckbox(checkbox.Component), "check")
	}
}

func NewCheckbox(component IComponent) *Checkbox {
	return &Checkbox{
		Component: component,
	}
}
