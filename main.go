package main

import (
	"fmt"
	"reflect"
	"test/calc"
	"test/core"
	"test/core/middlware"
)

func main() {
	input := core.GetInput()
	firstOperand, mathOperator, lastOperand := core.InputExpression(input)
	if reflect.TypeOf(lastOperand).Kind() == reflect.Int {
		output := calc.StringAndNumber(firstOperand, mathOperator, lastOperand)
		result := middlware.CheckLength(output, 40)
		fmt.Printf(result)
	} else {
		output := calc.StringAndString(firstOperand, mathOperator, lastOperand)
		result := middlware.CheckLength(output, 40)
		fmt.Println(result)

	}
}
