package valid_anagram

import "testing"

func TestProvided(t *testing.T) {
	if isAnagram("anagram", "nagaram") != true {
		t.Fatal("anagram and nagaram are anagrams")
	}

	if isAnagram("rat", "car") == true {
		t.Fatal("rat and car are  not anagrams")
	}
}

func TestBasic(t *testing.T) {
	if isAnagram("a", "a") != true {
		t.Fatal("a and a are anagrams")
	}

	if isAnagram("a", "b") == true {
		t.Fatal("a and b are not anagrams")
	}

	if isAnagram("aaaa", "aaab") == true {
		t.Fatal("aaaa and aaab are not anagrams")
	}

	if isAnagram("aaaa", "aaa") == true {
		t.Fatal("a and b are not anagrams")
	}
}
