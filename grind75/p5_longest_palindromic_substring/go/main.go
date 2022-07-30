package p5_longest_palindromic_substring

// Naive implementation, terrible performance, worse than O(n^2).
func longestPalindromeSlow(s string) string {
	if isPalindrome(s) {
		return s
	}

	left := longestPalindromeSlow(s[:len(s)-1])
	right := longestPalindromeSlow(s[1:])

	if len(left) > len(right) {
		return left
	} else {
		return right
	}
}

func isPalindrome(s string) bool {
	low := 0
	high := len(s) - 1

	for low < high {
		if s[low] != s[high] {
			return false
		}

		low++
		high--
	}

	return true
}

func longestPalindrome(s string) string {
	start := 0
	end := 0

	for i := range s {
		first := find(s, i, i)
		second := find(s, i, i+1)

		longest := first
		if longest < second {
			longest = second
		}

		if longest > end-start {
			start = i - (longest / 2)
			end = i + (longest / 2) + 1
			if longest%2 == 0 {
				start++
			}
		}
	}

	return s[start:end]
}

func find(s string, start, end int) int {
	length := 0

	for start >= 0 && end < len(s) && s[start] == s[end] {
		length = end - start + 1
		start--
		end++
	}

	return length
}
