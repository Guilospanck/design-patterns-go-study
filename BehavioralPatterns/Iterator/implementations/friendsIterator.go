package implementations

import "encoding/json"

type FriendsIterator struct {
	index   int
	Friends []*Friend
}

func (iterator *FriendsIterator) HasNext() bool {
	return iterator.index < len(iterator.Friends)
}

func (iterator *FriendsIterator) GetNext() string {
	if iterator.HasNext() {
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
