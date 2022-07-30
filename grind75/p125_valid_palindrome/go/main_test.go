package valid_palindrome

import "testing"

func TestProvided(t *testing.T) {
	word := "A man, a plan, a canal: Panama"
	if isPalindrome(word) != true {
		t.Fatalf("%s is a palindrome", word)
	}

	word = "race a car"
	if isPalindrome(word) != false {
		t.Fatalf("%s is not a palindrome", word)
	}

	word = " "
	if isPalindrome(word) != true {
		t.Fatalf("%s is a palindrome", word)
	}
}

func TestBasic(t *testing.T) {
	word := "abcdcba"
	if isPalindrome(word) != true {
		t.Fatalf("%s is a palindrome", word)
	}

	word = "abcdefdcba"
	if isPalindrome(word) != false {
		t.Fatalf("%s is not a palindrome", word)
	}
}
