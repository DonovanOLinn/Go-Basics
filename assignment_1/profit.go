package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue")
	fmt.Scan(&revenue)
	fmt.Print("Expenses")
	fmt.Scan(&expenses)
	fmt.Print("taxRate")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax - (earningsBeforeTax * taxRate/100)
	ratio := earningsBeforeTax / profit

	fmt.Println(earningsBeforeTax, profit, ratio)
}