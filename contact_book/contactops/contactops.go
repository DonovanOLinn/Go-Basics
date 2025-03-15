package contactops

import (
	"fmt"
	"bufio"
	"os"

)

type Contact struct {
	Name string
	Phone string
	Email string
}

func AddContact() Contact{
	name := GetUserInput("What is your name? ")
	phone := GetUserInput("What is your phone number? ")
	email := GetUserInput("What is your email? ")

	newContact := Contact {
		Name: name,
		Phone: phone,
		Email: email,
	}

	return newContact

}

func ViewAllContacts(contacts []Contact) {
	for i:=0 ; i < len(contacts) ; i++ {
		fmt.Printf("This is your contact! Name: %v, Email: %v, Phone: %v\n", contacts[i].Name, contacts[i].Email, contacts[i].Phone)
	}
}

func GetUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)

	input, _ := reader.ReadString('\n')
	return input[:len(input)-2]
}