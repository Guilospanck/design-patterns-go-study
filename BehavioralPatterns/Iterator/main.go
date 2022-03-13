package main

import (
	"base/BehavioralPatterns/Iterator/implementations"
	"encoding/json"
	"fmt"
)

func main() {
	// create users
	user1 := &implementations.User{
		Name: "Guilherme",
		Age:  25,
	}
	user2 := &implementations.User{
		Name: "Larry",
		Age:  50,
	}

	// collection of users
	userCollection := implementations.NewUserCollection([]*implementations.User{user1, user2})

	// iterator
	iterator := userCollection.CreateIterator()

	userUnmarshalled := &implementations.User{}

	// iterates
	for iterator.HasNext() {
		user := iterator.GetNext()
		json.Unmarshal([]byte(user), userUnmarshalled)
		fmt.Printf("User is %+v\n", userUnmarshalled)
	}

	/* ============= Friend =================== */
	friend1 := implementations.NewFriend("Friend1", "Neighborhood1", "City1", 1)
	friend2 := implementations.NewFriend("Friend2", "Neighborhood2", "City2", 2)
	friend3 := implementations.NewFriend("Friend3", "Neighborhood3", "City3", 3)

	// friends collection
	friendsCollection := implementations.NewFriendsCollection([]*implementations.Friend{friend1, friend2, friend3})

	// creates iterator for this collection
	friendsIterator := friendsCollection.CreateIterator()

	friendUnmarshalled := &implementations.Friend{}
	// iterates
	for friendsIterator.HasNext() {
		friend := friendsIterator.GetNext()
		err := json.Unmarshal([]byte(friend), friendUnmarshalled)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("Friend: %+v\n", friendUnmarshalled)
	}

}
