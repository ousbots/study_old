package product_of_array_except_self

import "testing"

func TestProvided(t *testing.T) {
	expected := []int{24, 12, 8, 6}
	if ans := productExceptSelf([]int{1, 2, 3, 4}); !comp(ans, expected) {
		t.Fatalf("%v is not %v", ans, expected)
	}

	expected = []int{0, 0, 9, 0, 0}
	if ans := productExceptSelf([]int{-1, 1, 0, -3, 3}); !comp(ans, expected) {
		t.Fatalf("%v is not %v", ans, expected)
	}
}

func TestBasic(t *testing.T) {
	expected := []int{0, 0, 0, 0, 0, 0, 0}
	if ans := productExceptSelf([]int{1, 0, 2, 0, 3, 0, 4}); !comp(ans, expected) {
		t.Fatalf("%v is not %v", ans, expected)
	}
}

func comp(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
