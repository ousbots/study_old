package middle_of_the_linked_list

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	output := "["

	if l != nil {
		output += strconv.Itoa(l.Val)
		l = l.Next
	}

	for {
		if l == nil {
			break
		}

		output += ", " + strconv.Itoa(l.Val)
		l = l.Next
	}

	output += "]"
	return output
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	middle := head
	end := head

	for end != nil && end.Next != nil {
		middle = middle.Next

		end = end.Next
		if end != nil {
			end = end.Next
		}
	}

	return middle
}
