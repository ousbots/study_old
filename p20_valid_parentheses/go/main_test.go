package valid_parentheses

import "testing"

func TestProvided(t *testing.T) {
	if !isValid("()") {
		t.Fatalf("should be true")
	}
	
	if !isValid("()[]{}") {
		t.Fatalf("should be true")
	}
	
	if isValid("(]") {
		t.Fatalf("should be false")
	}
}

func TestBasic(t *testing.T) {
	if isValid("(([{{}})])") {
		t.Fatalf("out of order should be false")
	}
	
	if isValid("(") {
		t.Fatalf("unclosed should be false")
	}
	if isValid("[") {
		t.Fatalf("unclosed should be false")
	}
	if isValid("{") {
		t.Fatalf("unclosed should be false")
	}

	if isValid(")") {
		t.Fatalf("unopened should be false")
	}
	if isValid("]") {
		t.Fatalf("unopend should be false")
	}
	if isValid("}") {
		t.Fatalf("unopened should be false")
	}
	
	if !isValid("([(([{([()]{}{{([])[]}})()}[{}]()]))])") {
		t.Fatalf("should be true")
	}
}