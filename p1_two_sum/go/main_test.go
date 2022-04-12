package two_sum

import (
	"sort"
	"testing"
)

func TestProvided(t *testing.T) {
	ans := twoSum([]int{2, 7, 11, 15}, 9)
	if !sliceEqual(ans, []int{0, 1}) {
		t.Fatalf("expected [0 1] instead of %v", ans)
	}

	ans = twoSum([]int{3, 2, 4}, 6)
	if !sliceEqual(ans, []int{1, 2}) {
		t.Fatalf("expected [1 2] instead of %v", ans)
	}

	ans = twoSum([]int{3, 3}, 6)
	if !sliceEqual(ans, []int{0, 1}) {
		t.Fatalf("expected [0 1] instead of %v", ans)
	}
}

func TestBasic(t *testing.T) {
	ans := twoSum([]int{0, 0}, 0)
	if !sliceEqual(ans, []int{0, 1}) {
		t.Fatalf("expected [0 1] instead of %v", ans)
	}

	ans = twoSum([]int{1, 5, 1, 5}, 10)
	if !sliceEqual(ans, []int{1, 3}) {
		t.Fatalf("expected [1 3] instead of %v", ans)
	}

	ans = twoSum([]int{0, 1, 16, 5, 30}, 6)
	if !sliceEqual(ans, []int{1, 3}) {
		t.Fatalf("expected [1 3] instead of %v", ans)
	}

	ans = twoSum([]int{99, 11, 33, -1}, 10)
	if !sliceEqual(ans, []int{1, 3}) {
		t.Fatalf("expected [1 3] instead of %v", ans)
	}
}

func sliceEqual(left, right []int) bool {
	if len(left) != len(right) {
		return false
	}

	sort.Slice(left, func(i, j int) bool { return left[i] < left[j] })
	sort.Slice(right, func(i, j int) bool { return right[i] < right[j] })

	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}

	return true
}
