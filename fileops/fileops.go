package fileops

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"

	"github.com/danilobml/price-calculator/conversion"
)

func GetFloatsFromFile(filename string) ([]float64, error) {
	values := []float64{}

	file, err := os.Open(filename)
	if err != nil {
		return []float64{}, errors.New("failed to open file " + filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		err = scanner.Err()
		if err != nil {
			return []float64{}, errors.New("failed to scan line in file " + filename)
		}
		floatValue, err := conversion.StringToFloat(line)
		if err != nil {
			return []float64{}, err
		}
		values = append(values, floatValue)
	}
	return values, nil
}

func WriteJSON(data any, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return errors.New("failed to create to json file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to generate json")
	}
	return nil
}
