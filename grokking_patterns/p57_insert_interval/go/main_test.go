package insert_interval

import "testing"

type testCase struct {
	interval    [][]int
	newInterval []int
	expected    [][]int
}

func comp(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i][0] != b[i][0] || a[i][1] != b[i][1] {
			return false
		}
	}

	return true
}

func TestBasic(t *testing.T) {
	tests := []testCase{
		{interval: [][]int{{1, 2}}, newInterval: []int{3, 4}, expected: [][]int{{1, 2}, {3, 4}}},
		{interval: [][]int{{1, 2}, {3, 4}}, newInterval: []int{4, 5}, expected: [][]int{{1, 2}, {3, 5}}},
		{interval: [][]int{{1, 2}, {3, 4}}, newInterval: []int{2, 5}, expected: [][]int{{1, 5}}},
		{interval: [][]int{{1, 3}, {6, 9}}, newInterval: []int{2, 5}, expected: [][]int{{1, 5}, {6, 9}}},
		{interval: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, newInterval: []int{4, 8}, expected: [][]int{{1, 2}, {3, 10}, {12, 16}}},
	}

	for _, test := range tests {
		ans := insert(test.interval, test.newInterval)
		if !comp(ans, test.expected) {
			t.Fatalf("case interval=%v newInterval=%v expected %v not %v", test.interval, test.newInterval, test.expected, ans)
		}
	}
}
