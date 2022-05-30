package p17_letter_combinations_of_a_phone_number

func letterCombinations(digits string) []string {
	buttons := map[byte][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	return build(digits, buttons, []rune{}, 0)
}

func build(digits string, buttons map[byte][]rune, set []rune, index int) []string {
	if index >= len(digits) {
		if len(set) == 0 {
			return []string{}
		}
		return []string{string(set)}
	}

	letters := buttons[digits[index]]

	output := []string{}
	for _, char := range letters {
		tmp := make([]rune, len(set))
		copy(tmp, set)
		tmp = append(tmp, char)

		output = append(output, build(digits, buttons, tmp, index+1)...)
	}

	return output
}

func letters(num rune) []rune {
	switch num {
	case '2':
		return []rune{'a', 'b', 'c'}

	case '3':
		return []rune{'d', 'e', 'f'}

	case '4':
		return []rune{'g', 'h', 'i'}

	case '5':
		return []rune{'j', 'k', 'l'}

	case '6':
		return []rune{'m', 'n', 'o'}

	case '7':
		return []rune{'p', 'q', 'r', 's'}

	case '8':
		return []rune{'t', 'u', 'v'}

	case '9':
		return []rune{'w', 'x', 'y', 'z'}

	default:
		return []rune{}
	}
}
