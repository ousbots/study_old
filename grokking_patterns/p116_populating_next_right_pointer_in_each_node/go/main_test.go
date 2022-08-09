package populating_next_right_pointers_in_each_node

import "testing"

type testCase struct {
	head     *Node
	expected *Node
}

func comp(a, b *Node) bool {
	for a != nil && b != nil {
		if a.Val != b.Val || a.Next != b.Next {
			return false
		}

		a = a.Next
		b = b.Next
	}

	if a != b {
		return false
	}

	return true
}

func TestBasic(t *testing.T) {
	tests := []testCase{
		{head: &Node{Val: 1, Left: &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}}, Right: &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}}}, expected: &Node{Val: 1, Left: &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}}, Right: &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}}}},
		{head: &Node{Val: 1}, expected: &Node{Val: 1}},
		{head: nil, expected: nil},
	}

	tests[0].head.Left.Next = tests[0].head.Right
	tests[0].head.Left.Left.Next = tests[0].head.Left.Right
	tests[0].head.Left.Right.Next = tests[0].head.Right.Left
	tests[0].head.Right.Left.Next = tests[0].head.Right.Right

	for _, test := range tests {
		if !comp(connect(test.head), test.expected) {
			t.Fatal("case failed, use debugger to explore")
		}
	}
}
