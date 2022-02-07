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

	// Put some products into boxes
	productsToBox1 := []IComponent{
		product1, product2, product3,
	}

	productsToBox2 := []IComponent{
		product4, product5, product6,
	}

	productsToBox3 := []IComponent{
		product7, product8, product9,
	}

	// Create some boxes with only products in them
	box1 := NewBox(productsToBox1) // 350
	box2 := NewBox(productsToBox2) // 350
	box3 := NewBox(productsToBox3) // 350

	// create boxes with boxes (which have products)
	box4 := NewBox([]IComponent{box1}) // 400
	box5 := NewBox([]IComponent{box2}) // 400
	box6 := NewBox([]IComponent{box3}) // 400

	// create boxes with more than one box
	box7 := NewBox([]IComponent{box4, box5}) // 850
	box8 := NewBox([]IComponent{box6, box1}) // 800
	box9 := NewBox([]IComponent{box1, box8}) // 1200

	// get Container
	container := NewContainer([]IComponent{
		box1, box2, box3, box4, box5, box6, box7, box8, box9, // 5100.00
		product1, product2, product3, product4, product5, product6, product7, product8, product9, // 900.00
	})

	// Finally gets the price of the bundle
	// Products cost 100, boxes costs 50
	fmt.Println(container.GetContainerPrice()) // 6000.00

}
