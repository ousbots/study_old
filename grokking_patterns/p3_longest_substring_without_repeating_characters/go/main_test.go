package longest_substring_without_repeating_characters

import "testing"

type TestCase struct {
	input  string
	output int
}

func TestProvided(t *testing.T) {
	cases := []TestCase{
		{input: "abcabcbb", output: 3},
		{input: "bbbbb", output: 1},
		{input: "pwwkew", output: 3},
	}

	for _, test := range cases {
		ans := lengthOfLongestSubstring(test.input)
		if ans != test.output {
			t.Fatalf("for %v expected %d not %d", test.input, test.output, ans)
		}
	}
}

func TestBasic(t *testing.T) {
	cases := []TestCase{
		{input: "a", output: 1},
		{input: "aa", output: 1},
		{input: "abab", output: 2},
		{input: "abcabcdabca", output: 4},
	}

	for _, test := range cases {
		ans := lengthOfLongestSubstring(test.input)
		if ans != test.output {
			t.Fatalf("for %v expected %d not %d", test.input, test.output, ans)
		}
	}
}
