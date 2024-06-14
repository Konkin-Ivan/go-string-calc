package calc

import (
	"fmt"
	"test/calc/middleware"
)

func StringAndString(firstOperand string, mathOperator string, lastOperand interface{}) string {
	var n string
	if str, ok := lastOperand.(string); ok {
		n = middleware.TrimQuotes(str)
	} else {
		panic("lastOperand is not of type string")
	}
	operand := middleware.TrimQuotes(firstOperand)

	switch mathOperator {
	case "+":
		result := fmt.Sprintf("\"%v%v\"", operand, n)
		return result
	case "-":
		calc := middleware.RemoveSubstring(operand, n)
		result := fmt.Sprintf("\"%v\"", calc)
		return result
	default:
		panic("Unknown mathematical operator SaS")
	}
}
