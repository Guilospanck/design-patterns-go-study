package main

type Container struct {
	Elements []IComponent
}

func (c *Container) GetContainerPrice() float64 {
	price := 0.0
	for _, elem := range c.Elements {
		price += elem.Price()
	}

	return price
}

func NewContainer(elements []IComponent) *Container {
	return &Container{
		Elements: elements,
	}
}
