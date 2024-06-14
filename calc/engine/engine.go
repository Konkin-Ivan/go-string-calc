package engine

import (
	"math"
)

func Concatenator(input string, quantity int) string {
	var parts = input
	var n = quantity
	var result string

	for i := 0; i < n; i++ {
		result += parts
	}

	return result
}

func Divider(input string, quantity int) string {
	if quantity <= 0 {
		return ""
	}

	length := len(input)
	partSize := int(math.Ceil(float64(length) / float64(quantity)))

	if partSize > length {
		return input
	}

	return input[:partSize]
}
