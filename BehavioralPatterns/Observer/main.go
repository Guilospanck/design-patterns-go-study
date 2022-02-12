package main

func main() {
	// Observers (Subscribers)
	firstSubscriber := NewCustomerSubscriber("guilherme@mail.com")
	secondSubscriber := NewCustomerSubscriber("larry@mail.com")

	// creates data
	newData := Data{
		Name: "Product 1",
		Size: 68,
	}

	// subject (Publisher) that is the "father" Of All others publishers
	publisher := NewPublisher(newData.Marshal())

	// Register subscribers (observers)
	publisher.AddSubscriber(firstSubscriber)
	publisher.AddSubscriber(secondSubscriber)

	// Updating state
	publisher.updateAvailabilityOfData()

}
