package main

import "encoding/json"

type FriendsIterator struct {
	index   int
	Friends []*Friend
}

func (iterator *FriendsIterator) hasNext() bool {
	return iterator.index < len(iterator.Friends)
}

func (iterator *FriendsIterator) getNext() string {
	if iterator.hasNext() {
		friend := iterator.Friends[iterator.index]
		iterator.index++

		friendByte, _ := json.Marshal(friend)
		friendStr := string(friendByte[:])

		return friendStr
	}

	return ""

}

func NewFriendsIterator(friends []*Friend) *FriendsIterator {
	return &FriendsIterator{
		Friends: friends,
	}
}
