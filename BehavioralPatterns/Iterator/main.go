package main

import "fmt"

func main() {

	// create users
	user1 := &User{
		name: "Guilherme",
		age:  25,
	}
	user2 := &User{
		name: "Larry",
		age:  50,
	}

	// collection of users
	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	// iterator
	iterator := userCollection.createIterator()

	// iterates
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}

}
