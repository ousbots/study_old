package reverse_linked_list

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) String() string {
	s := "["
	curr := &l

	max := 99
	for curr != nil {
		max--
		if max == 0 {
			break
		}

		s += fmt.Sprintf(" %d", curr.Val)
		curr = curr.Next
	}
	s += "]"

	return s
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	rev := &ListNode{Val: head.Val}
	node := head.Next

	for node != nil {
		rev = &ListNode{Val: node.Val, Next: rev}
		node = node.Next
	}

	return rev
}

func reverseListDFS(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	nodes := []*ListNode{}
	node := head

	for node != nil {
		nodes = append(nodes, node)
		node = node.Next
	}

	for i := len(nodes) - 1; i > 0; i-- {
		nodes[i].Next = nodes[i-1]
	}
	nodes[0].Next = nil

	return nodes[len(nodes)-1]
}
