package core

import (
	"fmt"
	"strings"
	"test/core/middlware"
)

func InputExpression(input string) (string, string, interface{}) {
	input = strings.TrimSpace(input)
	parts := middlware.Fields(input)
	//fmt.Printf("Parts0: %T Parts1: %T Parts2: %T", parts[0], parts[1], parts[2])
	if len(parts) != 3 {
		panic("Invalid input: the mathematical expression must contain two operands and one operator (+, -, /, *).")
	} else if len(parts[0]) > 10 && len(parts[2]) > 10 {
		panic("Invalid input: The entered value must be no more than 10 characters.")
	}
	intOrString, _ := middlware.ThisType(parts[2])

	isValid := validateInput(parts[0], parts[1], intOrString)
	if isValid == false {
		panic(fmt.Sprintf("Invalid expression value: %s\n", parts))
	}

	return parts[0], parts[1], intOrString
}

func validateInput(firstNumber string, operator string, lastNumber interface{}) bool {

	var firstCheck bool
	var lastCheck bool

	firstCheck = middlware.CheckQuotes(firstNumber)

	operatorsCheck := middlware.CheckOperators(operator)

	switch lastNumber.(type) {
	case string:
		lastCheck = middlware.CheckQuotes(lastNumber.(string))
	case int:
		lastCheck = middlware.CheckNumbers(lastNumber.(int))

	}

	if firstCheck == true && operatorsCheck == true && lastCheck == true {
		return true
	} else {
		return false
	}

}
