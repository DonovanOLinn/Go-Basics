package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 6.5


func main() {
	
	var investmentAmount float64
	var expectedReturnRate = 5.5
	var years float64 = 10

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	// fmt.Print("Expected Return Rate: ")
	outputText("Expected return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Years Invested: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)
	// fmt.Println("Future Value:", futureValue)
	fmt.Printf("Future Value: %.2f\nFuture Value (adjusted for Inflation): %.2f", futureValue, futureRealValue)
	fmt.Println("Future Value (adjusted for Inflation): ", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	// fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// rfv := fv / math.Pow((1+inflationRate/100), years)
	// IF YOU DECLARE YOUR RETURN VALUES IN THE RETURN DECLARATION ABOVE, YOU DON'T NEED TO DEFINE THEM IN CODE
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow((1+inflationRate/100), years)
	// return fv, rfv 
	// Since we know what we are returning above, we don't need to explicitly state that we are returning them below. Probably should, but don't have to.
	return
}