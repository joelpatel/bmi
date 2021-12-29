package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joelpatel/bmi/info"
)

func main() {
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)
	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')
	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	weightInput = strings.Replace(weightInput, "\n", "", -1) // -1 implies, make these changes to all the matching "\n"s.
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	bmi := weight / (height * height)

	fmt.Printf(info.OutputString, bmi)

}
