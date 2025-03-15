package main

import (
	"example.com/contact-book/contactops"
	"fmt"
)





func main() {
	var contacts []contactops.Contact 
	fmt.Print("Welcome to your Contact list App!")

	for {

		response := contactops.GetUserInput("Would you like to:\n1. Add a contact?\n2. View all contacts?\n3. Quit?\n")
		switch response {
			case "1": 
			contacts = append(contacts, contactops.AddContact())
		case "2":
			contactops.ViewAllContacts(contacts) 
		case "3":
			return
		}
	}

}

