package reorder_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	end := head
	middle := head

	for end.Next != nil && end.Next.Next != nil {
		middle, end = middle.Next, end.Next.Next
	}

	if end.Next != nil {
		middle, end = middle.Next, end.Next
	}

	tempB, tempF := middle, middle.Next
	for tempF != nil {
		next := tempF.Next
		tempF.Next = tempB
		tempB = tempF
		tempF = next
	}

	front := head

	for front != middle {
		next := front.Next
		front.Next = end
		end = end.Next
		front.Next.Next = next
		front = next
	}
	middle.Next = nil
}
