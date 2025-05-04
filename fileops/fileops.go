package fileops

import (
	"bufio"
	"strconv"
	"os"
)

const filename string = "prices.txt"

func GetPrices() ([]float64, error) {
	prices := []float64{}

	file, err := os.Open(filename)
	if err != nil {
		return []float64{}, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return []float64{}, err
		}
		prices = append(prices, num)
	}
	return prices, nil
}
