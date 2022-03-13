package implementations

import (
	"encoding/json"
)

// Concrete Iterator
type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) HasNext() bool {
	return u.index < len(u.users)
}

func (u *UserIterator) GetNext() string {
	if u.HasNext() {
		user := u.users[u.index]
		u.index++
		userBytes, _ := json.Marshal(user)
		userStr := string(userBytes[:])

		return userStr
	}
	return ""
}
