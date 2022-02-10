package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// create users
	user1 := &User{
		Name: "Guilherme",
		Age:  25,
	}
	user2 := &User{
		Name: "Larry",
		Age:  50,
	}

	// collection of users
	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	// iterator
	iterator := userCollection.createIterator()

	userUnmarshalled := &User{}
	// iterates
	for iterator.hasNext() {
		user := iterator.getNext()
		json.Unmarshal([]byte(user), userUnmarshalled)
		fmt.Printf("User is %+v\n", userUnmarshalled)
	}

	/* ============= Friend =================== */
	friend1 := NewFriend("Friend1", "Neighborhood1", "City1", 1)
	friend2 := NewFriend("Friend2", "Neighborhood2", "City2", 2)
	friend3 := NewFriend("Friend3", "Neighborhood3", "City3", 3)

	// friends collection
	friendsCollection := NewFriendsCollection([]*Friend{friend1, friend2, friend3})

	// creates iterator for this collection
	friendsIterator := friendsCollection.createIterator()

	friendUnmarshalled := &Friend{}
	// iterates
	for friendsIterator.hasNext() {
		friend := friendsIterator.getNext()
		err := json.Unmarshal([]byte(friend), friendUnmarshalled)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("Friend: %+v\n", friendUnmarshalled)
	}

}
