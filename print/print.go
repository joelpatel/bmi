package print

import (
	"fmt"

	"github.com/joelpatel/bmi/info"
)

func Print_header() {
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)
}

func PrintBMI(bmi float64) {
	fmt.Printf(info.OutputString, bmi)
}
