package ransom_note

import "testing"

func TestProvided(t *testing.T) {
	if canConstruct("a", "b") != false {
		t.Fatal("expected false")
	}

	if canConstruct("aa", "ab") != false {
		t.Fatal("expected false")
	}

	if canConstruct("aa", "aab") != true {
		t.Fatal("expected true")
	}
}

func TestBasic(t *testing.T) {
	if canConstruct("a", "a") != true {
		t.Fatal("expected true")
	}

	if canConstruct("allaboutthatransomdawgineeditmorethanyou", "allaboutthatransomdawgineeditmorethanyoubruhlikeforrealsies") != true {
		t.Fatal("expected true")
	}

	if canConstruct("allaboutthatransomdawgineeditmorethanyoubruhword", "allaboutthatransomdawgineeditmorethanyoubruhworm") != false {
		t.Fatal("expected false")
	}
}
