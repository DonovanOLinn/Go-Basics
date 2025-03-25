package main

import (
	"fmt"
)

func main() {
	// var variable_name [SIZE] variable_type
	
	// var balance [10] float32
	
	// You can initialize a GO array by using a single statement as follows: 
	
	// var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	// You can also omit the size of an array by writing: 

	balance := [] float64 {1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)
	
	// You can also then assign a value to a single element of the array like this: 
	
	balance[3] = 50.0
	fmt.Println(balance)

}