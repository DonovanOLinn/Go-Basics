package main

import (
	"fmt"

	"example.com/price-calc/cmdmanager"
	// "example.com/price-calc/filemanager"
	"example.com/price-calc/prices"
)



func main() {
	var taxRates []float64 = []float64{0,0.07,0.10,0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}

}