package reorder_list

import "testing"
import "fmt"

type testCase struct {
	head     *ListNode
	expected *ListNode
}

// Generates a linked list from the given values and the loop position and returns the
// corresponding test case.
func genCase(values, answer []int) testCase {
	nodes := []*ListNode{}
	for i, value := range values {
		nodes = append(nodes, &ListNode{Val: value})
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}

	expected := []*ListNode{}
	for i, value := range answer {
		expected = append(expected, &ListNode{Val: value})
		if i > 0 {
			expected[i-1].Next = expected[i]
		}
	}

	if len(nodes) == 0 {
		return testCase{head: nil, expected: nil}
	}

	return testCase{head: nodes[0], expected: expected[0]}
}

func toList(a *ListNode) []int {
	list := []int{}
	temp := a
	for temp != nil {
		list = append(list, temp.Val)
		temp = temp.Next
	}

	return list
}

func checkList(a, b *ListNode) bool {
	fmt.Printf("checking\n")
	for a != nil && b != nil {
		fmt.Printf("loop\n")
		if a.Val != b.Val {
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
		genCase([]int{1, 2, 3, 4}, []int{1, 4, 2, 3}),
		genCase([]int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}),
		genCase([]int{1, 2}, []int{1, 2}),
		genCase([]int{1, 2, 3}, []int{1, 3, 2}),
	}

	for _, test := range tests {
		fmt.Printf("test\n")
		orig := toList(test.head)
		expected := toList(test.expected)
		fmt.Printf("toList\n")
		reorderList(test.head)
		ans := toList(test.head)
		fmt.Printf("reorder\n")
		if !checkList(test.head, test.expected) {
			t.Fatalf("case %v expected %v not %v", orig, expected, ans)
		}
	}
}
