package main

import (
	"fmt"
	"log"
)

func main() {
	// publisher (subject)
	publisher := NewPublisher()

	// observers (subscribers)
	subscriber1 := NewSubscriber()
	subscriber2 := NewSubscriber()
	subscriber3 := NewSubscriber()

	// register subscribers to publisher
	topic1 := "WSL/GUI/10"
	topic2 := "WSL/LARRY/50"

	log.Println("=============== [ADDING SUBSCRIBERS] ================")
	publisher.AddSubscriber(topic1, subscriber1)
	publisher.AddSubscriber(topic2, subscriber2)
	publisher.AddSubscriber(topic1, subscriber3)
	fmt.Println()

	// notify just one subscriber
	log.Println("=============== [NOTIFYING ONE SUBSCRIBER] ================")
	publisher.Notify(topic1, subscriber1)
	fmt.Println()

	// notify all topics
	log.Println("=============== [NOTIFYING ALL TOPICS] ================")
	publisher.NotifyAll()
	fmt.Println()

	// notify just one topic
	log.Printf("=============== [NOTIFYING TOPIC %s] ================\n", topic1)
	publisher.NotifyAllFromTopic(topic1)
	fmt.Println()

	log.Printf("=============== [NOTIFYING TOPIC %s] ================\n", topic2)
	publisher.NotifyAllFromTopic(topic2)
	fmt.Println()

	// changing state
	publisher.UpdateState()
	log.Println("=============== [NOTIFYING ALL TOPICS AFTER CHANGING STATE] ================")
	publisher.NotifyAll()
	fmt.Println()

	// removing a subscriber
	publisher.RemoveSubscriber(topic1, subscriber3)
	log.Println("=============== [NOTIFYING ALL TOPICS AFTER REMOVING SUBSCRIBER] ================")
	publisher.NotifyAll()
	fmt.Println()

	// removing a subscriber from a topic in which he's not in.
	publisher.RemoveSubscriber(topic1, subscriber2)
	log.Println("=============== [NOTIFYING ALL TOPICS AFTER REMOVING SUBSCRIBER NOT IN TOPIC] ================")
	publisher.NotifyAll()
	fmt.Println()

}
