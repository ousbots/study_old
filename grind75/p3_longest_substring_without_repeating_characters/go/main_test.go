package longest_substring_without_repeating_characters

import "testing"

func TestProvided(t *testing.T) {
	if ans := lengthOfLongestSubstring("abcabcbb"); ans != 3 {
		t.Fatalf("expected 3, not %d", ans)
	}

	if ans := lengthOfLongestSubstring("bbbbb"); ans != 1 {
		t.Fatalf("expected 1, not %d", ans)
	}

	if ans := lengthOfLongestSubstring("pwwkew"); ans != 3 {
		t.Fatalf("expected 3, not %d", ans)
	}
}

func TestBasic(t *testing.T) {
	if ans := lengthOfLongestSubstring("abcdefghijklmnopqrstuvwxyzabcdefg"); ans != 26 {
		t.Fatalf("expected 26, not %d", ans)
	}

	if ans := lengthOfLongestSubstring("abcdefghijklmnopaqrstuavwxyzabcdefg"); ans != 21 {
		t.Fatalf("expected 21, not %d", ans)
	}

	if ans := lengthOfLongestSubstring(" "); ans != 1 {
		t.Fatalf("expected 1, not %d", ans)
	}
}