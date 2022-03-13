package implementations

// Concrete Collection
type UserCollection struct {
	users []*User
}

func (u *UserCollection) CreateIterator() *UserIterator {
	return &UserIterator{
		users: u.users,
	}
}

func NewUserCollection(users []*User) *UserCollection {
	return &UserCollection{
		users: users,
	}
}
