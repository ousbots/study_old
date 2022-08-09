package rotate_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// Returns the linked list rotated right by k nodes.
//
// This solution iterates over the list to determine its length and the tail node. The real offset
// is calculated, then the list is iterated over to the split point (count - offset) and that tail
// is moved to the front of the list. This is very similar to the "slow" solution, just a little
// cleaner and faster.
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	tail, split, count := head, head, 1

	for tail.Next != nil {
		count++
		tail = tail.Next
	}

	offset := k % count
	for i := 1; i < count-offset; i++ {
		split = split.Next
	}

	tail.Next = head
	head = split.Next
	split.Next = nil

	return head
}

// Returns the linked list rotated right by k nodes.
//
// This solution iterates over the list and  calculates the length and the real offset, then
// iterates to the split point, and splits and moves the tail to the front.
func rotateRightSlow(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	length := 0
	temp := head
	for temp != nil {
		length++
		temp = temp.Next
	}

	offset := k % length
	if offset == 0 {
		return head
	}

	temp = head
	new_head := head
	for temp.Next != nil {
		if offset == 0 {
			new_head = new_head.Next
		} else {
			offset--
		}

		temp = temp.Next
	}

	new_head.Next, new_head = nil, new_head.Next
	temp.Next = head

	return new_head
}
