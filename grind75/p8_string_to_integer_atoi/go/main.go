package p8_string_to_integer_atoi

import "math"

func myAtoi(s string) int {
	value := int64(0)

	index := 0
	for i := range s {
		if s[i] != ' ' {
			break
		}
		index++
	}

	factor := int64(1)
	if index < len(s) {
		if s[index] == '-' {
			factor = -1
			index++
		} else if s[index] == '+' {
			index++
		}
	}

loop:
	for _, letter := range s[index:] {
		if value >= math.MaxInt32 {
			break
		}

		switch letter {
		case '0':
			value = (value * 10) + 0
		case '1':
			value = (value * 10) + 1
		case '2':
			value = (value * 10) + 2
		case '3':
			value = (value * 10) + 3
		case '4':
			value = (value * 10) + 4
		case '5':
			value = (value * 10) + 5
		case '6':
			value = (value * 10) + 6
		case '7':
			value = (value * 10) + 7
		case '8':
			value = (value * 10) + 8
		case '9':
			value = (value * 10) + 9
		default:
			break loop
		}
	}

	value = value * factor

	if value > math.MaxInt32 {
		value = math.MaxInt32
	} else if value < math.MinInt32 {
		value = math.MinInt32
	}

	return int(value)
}
