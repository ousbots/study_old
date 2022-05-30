package p11_container_with_most_water

import "testing"

func TestBasic(t *testing.T) {
	valid := 1
	if ans := maxArea([]int{1, 2}); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = 2
	if ans := maxArea([]int{1, 2, 3}); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = 9
	if ans := maxArea([]int{3, 1, 2, 3}); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = 9
	if ans := maxArea([]int{1, 1, 9, 9, 1, 1, 1, 1}); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = 12
	if ans := maxArea([]int{1, 1, 1, 1, 1, 8, 1, 1, 1, 1, 1, 1, 1}); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}
}