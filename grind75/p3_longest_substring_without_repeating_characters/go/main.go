package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	count := make(map[rune]int)

	back := 0
	max := 0

	for index, elem := range runes {
		count[elem] += 1

		if count[elem] > 1 {
			for ; back < index; back++ {
				count[runes[back]] -= 1

				if runes[back] == elem {
					break
				}
			}
			back += 1
		} else {
			length := index - back + 1
			if length > max {
				max = length
			}
		}
	}

	return max
}
