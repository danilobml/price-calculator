package main

import (
	"fmt"

	"github.com/danilobml/price-calculator/fileops"
	"github.com/danilobml/price-calculator/price"
)

func main() {
	prices, err := fileops.GetPrices()
	if err != nil {
		fmt.Println("Error reading prices file:", err)
		return
	}
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		newJob := price.NewTaxIncludedPriceJob(taxRate, prices)
		newJob.Process()
		fmt.Println(*newJob)
	}

	
}
