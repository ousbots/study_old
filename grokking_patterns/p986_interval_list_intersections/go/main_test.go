package interval_list_intersections

import "testing"

type testCase struct {
	firstList  [][]int
	secondList [][]int
	expected   [][]int
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
		{firstList: [][]int{{1, 5}, {7, 9}, {10, 13}}, secondList: [][]int{{2, 3}, {5, 5}, {6, 11}, {13, 15}}, expected: [][]int{{2, 3}, {5, 5}, {7, 9}, {10, 11}, {13, 13}}},
		{firstList: [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, secondList: [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}, expected: [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}}},
		{firstList: [][]int{{1, 3}, {5, 9}}, secondList: [][]int{}, expected: [][]int{}},
		{firstList: [][]int{{1, 3}, {5, 9}}, secondList: [][]int{{0, 0}, {4, 4}, {10, 20}}, expected: [][]int{}},
	}

	for _, test := range tests {
		ans := intervalIntersection(test.firstList, test.secondList)
		if !comp(ans, test.expected) {
			t.Fatalf("case firstList=%v secondList=%v expected %v not %v", test.firstList, test.secondList, test.expected, ans)
		}
	}
}
