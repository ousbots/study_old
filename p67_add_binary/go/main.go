package add_binary

func addBinary(a, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	sum := []rune{}

	carry := false
	for i, j := len(a)-1, len(b)-1; i >= 0; i, j = i-1, j-1 {
		x := a[i]
		y := '0'
		if j >= 0 {
			y = rune(b[j])
		}

		if x == '1' && y == '1' {
			if carry {
				sum = append([]rune{'1'}, sum...)
			} else {
				sum = append([]rune{'0'}, sum...)
				carry = true
			}
		} else if x == '1' || y == '1' {
			if carry {
				sum = append([]rune{'0'}, sum...)
			} else {
				sum = append([]rune{'1'}, sum...)
			}
		} else {
			if carry {
				sum = append([]rune{'1'}, sum...)
				carry = false
			} else {
				sum = append([]rune{'0'}, sum...)
			}
		}
	}

	if carry {
		sum = append([]rune{'1'}, sum...)
	}

	return string(sum)
}
