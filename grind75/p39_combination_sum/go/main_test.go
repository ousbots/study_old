package combination_sum

import (
	"sort"
	"testing"
)

func TestProvided(t *testing.T) {
	valid := [][]int{{2, 2, 3}, {7}}
	if ans := combinationSum([]int{2, 3, 6, 7}, 7); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}
	if ans := combinationSum([]int{2, 3, 5}, 8); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func TestBasic(t *testing.T) {
	valid := [][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 2}, {1, 1, 1, 1, 1, 1, 3}, {1, 1, 1, 1, 1, 2, 2}, {1, 1, 1, 1, 2, 3}, {1, 1, 1, 1, 5}, {1, 1, 1, 2, 2, 2}, {1, 1, 1, 3, 3}, {1, 1, 1, 6}, {1, 1, 2, 2, 3}, {1, 1, 2, 5}, {1, 1, 7}, {1, 2, 2, 2, 2}, {1, 2, 3, 3}, {1, 2, 6}, {1, 3, 5}, {2, 2, 2, 3}, {2, 2, 5}, {2, 7}, {3, 3, 3}, {3, 6}}
	if ans := combinationSum([]int{2, 7, 6, 3, 5, 1}, 9); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func TestComp(t *testing.T) {
	if comp([][]int{{1}}, [][]int{{1}}) != true {
		t.Fatal("[[1]] and [[1]] are the same")
	}

	if comp([][]int{{1}, {3, 2, 1}, {5, 4}}, [][]int{{1}, {4, 5}, {1, 2, 3}}) != true {
		t.Fatal("[[1] [3 2 1] [5 4]] and [[1] [4 5] [1 2 3]] are the same")
	}

	if comp([][]int{{1}, {1, 9, 1}, {1, 1, 3}}, [][]int{{1}, {9, 1, 1}, {1, 1, 1}}) != false {
		t.Fatal("[[1] [1 9 1] [1 1 3]] and [[1] [9 1 1] [1 1 1]] are not the same")
	}
}

func comp(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for _, x := range a {
		sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
	}

	for _, x := range b {
		sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
	}

	for _, check := range a {
		found := false

		index := 0
	loop:
		for _, tmp := range b {
			if len(check) != len(tmp) {
				b[index] = tmp
				index++
				continue
			}

			for i := range check {
				if check[i] != tmp[i] {
					b[index] = tmp
					index++
					continue loop
				}
			}

			found = true
			break
		}

		if !found {
			return false
		}
	}

	return true
}
