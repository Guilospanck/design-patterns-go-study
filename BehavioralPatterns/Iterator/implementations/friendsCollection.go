package implementations

import "base/BehavioralPatterns/Iterator/interfaces"

type Friend struct {
	Name, Neighborhood, City string
	Age                      int
}

func NewFriend(name, neighborhood, city string, age int) *Friend {
	return &Friend{
		Name:         name,
		Neighborhood: neighborhood,
		City:         city,
		Age:          age,
	}
}

type FriendsCollection struct {
	Friends []*Friend
}

func (collection *FriendsCollection) CreateIterator() interfaces.IIterator {
	return NewFriendsIterator(collection.Friends)
}

func (collection *FriendsCollection) AddFriend(friend Friend) {
	collection.Friends = append(collection.Friends, &friend)
}

func NewFriendsCollection(friends []*Friend) *FriendsCollection {
	return &FriendsCollection{
		Friends: friends,
	}
}
