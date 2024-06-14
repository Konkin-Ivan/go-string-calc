package middleware

import "strings"

func TrimQuotes(value string) string {
	return strings.Trim(value, "\"")
}

func RemoveSubstring(input, toRemove string) string {
	return strings.Replace(input, toRemove, "", 1)
}
