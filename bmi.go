package main

import (
	"github.com/joelpatel/bmi/calculateBMI"
	"github.com/joelpatel/bmi/convert"
	"github.com/joelpatel/bmi/input"
	"github.com/joelpatel/bmi/print"
)

func main() {
	print.Print_header()
	weightInput, heightInput := input.Input()
	weight, height := convert.ConvertToFloat64(weightInput, heightInput)

	bmi := calculateBMI.CalculateBMI(weight, height)
	print.PrintBMI(bmi)
}
