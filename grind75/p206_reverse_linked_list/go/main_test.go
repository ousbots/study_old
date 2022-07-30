package reverse_linked_list

import "testing"

func TestProvided(t *testing.T) {
	input := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	expected := &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}
	if ans := reverseList(input); !comp(ans, expected) {
		t.Fatalf("%v is not %v", ans, expected)
	}

	input = &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	expected = &ListNode{Val: 2, Next: &ListNode{Val: 1}}
	if ans := reverseList(input); !comp(ans, expected) {
		t.Fatalf("%v is not %v", ans, expected)
	}

	if ans := reverseList(nil); !comp(ans, nil) {
		t.Fatalf("%v is not %v", ans, expected)
	}
}

func TestBasic(t *testing.T) {
	input := &ListNode{Val: 1}
	expected := &ListNode{Val: 1}
	if ans := reverseList(input); !comp(ans, expected) {
		t.Fatalf("%v is not %v", ans, expected)
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
