package core

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter an expression (example: \"Hello\" + \"World!\" or: \"World!!!\" / 3): ")
	input, _ := reader.ReadString('\n')
	return input
}
