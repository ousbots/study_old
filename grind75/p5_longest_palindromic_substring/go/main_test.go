package p5_longest_palindromic_substring

import "testing"

func TestBasic(t *testing.T) {
	valid := "a"
	if ans := longestPalindrome("a"); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = "aa"
	if ans := longestPalindrome("aab"); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = "aba"
	if ans := longestPalindrome("aaba"); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = "cmooomc"
	if ans := longestPalindrome("abbacmooomcfes"); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func TestExtra(t *testing.T) {
	valid := "ddtattarrattatdd"
	if ans := longestPalindrome("babaddtattarrattatddetartrateedredividerb"); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}
}
