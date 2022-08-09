package rotate_list

import "testing"

type testCase struct {
	input    []int
	k        int
	expected []int
}

func toList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	nodes := make([]*ListNode, len(vals))

	for i := range vals {
		nodes[i] = &ListNode{Val: vals[i]}
	}

	for i := range nodes {
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}

	return nodes[0]
}

func toSlice(head *ListNode) []int {
	vals := []int{}
	temp := head

	for temp != nil {
		vals = append(vals, temp.Val)
		temp = temp.Next
	}

	return vals
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
		{input: []int{}, k: 4, expected: []int{}},
		{input: []int{1, 2, 3, 4, 5}, k: 2, expected: []int{4, 5, 1, 2, 3}},
		{input: []int{0, 1, 2}, k: 4, expected: []int{2, 0, 1}},
		{input: []int{0, 1, 2, 3}, k: 4, expected: []int{0, 1, 2, 3}},
		{input: []int{0, 1, 2, 3}, k: 0, expected: []int{0, 1, 2, 3}},
		{input: []int{0, 1, 2, 3}, k: 5, expected: []int{3, 0, 1, 2}},
		{input: []int{0, 1, 2, 3}, k: 1, expected: []int{3, 0, 1, 2}},
	}

	for _, test := range tests {
		ans := toSlice(rotateRight(toList(test.input), test.k))
		if !comp(ans, test.expected) {
			t.Fatalf("case input=%v k=%d expected %v not %v", test.input, test.k, test.expected, ans)
		}
	}
}
