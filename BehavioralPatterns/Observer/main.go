package main

func main() {
	// creates data
	newData := Data{
		Name: "Product 1",
		Size: 68,
	}

	// subject (Publisher)
	publisher := NewPublisher(newData)

	// Observers (Subscribers)
	firstSubscriber := NewCustomerSubscriber("guilherme@mail.com")
	secondSubscriber := NewCustomerSubscriber("larry@mail.com")

	// Register subscribers (observers)
	publisher.AddSubscriber(firstSubscriber)
	publisher.AddSubscriber(secondSubscriber)

	// Updating state
	publisher.updateAvailabilityOfData()

}
