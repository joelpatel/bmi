package input

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joelpatel/bmi/info"
)

var reader = bufio.NewReader(os.Stdin)

func Input() (string, string) {
	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')
	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')
	return weightInput, heightInput
}
