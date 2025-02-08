package main

import "fmt"

func add(number1 int, number2 int) (int) {
	value := number1 + number2
	return value
} 
func subtract(number1 int, number2 int) (int) {
	value := number1 - number2
	return value
} 
func multiply(number1 int, number2 int) (int) {
	value := number1 * number2
	return value
} 
func divide(number1 int, number2 int) (float64) {
	value := float64(number1) / float64(number2)
	return value
} 


func main() {
	var response int
	var input1 int
	var input2 int

	
	for {
		fmt.Println(`Welcome to your Calculator! What would you like to do? 
1. Add
2. Subtract
3. Multiple
4. Divide
5. Exit
		`)
		fmt.Print("Choice: ")
		fmt.Scanln(&response) 

		if response > 4 || response < 1{
			fmt.Print("Thank you for using the calculator today!")
			return
		}
	
		fmt.Print("First number: ")
		fmt.Scanln(&input1)
	
		fmt.Print("Second number: ")
		fmt.Scanln(&input2)
		if response == 1 {
			response := add(input1, input2)
			fmt.Println("Your result: ", response)
		} else if response == 2 {
			response := subtract(input1, input2)
			fmt.Println("Your result: ", response)
		} else if response == 3 {
			response := multiply(input1,input2)
			fmt.Println("Your result: ", response)
		} else if response == 4 {
			response := divide(input1, input2)
			// Should probably do a int => float conversion
			fmt.Printf("Your result: %.02f\n", response)
		} else {
			break
		}
	}

}