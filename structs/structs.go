package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName: userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	// outputUserDetails(&appUser)
	appUser.outputUserDetails()
}

// You can attach methods to structs as seen on the line below
func (u user) outputUserDetails() {
	// With the use of the *user, you are pointing at the memory location. BUT the below line still works. Why? 
	// The reason is works is because GO knows what you are trying to do. In reality it would be technically better to do this: 
	// (*u).firstName but no one does because the below line is much MUCH easier to read
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}