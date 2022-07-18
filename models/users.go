package models

import (
	"errors"
	"fmt"
)

type UserType int

const (
	Admin UserType = iota
	Normal
	Guest
)

type User struct {
	ID       int
	Name     string
	Type     UserType
	Password string
	Email    string
	Passport string
}

var (
	users         []*User
	loggedInUsers []*User
	nextID        = 1
)

func GetUsers() []*User {
	return users
}

func GetLoggedInUsers() []*User {
	return loggedInUsers
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not have an ID assigned")
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
	return User{}, fmt.Errorf("User id: %d not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, us := range users {
		if us.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User id: %d not found", u.ID)
}

func RemoveUserByID(id int) error {
	for i, us := range users {
		if us.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User id: %d not found", id)
}

func GetUserByDetail(u User)(User,error){
	for _, us := range users {
		if us.Email==u.Email && us.Password==u.Password {			
			return *us, nil
		}
	}
	return User{},fmt.Errorf("Wrong User details")
}