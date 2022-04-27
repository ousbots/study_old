package evaluate_reverse_polish_notation

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+":
			val := stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, val)

		case "-":
			val := stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, val)

		case "*":
			val := stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, val)

		case "/":
			val := stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, val)

		default:
			val, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			stack = append(stack, val)
		}
	}

	return stack[0]
}
