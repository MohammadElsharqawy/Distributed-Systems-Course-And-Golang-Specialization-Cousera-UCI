package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID      int
	Name    string // must be Name not name :D, won't work
	Age     int
	Address Address
}

var (
	users  []*User
	nextID = 1
)

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user must not include id or it must be zero \n")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)

	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}
