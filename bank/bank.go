package main

import "fmt"
import "os"

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() {
	os.ReadFile(accountBalanceFile)
}

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")
	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		
		var choice int
		fmt.Print(("Your Choice: "))
		fmt.Scan(&choice)

		switch choice {
		case 1: 
			fmt.Println("Your balance is", accountBalance)
		case 2: 
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3: 
			fmt.Print("Your withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be grater than 0")
				continue
			}
			
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. Can't draw funds more than account balance")
			}
			accountBalance -= withdrawAmount
			writeBalanceToFile((accountBalance))
			fmt.Println("Balance updated! New amount:", accountBalance)
		default: 
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our Bank")
			return
			// break //If in a switch statement, it will break the switch statement, not the loop
		}

			
	}

}