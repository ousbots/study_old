package reverse_linked_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

// Reverse the nodes of the given linked list between positions left and right inclusive.
//
// This solution iterates over the list once, up to the start of the reverse interval, recording
// the before node, then iterates over the interval, reversing the nodes in the interval, then
// sets the node at the beginning of the interval to point after the interval and if there is a
// node before the interval it points to the last node in the interval. If the reverse interval
// starts at the beginning of the list, then the final node in the interval is returned as the
// head, otherwise the original head is returned.
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	var before *ListNode
	current := head

	for i := 1; i < left; i++ {
		before = current
		current = current.Next
	}

	start := current
	tempB := current
	tempF := tempB.Next

	for i := left; i < right; i++ {
		current = tempF.Next
		tempF.Next = tempB
		tempB = tempF
		tempF = current
	}

	start.Next = current

	if left == 1 {
		return tempB
	}

	before.Next = tempB

	return head
}

// Reverse the nodes of the given linked list between positions left and right inclusive.
//
// This solution first finds the nodes that are before, start, end, and after the reverse interval.
// Then the interval is reversed and the reversed section "inserted" into the surrounding nodes.
// This solution iterates the list twice, once to determine the nodes of interest and another time
// to reverse the nodes in the interval, so isn't as fast as it could be.
func reverseBetweenSlow(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	if left == right {
		return head
	}

	var start, end, before, after *ListNode

	pos := 1
	temp := head
	for temp != nil {
		if pos+1 == left {
			before = temp
		}

		if pos == left {
			start = temp
		}

		if pos == right {
			end = temp
		}

		if pos-1 == right {
			after = temp
			break
		}

		temp = temp.Next
		pos++
	}

	tempB := start
	tempF := start.Next
	for tempF != nil && tempF != after {
		next := tempF.Next
		tempF.Next = tempB
		tempB = tempF
		tempF = next
	}

	start.Next = after
	if before != nil {
		before.Next = end
	}

	if left == 1 {
		return end
	}

	return head
}
