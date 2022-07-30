package ransom_note

func canConstructUtf8(ransomNote string, magazine string) bool {
	count := make(map[rune]int)

	for _, elem := range ransomNote {
		count[elem] -= 1
	}

	for _, elem := range magazine {
		count[elem] += 1
	}

	for _, val := range count {
		if val < 0 {
			return false
		}
	}

	return true
}

// Faster, but only works for ascii.
func canConstruct(ransomNote string, magazine string) bool {
	letters := new([26]rune)

	for _, elem := range ransomNote {
		letters[elem-'a'] -= 1
	}

	for _, elem := range magazine {
		letters[elem-'a'] += 1
	}

	for _, count := range letters {
		if count < 0 {
			return false
		}
	}

	return true
}
