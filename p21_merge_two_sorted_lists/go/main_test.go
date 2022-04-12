package merge_two_sorted_lists

import "testing"

func TestProvided(t *testing.T) {
	output := mergeTwoLists(toNode([]int{1, 2, 4}), toNode([]int{1, 3, 4}))
	if !comp(output, toNode([]int{1, 1, 2, 3, 4, 4})) {
		t.Fatalf("expected [1 1 2 3 4 4] not %v", toSlice(output))
	}

	output = mergeTwoLists(toNode([]int{}), toNode([]int{}))
	if !comp(output, toNode([]int{})) {
		t.Fatalf("expected [] not %v", toSlice(output))
	}

	output = mergeTwoLists(toNode([]int{}), toNode([]int{0}))
	if !comp(output, toNode([]int{0})) {
		t.Fatalf("expected [] not %v", toSlice(output))
	}
}

func TestBasic(t *testing.T) {
	output := mergeTwoLists(toNode([]int{1, 2, 3}), toNode([]int{}))
	if !comp(output, toNode([]int{1, 2, 3})) {
		t.Fatalf("expected [1 2 3] not %v", toSlice(output))
	}

	output = mergeTwoLists(toNode([]int{}), toNode([]int{1, 2, 3}))
	if !comp(output, toNode([]int{1, 2, 3})) {
		t.Fatalf("expected [1 2 3] not %v", toSlice(output))
	}

	output = mergeTwoLists(toNode([]int{1, 2, 6, 7, 8}), toNode([]int{3, 4, 5}))
	if !comp(output, toNode([]int{1, 2, 3, 4, 5, 6, 7, 8})) {
		t.Fatalf("expected [1 2 3 4 5 6 7 8] not %v", toSlice(output))
	}

	output = mergeTwoLists(toNode([]int{1, 2, 6, 7}), toNode([]int{3, 4, 5, 8}))
	if !comp(output, toNode([]int{1, 2, 3, 4, 5, 6, 7, 8})) {
		t.Fatalf("expected [1 2 3 4 5 6 7 8] not %v", toSlice(output))
	}
}

func TestProvidedSimple(t *testing.T) {
	output := mergeTwoListsSimple(toNode([]int{1, 2, 4}), toNode([]int{1, 3, 4}))
	if !comp(output, toNode([]int{1, 1, 2, 3, 4, 4})) {
		t.Fatalf("expected [1 1 2 3 4 4] not %v", toSlice(output))
	}

	output = mergeTwoListsSimple(toNode([]int{}), toNode([]int{}))
	if !comp(output, toNode([]int{})) {
		t.Fatalf("expected [] not %v", toSlice(output))
	}

	output = mergeTwoListsSimple(toNode([]int{}), toNode([]int{0}))
	if !comp(output, toNode([]int{0})) {
		t.Fatalf("expected [] not %v", toSlice(output))
	}
}

func TestBasicSimple(t *testing.T) {
	output := mergeTwoListsSimple(toNode([]int{1, 2, 3}), toNode([]int{}))
	if !comp(output, toNode([]int{1, 2, 3})) {
		t.Fatalf("expected [1 2 3] not %v", toSlice(output))
	}

	output = mergeTwoListsSimple(toNode([]int{}), toNode([]int{1, 2, 3}))
	if !comp(output, toNode([]int{1, 2, 3})) {
		t.Fatalf("expected [1 2 3] not %v", toSlice(output))
	}

	output = mergeTwoListsSimple(toNode([]int{1, 2, 6, 7, 8}), toNode([]int{3, 4, 5}))
	if !comp(output, toNode([]int{1, 2, 3, 4, 5, 6, 7, 8})) {
		t.Fatalf("expected [1 2 3 4 5 6 7 8] not %v", toSlice(output))
	}

	output = mergeTwoListsSimple(toNode([]int{1, 2, 6, 7}), toNode([]int{3, 4, 5, 8}))
	if !comp(output, toNode([]int{1, 2, 3, 4, 5, 6, 7, 8})) {
		t.Fatalf("expected [1 2 3 4 5 6 7 8] not %v", toSlice(output))
	}
}

func toNode(values []int) *ListNode {
	head := &ListNode{}
	current := head

	for _, val := range values {
		current.Next = &ListNode{
			Val:  val,
			Next: nil,
		}

		current = current.Next
	}

	return head.Next
}

func toSlice(head *ListNode) []int {
	values := []int{}
	current := head

	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}

	return values
}

func comp(first, second *ListNode) bool {
	for first != nil && second != nil {
		if first.Val != second.Val {
			return false
		}

		first = first.Next
		second = second.Next
	}

	if first == nil && second == nil {
		return true
	}

	return false
}
