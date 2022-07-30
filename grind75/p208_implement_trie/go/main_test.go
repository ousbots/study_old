package implement_trie

import "testing"

func TestProvided(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")

	if trie.Search("apple") != true {
		t.Fatal("expected true")
	}

	if trie.Search("app") != false {
		t.Fatal("expected false")
	}

	if trie.StartsWith("app") != true {
		t.Fatal("expected true")
	}

	trie.Insert("app")

	if trie.Search("app") != true {
		t.Fatal("expected true")
	}
}

func TestBasic(t *testing.T) {
	trie := Constructor()
	trie.Insert("a")

	if trie.Search("a") != true {
		t.Fatal("expected true")
	}

	if trie.Search("ab") != false {
		t.Fatal("expected false")
	}

	if trie.StartsWith("a") != true {
		t.Fatal("expected true")
	}

	trie.Insert("ab")
	trie.Insert("ac")

	if trie.Search("ab") != true {
		t.Fatal("expected true")
	}

	if trie.Search("ac") != true {
		t.Fatal("expected true")
	}

	if trie.Search("ad") != false {
		t.Fatal("expected false")
	}

	if trie.StartsWith("a") != true {
		t.Fatal("expected true")
	}

	if trie.StartsWith("ab") != true {
		t.Fatal("expected true")
	}

	if trie.StartsWith("ac") != true {
		t.Fatal("expected true")
	}

	if trie.StartsWith("b") != false {
		t.Fatal("expected false")
	}

	if trie.StartsWith("c") != false {
		t.Fatal("expected false")
	}

	trie.Insert("falafel")
	trie.Insert("falafe")

	if trie.Search("falaf") != false {
		t.Fatal("expected false")
	}

	if trie.Search("falafel") != true {
		t.Fatal("expected true")
	}

	if trie.Search("falafe") != true {
		t.Fatal("expected true")
	}

	if trie.StartsWith("falaf") != true {
		t.Fatal("expected true")
	}

	if trie.StartsWith("falafe") != true {
		t.Fatal("expected true")
	}

	if trie.StartsWith("falafel") != true {
		t.Fatal("expected true")
	}
}
