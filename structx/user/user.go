package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	dateOfBirth string
	createdAt   time.Time
}

type Admin struct {
	email    string
	password string
	User
}

// Receiver func, adding a method to a struct :)
func (u User) GetUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.dateOfBirth)
}

// user pointer with receiver func, else a copy is passed to the receiver func and the underlying user will not change
func (u *User) ClearUserDetails() {
	u.firstName = ""
	u.lastName = ""
	u.createdAt = time.Time{}
}

func New(firstName, lastName, dateOfBirth string) (*User, error) {
	if firstName == "" || lastName == "" || dateOfBirth == "" {
		return nil, errors.New("first name, last name and date of birth are required")
	}
	return &User{firstName: firstName, lastName: lastName, dateOfBirth: dateOfBirth, createdAt: time.Now()}, nil
}

func NewAdmin(email, password string) (Admin, error) {
	if email == "" || password == "" {
		return Admin{}, errors.New("email and password are required")
	}
	return Admin{
		email: email, password: password,
		User: User{firstName: "Admin", lastName: "Master", dateOfBirth: "01/01/2010", createdAt: time.Now()},
	}, nil
}
