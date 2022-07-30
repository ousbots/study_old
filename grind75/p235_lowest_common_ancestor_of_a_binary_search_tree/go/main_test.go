package lowest_common_ancestor_of_a_binary_search_tree

import (
	"strconv"
	"testing"
)

func TestProvided(t *testing.T) {
	input := toTree([]string{"6", "2", "8", "0", "4", "7", "9", "null", "null", "3", "5"})
	if ans := lowestCommonAncestor(input, &TreeNode{Val: 2}, &TreeNode{Val: 8}); ans == nil || ans.Val != 6 {
		t.Fatalf("expected 6, not %v\n", ans.Val)
	}

	input = toTree([]string{"6", "2", "8", "0", "4", "7", "9", "null", "null", "3", "5"})
	if ans := lowestCommonAncestor(input, &TreeNode{Val: 2}, &TreeNode{Val: 4}); ans == nil || ans.Val != 2 {
		t.Fatalf("expected 2, not %v\n", ans.Val)
	}

	input = toTree([]string{"2", "1"})
	if ans := lowestCommonAncestor(input, &TreeNode{Val: 2}, &TreeNode{Val: 1}); ans == nil || ans.Val != 2 {
		t.Fatalf("expected 2, not %v\n", ans.Val)
	}
}

func TestBasic(t *testing.T) {
	input := toTree([]string{"6", "2", "8", "0", "4", "7", "9", "null", "null", "3", "5"})
	if ans := lowestCommonAncestor(input, &TreeNode{Val: 0}, &TreeNode{Val: 4}); ans == nil || ans.Val != 2 {
		t.Fatalf("expected 2, not %v\n", ans.Val)
	}

	if ans := lowestCommonAncestor(input, &TreeNode{Val: 4}, &TreeNode{Val: 3}); ans == nil || ans.Val != 4 {
		t.Fatalf("expected 4, not %v\n", ans.Val)
	}

	if ans := lowestCommonAncestor(input, &TreeNode{Val: 5}, &TreeNode{Val: 4}); ans == nil || ans.Val != 4 {
		t.Fatalf("expected 4, not %v\n", ans.Val)
	}

	if ans := lowestCommonAncestor(input, &TreeNode{Val: 9}, &TreeNode{Val: 0}); ans == nil || ans.Val != 6 {
		t.Fatalf("expected 6, not %v\n", ans.Val)
	}

	if ans := lowestCommonAncestor(input, &TreeNode{Val: 0}, &TreeNode{Val: 3}); ans == nil || ans.Val != 2 {
		t.Fatalf("expected 2, not %v\n", ans.Val)
	}
}

func compareTrees(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return a.Val == b.Val && compareTrees(a.Left, b.Left) && compareTrees(a.Right, b.Right)
}

func toTree(values []string) *TreeNode {
	if len(values) == 0 || values[0] == "null" {
		return nil
	}

	val, err := strconv.Atoi(values[0])
	if err != nil {
		panic(err)
	}
	values = values[1:]

	root := &TreeNode{Val: val}
	parents := []*TreeNode{root}

loop:
	for len(parents) != 0 {
		temp := []*TreeNode{}

		for _, node := range parents {
			if len(values) == 0 {
				break loop
			}

			data := values[0]
			values = values[1:]

			if data != "null" {
				val, err := strconv.Atoi(data)
				if err != nil {
					panic(err)
				}
				node.Left = &TreeNode{Val: val}
				temp = append(temp, node.Left)
			}

			if len(values) == 0 {
				break loop
			}

			data = values[0]
			values = values[1:]

			if data != "null" {
				val, err := strconv.Atoi(data)
				if err != nil {
					panic(err)
				}
				node.Right = &TreeNode{Val: val}
				temp = append(temp, node.Right)
			}
		}

		parents = temp
	}

	return root
}

func toSlice(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	values := []int{}
	parents := []*TreeNode{root}

	for len(parents) != 0 {
		temp := []*TreeNode{}

		for _, node := range parents {
			values = append(values, node.Val)
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}

		parents = temp
	}

	return values
}

func TestToTree(t *testing.T) {
	root := toTree([]string{})
	if root != nil {
		t.Fatalf("expected nil, not %v", toSlice(root))
	}

	root = toTree([]string{"0"})
	expected := &TreeNode{Val: 0}
	if !compareTrees(root, expected) {
		t.Fatalf("expected %v, not %v", toSlice(expected), toSlice(root))
	}

	root = toTree([]string{"0", "1"})
	expected = &TreeNode{Val: 0, Left: &TreeNode{Val: 1}}
	if !compareTrees(root, expected) {
		t.Fatalf("expected %v, not %v", toSlice(expected), toSlice(root))
	}

	root = toTree([]string{"0", "1", "2"})
	expected = &TreeNode{Val: 0, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
	if !compareTrees(root, expected) {
		t.Fatalf("expected %v, not %v", toSlice(expected), toSlice(root))
	}

	root = toTree([]string{"0", "1", "2", "null", "null", "3"})
	expected = &TreeNode{Val: 0, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	if !compareTrees(root, expected) {
		t.Fatalf("expected %v, not %v", toSlice(expected), toSlice(root))
	}
}
