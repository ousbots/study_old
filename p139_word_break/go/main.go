package p139_word_break

import "fmt"

func wordBreakWrong(s string, wordDict []string) bool {
	words := make(map[string]bool)
	max := 0
	for _, word := range wordDict {
		if len(word) > max {
			max = len(word)
		}
		words[word] = true
	}

	return search(s, words, make(map[int]bool), 0, max)
}

func search(s string, words map[string]bool, found map[int]bool, index int, longest int) bool {
	fmt.Printf("checking %s\n", s)
	if len(s) == 0 {
		return true
	}

	if _, ok := found[index]; ok {
		return found[index]
	}

	max := longest
	if max > len(s) {
		max = len(s)
	}

	for i := 1; i <= max; i++ {
		if _, ok := words[s[:i]]; ok {
			if search(s[i:], words, found, index+1, longest) {
				found[index+1] = true
				return true
			} else {
				found[index+1] = false
			}
		}
	}

	return false
}

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		words[word] = true
	}

	found := make(map[int]bool, len(s)+1)
	found[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if words[s[j:i]] && found[j] {
				found[i] = true
				break
			}
		}
	}

	return found[len(s)]
}
