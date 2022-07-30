package valid_palindrome

import "strings"

func isPalindrome(s string) bool {
	check := strings.Split(clean(strings.ToLower(s)), "")

	left := 0
	right := len(check) - 1

	for left <= right {
		if check[left] != check[right] {
			return false
		}

		left++
		right--
	}

	return true

}

// clean removes any non-alphanumeric characters from the given string. The input must be lower case.
func clean(s string) string {
	bytes := []byte(s)
	index := 0

	for _, elem := range bytes {
		if ('a' <= elem && elem <= 'z') || ('0' <= elem && elem <= '9') {
			bytes[index] = elem
			index++
		}
	}

	return string(bytes[:index])
}
