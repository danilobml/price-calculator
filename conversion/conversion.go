package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(str string) (float64, error) {
	floatValue, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0, errors.New("failed to convert string to float: " + str)
	}
	return floatValue, nil
}
