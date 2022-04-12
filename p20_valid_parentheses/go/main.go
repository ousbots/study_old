package valid_parentheses

func isValid(s string) bool {
	open := []rune{}
	
	for _, letter := range s {
		switch letter {
			case '(':
				open = append(open, letter)
			case '[':
				open = append(open, letter)
			case '{':
				open = append(open, letter)

			case ')':
				if len(open) == 0 {
					return false
				}
				last := open[len(open)-1]
				open = open[:len(open)-1]
				if last != '(' {
					return false
				}

			case ']':
				if len(open) == 0 {
					return false
				}
				last := open[len(open)-1]
				open = open[:len(open)-1]
				if last != '[' {
					return false
				}

			case '}':
				if len(open) == 0 {
					return false
				}
				last := open[len(open)-1]
				open = open[:len(open)-1]
				if last != '{' {
					return false
				}

			default:
				return false
		}
	}
		
	if len(open) == 0 {
		return true
	}

	return false
}
