package linked_list_cycle

import "testing"
import "fmt"

func TestProvided(t *testing.T) {
	input := toLL([]int{3, 2, 0, -4}, 1)
	fmt.Printf("input %v\n", input)
	if hasCycle(input) != true {
		t.Fatal("expected true")
	}

	input = toLL([]int{1, 2}, 0)
	if hasCycle(input) != true {
		t.Fatal("expected true")
	}

	input = toLL([]int{1}, -1)
	if hasCycle(input) != false {
		t.Fatal("expected false")
	}
}

func toLL(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	root := ListNode{Val: values[0]}
	current := &root
	values = values[1:]

	for _, val := range values {
		node := &ListNode{Val: val}
		current.Next = node
		current = node
	}

	if pos >= 0 {
		link := &root
		for pos > 1 && link != nil {
			link = link.Next
			pos--
		}

		current.Next = link
	}

	return &root
}
