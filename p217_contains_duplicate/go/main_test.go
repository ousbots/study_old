package p217_contains_duplicate

import "testing"

func TestProvided(t *testing.T) {
	input := []int{1, 2, 3, 1}
	if !containsDuplicate(input) {
		t.Fatalf("%v expected true", input)
	}

	input = []int{1, 2, 3, 4}
	if containsDuplicate(input) {
		t.Fatalf("%v expected false", input)
	}

	input = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	if !containsDuplicate(input) {
		t.Fatalf("%v expected true", input)
	}
}
