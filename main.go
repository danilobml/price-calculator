package main

import (
	"fmt"

	"github.com/danilobml/price-calculator/fileops"
	"github.com/danilobml/price-calculator/price"
)

const pricesFilename string = "prices.txt"
const taxRatesFilename string = "rates.txt"

const jobsSavingFile string = "jobs.json"

func main() {
	prices, err := fileops.GetFloatsFromFile(pricesFilename)
	if err != nil {
		fmt.Println("Error reading prices file:", err)
		return
	}
	taxRates, err := fileops.GetFloatsFromFile(taxRatesFilename)
	if err != nil {
		fmt.Println("Error reading prices file:", err)
		return
	}

	jobs := price.GenerateJobs(taxRates, prices)

	fileops.WriteJSON(jobs, jobsSavingFile)
}
