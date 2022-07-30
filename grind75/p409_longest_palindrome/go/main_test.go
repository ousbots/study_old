package longest_palindrome

import "testing"

func TestProvided(t *testing.T) {
	if ans := longestPalindrome("abccccdd"); ans != 7 {
		t.Fatalf("%d is not 7", ans)
	}

	if ans := longestPalindrome("a"); ans != 1 {
		t.Fatalf("%d is not 1", ans)
	}

	if ans := longestPalindrome("bb"); ans != 2 {
		t.Fatalf("%d is not 2", ans)
	}
}

func TestBasic(t *testing.T) {
	if ans := longestPalindrome("aA"); ans != 1 {
		t.Fatalf("%d is not 1", ans)
	}

	if ans := longestPalindrome("aAbbcc"); ans != 5 {
		t.Fatalf("%d is not 5", ans)
	}
}
