package main

import (
	"fmt"
)

func main() {
	age := 32 //regular variable

	// Will make into type: *int which specifies it is a pointer to an integer
	var agePointer *int // Typically you wouldn't use pointer in value name

	agePointer = &age

	// fmt.Println("Age:", age)
	fmt.Println("Age:", agePointer) //This will print off the pointer location
	fmt.Println("Age:", *agePointer) //Using the * will get the value associated with that pointer
	// For a pointer the "null value" is nil
	
	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}