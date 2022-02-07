package main

type Product struct{}

func (p *Product) Price() float64 {
	return 100
}

func NewProduct() IComponent {
	return &Product{}
}
