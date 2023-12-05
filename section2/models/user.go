package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID      int
	Name    string
	Age     int
	Address Address
}

var (
	users  []*User
	nextID = 1
)

func AddUser(user User) (User, error) {

	if user.ID != 0 {
		return User{}, errors.New("new user must not include id or it must be zero \n")
	}
	user.ID = nextID
	nextID++
	users = append(users, &user)

	return user, nil

}

func GetUserById(id int) (User, error) {

	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with id %d not found", id)

}
