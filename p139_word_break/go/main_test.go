package p139_word_break

import "testing"

func TestProvided(t *testing.T) {
	if !wordBreak("leetcode", []string{"leet", "code"}) {
		t.Fatal("expected true")
	}

	if !wordBreak("applepenapple", []string{"apple", "pen"}) {
		t.Fatal("expected true")
	}

	if wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}) {
		t.Fatal("expected false")
	}
}

func TestBasic(t *testing.T) {
	if !wordBreak("abcdefg", []string{"a", "ab", "bcd", "c", "defg"}) {
		t.Fatal("expected true")
	}

	if !wordBreak("abcdefg", []string{"abcdefg", "a", "q"}) {
		t.Fatal("expected true")
	}

	if wordBreak("abcdefg", []string{"q"}) {
		t.Fatal("expected false")
	}

	if !wordBreak("abcdefg", []string{"abcde", "b", "efg", "cd", "a"}) {
		t.Fatal("expected true")
	}

	if !wordBreak("abcdefg", []string{"abcde", "b", "efg", "cd", "a", "c", "d", "e", "f"}) {
		t.Fatal("expected true")
	}

	if wordBreak("aaaaaaaaaaaaaaaaaaaac", []string{"a", "aa", "aaa"}) {
		t.Fatal("expected false")
	}
}

func TestExtra(t *testing.T) {
	if wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}) {
		t.Fatal("expected false")
	}
}