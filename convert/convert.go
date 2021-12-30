package convert

import (
	"strconv"
	"strings"
)

func ConvertToFloat64(weightInput string, heightInput string) (float64, float64) {
	weightInput = strings.Replace(weightInput, "\n", "", -1) // -1 implies, make these changes to all the matching "\n"s.
	heightInput = strings.Replace(heightInput, "\n", "", -1)
	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)
	return weight, height
}
