package majority_element

import "testing"

func TestProvided(t *testing.T) {
	if ans := majorityElement([]int{3, 2, 3}); ans != 3 {
		t.Fatalf("%d is not %d", ans, 3)
	}

	if ans := majorityElement([]int{2, 2, 1, 1, 1, 2, 2}); ans != 2 {
		t.Fatalf("%d is not %d", ans, 2)
	}
}

func TestBasic(t *testing.T) {
	if ans := majorityElement([]int{1}); ans != 1 {
		t.Fatalf("%d is not %d", ans, 1)
	}
}
