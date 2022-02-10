package main

// Concrete Collection
type UserCollection struct {
	users []*User
}

func (u *UserCollection) createIterator() IIterator {
	return &UserIterator{
		users: u.users,
	}

}
