package middlware

import (
	"strconv"
	"strings"
)

func CheckQuotes(input string) bool {
	return strings.HasPrefix(input, `"`) && strings.HasSuffix(input, `"`)
}

func CheckOperators(input string) bool {
	switch input {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	case "/":
		return true
	default:
		return false
	}
}

func CheckNumbers(input int) bool {
	if input >= 1 && input <= 10 {
		return true
	}
	return false
}

func ThisType(input string) (interface{}, error) {
	validInt := map[string]bool{
		"1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true, "10": true,
	}

	if _, ok := validInt[input]; ok {
		return strconv.Atoi(input)
	} else {
		return input, nil
	}
}

func CheckLength(result string, length int) string {
	if len(result) > length {
		result = result[:length] + "..."
	}
	return result
}
