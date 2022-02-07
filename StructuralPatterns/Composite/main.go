package main

import "fmt"

func main() {
	// Create some products
	product1 := NewProduct()
	product2 := NewProduct()
	product3 := NewProduct()
	product4 := NewProduct()
	product5 := NewProduct()
	product6 := NewProduct()
	product7 := NewProduct()
	product8 := NewProduct()
	product9 := NewProduct()

	// Create some boxes with only products in them
	box1 := NewBox()
	box2 := NewBox()
	box3 := NewBox()
	box1.AddProducts([]IComponent{product1, product2, product3}) // 350
	box2.AddProducts([]IComponent{product4, product5, product6}) // 350
	box3.AddProducts([]IComponent{product7, product8, product9}) // 350

	// create boxes with boxes (which have products)
	box4 := NewBox()
	box4.AddProducts([]IComponent{box1}) // 400
	box5 := NewBox()
	box5.AddProducts([]IComponent{box2}) // 400
	box6 := NewBox()
	box6.AddProducts([]IComponent{box3}) // 400

	// create boxes with more than one box
	box7 := NewBox()
	box7.AddProducts([]IComponent{box4, box5}) // 850
	box8 := NewBox()
	box8.AddProducts([]IComponent{box6, box1}) // 800
	box9 := NewBox()
	box9.AddProducts([]IComponent{box1, box8}) // 1200

	// get Container
	container := NewContainer([]IComponent{
		box1, box2, box3, box4, box5, box6, box7, box8, box9, // 5100.00
		product1, product2, product3, product4, product5, product6, product7, product8, product9, // 900.00
	})

	// Finally gets the price of the bundle
	// Products cost 100, boxes costs 50
	fmt.Println(container.GetContainerPrice()) // 6000.00

}
