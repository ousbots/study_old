package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	count := make(map[byte]int)

	back, max := 0, 0
	for front := range s {
		count[s[front]]++

		if count[s[front]] > 1 {
			for ; back < front; back++ {
				count[s[back]]--

				if count[s[back]] == 0 {
					delete(count, s[back])
				}

				if s[back] == s[front] {
					back++
					break
				}
			}
		} else {
			if front-back >= max {
				max = front - back + 1
			}
		}
	}

	return max
}
