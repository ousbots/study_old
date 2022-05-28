package p75_sort_colors

import "testing"

func TestProvided(t *testing.T) {
	input := []int{2, 0, 2, 1, 1, 0}
	valid := []int{0, 0, 1, 1, 2, 2}
	sortColors(input)
	if len(input) != len(valid) {
		t.Fatalf("%v is not %v", input, valid)
	}
	for i := range valid {
		if input[i] != valid[i] {
			t.Fatalf("%v is not %v", input, valid)
		}
	}

	input = []int{2, 0, 1}
	valid = []int{0, 1, 2}
	sortColors(input)
	if len(input) != len(valid) {
		t.Fatalf("%v is not %v", input, valid)
	}
	for i := range valid {
		if input[i] != valid[i] {
			t.Fatalf("%v is not %v", input, valid)
		}
	}
}
