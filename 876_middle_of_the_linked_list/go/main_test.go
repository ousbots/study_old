package middle_of_the_linked_list

import "testing"

func TestProvided(t *testing.T) {
	input := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	valid := &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	if ans := middleNode(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}
	valid = &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}
	if ans := middleNode(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func TestBasic(t *testing.T) {
	input := &ListNode{Val: 1}
	valid := &ListNode{Val: 1}
	if ans := middleNode(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	valid = &ListNode{Val: 2}
	if ans := middleNode(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	valid = &ListNode{Val: 2, Next: &ListNode{Val: 3}}
	if ans := middleNode(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	valid = &ListNode{Val: 3, Next: &ListNode{Val: 4}}
	if ans := middleNode(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func comp(a, b *ListNode) bool {
	for {
		if a == nil && b == nil {
			return true
		}

		if a == nil || b == nil {
			return false
		}

		if a.Val != b.Val {
			return false
		}

		a, b = a.Next, b.Next
	}
}
