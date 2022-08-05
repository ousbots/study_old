package linked_list_cycle_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

// Returns the node that is the start of the loop.
//
// This solution uses the fast and slow pointers to detect if there is a cycle. If there is a
// cycle then then one of the pointers is reset to the head. Both pointers are then incremented
// by one until they meet, which is the starting node of the loop.
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	for ok := true; ok; ok = fast != slow {
		if slow == nil || fast == nil || fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
