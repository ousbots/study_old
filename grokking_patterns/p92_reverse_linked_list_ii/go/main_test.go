package reverse_linked_list_ii

import "testing"

type testCase struct {
	input    []int
	left     int
	right    int
	expected []int
}

// Returns the head of a linked list generated from the input values.
func toList(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}

	nodes := make([]*ListNode, len(input))

	for i := range input {
		nodes[i] = &ListNode{Val: input[i]}
	}

	for i := range nodes {
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}

	return nodes[0]
}

// Returns a slice of the values in the linked list.
func toSlice(head *ListNode) []int {
	values := []int{}
	temp := head

	for temp != nil {
		values = append(values, temp.Val)
		temp = temp.Next
	}

	return values
}

// Returns whether the two arrays have the same values and order.
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
		{input: []int{1, 2, 3, 4, 5}, left: 2, right: 4, expected: []int{1, 4, 3, 2, 5}},
		{input: []int{5}, left: 1, right: 1, expected: []int{5}},
		{input: []int{1, 2}, left: 1, right: 2, expected: []int{2, 1}},
	}

	for _, test := range tests {
		ans := toSlice(reverseBetween(toList(test.input), test.left, test.right))
		if !comp(ans, test.expected) {
			t.Fatalf("case input=%v left=%d right=%d expected %v not %v", test.input, test.left, test.right, test.expected, ans)
		}
	}
}
