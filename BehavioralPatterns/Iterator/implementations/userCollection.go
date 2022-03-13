package implementations

import "base/BehavioralPatterns/Iterator/interfaces"

// Concrete Collection
type UserCollection struct {
	users []*User
}

func (u *UserCollection) CreateIterator() interfaces.IIterator {
	return &UserIterator{
		users: u.users,
	}
}

func NewUserCollection(users []*User) *UserCollection {
	return &UserCollection{
		users: users,
	}
}
