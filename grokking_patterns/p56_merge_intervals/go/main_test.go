package merge_interval

import "testing"

type testCase struct {
	intervals [][]int
	expected  [][]int
}

func check(a, b [][]int) bool {
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
		{intervals: [][]int{{1, 2}}, expected: [][]int{{1, 2}}},
		{intervals: [][]int{{1, 2}, {0, 0}}, expected: [][]int{{0, 0}, {1, 2}}},
		{intervals: [][]int{{1, 2}, {3, 4}}, expected: [][]int{{1, 2}, {3, 4}}},
		{intervals: [][]int{{1, 3}, {3, 4}}, expected: [][]int{{1, 4}}},
		{intervals: [][]int{{1, 3}, {2, 3}}, expected: [][]int{{1, 3}}},
		{intervals: [][]int{{1, 3}, {2, 2}}, expected: [][]int{{1, 3}}},
		{intervals: [][]int{{1, 2}, {3, 5}, {4, 8}, {8, 12}, {13, 15}}, expected: [][]int{{1, 2}, {3, 12}, {13, 15}}},
		{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, expected: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{intervals: [][]int{{1, 4}, {4, 5}}, expected: [][]int{{1, 5}}},
	}

	for _, test := range tests {
		ans := merge(test.intervals)
		if !check(ans, test.expected) {
			t.Fatalf("case %v expected %v not %v", test.intervals, test.expected, ans)
		}
	}
}
