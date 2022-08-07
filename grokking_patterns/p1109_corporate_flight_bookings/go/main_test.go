package corporate_flight_bookings

import "testing"

type testCase struct {
	bookings [][]int
	n        int
	expected []int
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

func TestBasic(t *testing.T) {
	tests := []testCase{
		{bookings: [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, n: 5, expected: []int{10, 55, 45, 25, 25}},
		{bookings: [][]int{{1, 2, 10}, {2, 2, 15}}, n: 2, expected: []int{10, 25}},
		{bookings: [][]int{{1, 2, 10}, {2, 2, 15}}, n: 2, expected: []int{10, 25}},
		{bookings: [][]int{{1, 2, 10}, {4, 5, 9}, {5, 10, 1}}, n: 11, expected: []int{10, 10, 0, 9, 10, 1, 1, 1, 1, 1, 0}},
	}

	for _, test := range tests {
		ans := corpFlightBookings(test.bookings, test.n)
		if !comp(test.expected, ans) {
			t.Fatalf("case bookings=%v n=%d expected %v not %v", test.bookings, test.n, test.expected, ans)
		}
	}
}
