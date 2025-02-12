package user

import (
	"fmt"
	"time"
	"errors"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User User 
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email: email,
		password: password,
		User: User{
				firstName: "ADMIN",
				lastName: "ADMIN",
				birthDate: "---",
				createdAt: time.Now(),
		},
	}
}

func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name, and birth date are required.")
	}

	return &User{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// You can attach methods to structs as seen on the line below
func (u *User) OutputUserDetails() {
	// With the use of the *user, you are pointing at the memory location. BUT the below line still works. Why? 
	// The reason is works is because GO knows what you are trying to do. In reality it would be technically better to do this: 
	// (*u).firstName but no one does because the below line is much MUCH easier to read
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

