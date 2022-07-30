package longest_palindrome

func longestPalindrome(s string) int {
	letters := make(map[rune]int)
	count := 0

	for _, char := range s {
		val := letters[char] + 1
		if val == 2 {
			count += 2
			val = 0
		}
		letters[char] = val
	}

	for _, val := range letters {
		if val > 0 {
			count++
			break
		}
	}

	return count
}
