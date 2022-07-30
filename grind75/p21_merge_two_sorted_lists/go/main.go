package merge_two_sorted_lists

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head

	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				current.Next = list1
				list1 = list1.Next
			} else {
				current.Next = list2
				list2 = list2.Next
			}
		} else if list1 != nil {
			current.Next = list1
			list1 = list1.Next
		} else if list2 != nil {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	return head.Next
}

func mergeTwoListsSimple(list1 *ListNode, list2 *ListNode) *ListNode {
	values1 := []int{}
	values2 := []int{}

	for list1 != nil {
		values1 = append(values1, list1.Val)
		list1 = list1.Next
	}

	for list2 != nil {
		values2 = append(values2, list2.Val)
		list2 = list2.Next
	}

	final := append(values1, values2...)
	sort.Slice(final, func(i, j int) bool { return final[i] < final[j] })

	head := &ListNode{}
	current := head
	for _, val := range final {
		current.Next = &ListNode{Val: val, Next: nil}
		current = current.Next

	}

	return head.Next
}
