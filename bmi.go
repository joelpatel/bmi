package main

import (
	"github.com/joelpatel/bmi/packages/calculateBMI"
	"github.com/joelpatel/bmi/packages/convert"
	"github.com/joelpatel/bmi/packages/input"
	"github.com/joelpatel/bmi/packages/print"
)

func main() {
	print.Print_header()
	weightInput, heightInput := input.Input()
	weight, height := convert.ConvertToFloat64(weightInput, heightInput)

	bmi := calculateBMI.CalculateBMI(weight, height)
	print.PrintBMI(bmi)
}
