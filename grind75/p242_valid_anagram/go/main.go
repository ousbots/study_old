package valid_anagram

func isAnagramSlow(s string, t string) bool {
	sCount := make(map[rune]int32)
	tCount := make(map[rune]int32)

	for _, letter := range s {
		sCount[letter] += 1
	}

	for _, letter := range t {
		tCount[letter] += 1
	}

	if len(sCount) != len(tCount) {
		return false
	}

	for key := range sCount {
		if sCount[key] != tCount[key] {
			return false
		}
	}

	return true
}

func isAnagram(s string, t string) bool {
	count := make([]int32, 26)

	if len(s) != len(t) {
		return false
	}

	for index := range s {
		count[s[index]-'a'] += 1
		count[t[index]-'a'] -= 1
	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}

	return true
}

func isAnagramUnicode(s string, t string) bool {
	count := make(map[rune]int32, 26)

	if len(s) != len(t) {
		return false
	}

	for _, letter := range s {
		count[letter] += 1
	}

	for _, letter := range t {
		count[letter] -= 1
	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}

	return true
}
