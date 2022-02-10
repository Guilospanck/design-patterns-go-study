package main

import (
	"encoding/json"
)

// Concrete Iterator
type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) hasNext() bool {
	return u.index < len(u.users)
}

func (u *UserIterator) getNext() string {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		userBytes, _ := json.Marshal(user)
		userStr := string(userBytes[:])

		return userStr
	}
	return ""
}
