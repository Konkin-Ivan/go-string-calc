package calc

import (
	"fmt"
	"test/calc/engine"
	"test/calc/middleware"
)

func StringAndNumber(firstOperand string, mathOperator string, lastOperand interface{}) string {

	var operand = middleware.TrimQuotes(firstOperand)
	var n = lastOperand.(int)

	switch mathOperator {
	case "*":
		calc := engine.Concatenator(operand, n)
		result := fmt.Sprintf("\"%v\"", calc)
		return result
	case "/":
		calc := engine.Divider(operand, n)
		result := fmt.Sprintf("\"%v\"", calc)
		return result
	default:
		panic("Unknown mathematical operator SaN")
	}

	return "StringAndNumber"
}
