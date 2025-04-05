package main

import (
	"example.com/price-calc/prices"
	"fmt"
)

func main() {
	var taxRates []float64 = []float64{0,0.07,0.10,0.15}

	// var result map[float64][]float64 = make(map[float64][]float64)

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

}