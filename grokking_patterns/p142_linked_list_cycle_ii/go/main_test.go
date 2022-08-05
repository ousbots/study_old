package linked_list_cycle_ii

import "testing"

type testCase struct {
	head     *ListNode
	expected *ListNode
}

// Generates a linked list from the given values and the loop position and returns the
// corresponding test case.
func genCase(values []int, pos int) testCase {
	nodes := []*ListNode{}
	for i, value := range values {
		nodes = append(nodes, &ListNode{Val: value})
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}

	var expected *ListNode
	if pos >= 0 {
		expected = nodes[pos]
		nodes[len(nodes)-1].Next = expected
	}

	if len(nodes) == 0 {
		return testCase{head: nil, expected: nil}
	}

	return testCase{head: nodes[0], expected: expected}
}

func TestBasic(t *testing.T) {
	tests := []testCase{
		genCase([]int{3, 2, 0, -4}, 1),
		genCase([]int{1, 2}, 0),
		genCase([]int{1}, 0),
		genCase([]int{1}, -1),
		genCase([]int{}, -1),
	}

	for _, test := range tests {
		ans := detectCycle(test.head)
		if ans != test.expected {
			t.Fatalf("case %v expected %v not %v", test.head, test.expected, ans)
		}
	}
}
