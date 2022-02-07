package main

type Box struct {
	products []IComponent
}

func (b *Box) Price() float64 {
	packagePrice := 50.0

	for _, p := range b.products {
		packagePrice += p.Price()
	}

	return packagePrice
}

func (b *Box) AddProducts(products []IComponent) {
	b.products = append(b.products, products...)
}

func NewBox() *Box {
	return &Box{}
}
