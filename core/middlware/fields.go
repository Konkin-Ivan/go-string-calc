package middlware

func Fields(input string) []string {
	var result []string
	var inQuotes bool
	var start int

	for i, r := range input {
		switch r {
		case '"':
			inQuotes = !inQuotes
		case ' ':
			if !inQuotes {
				if start < i {
					result = append(result, input[start:i])
				}
				start = i + 1
			}
		}
	}

	if start < len(input) {
		result = append(result, input[start:])
	}

	return result
}
