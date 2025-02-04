package main

import (
	"fmt"
)

func main() {
	var revenue float64 = 1000
	var expenses float64 = 100
	var taxRate float64 = 5

	// fmt.Print("Revenue")
	// fmt.Scan(&revenue)
	revenue = getRevenue()
	fmt.Print("Expenses")
	fmt.Scan(&expenses)
	fmt.Print("taxRate")
	fmt.Scan(&taxRate)

	// earningsBeforeTax := revenue - expenses
	// profit := earningsBeforeTax - (earningsBeforeTax * taxRate/100)
	// ratio := earningsBeforeTax / profit

	// fmt.Println(earningsBeforeTax, profit, ratio)
	earningsBeforeTax, profit, ratio := earnings(revenue, expenses, taxRate)
	fmt.Println(earningsBeforeTax, profit, ratio)
}

func getRevenue() (revenue float64) {
	fmt.Print("Revenue")
	fmt.Scan((&revenue))
	return revenue
}

func earnings(revenue float64, expenses float64, taxRate float64)(earningsBeforeTax float64, profit float64, ratio float64) {
	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax - (earningsBeforeTax * taxRate/100)
	ratio = earningsBeforeTax / profit
	fmt.Printf("Earnings: %.2f\nProfit: %0f\nRatio: %.2f\n", earningsBeforeTax, profit, ratio)
	return
}