package minimum_height_trees

import (
	"sort"
	"testing"
)

func checkSlices(a, b []int) bool {
	sort.Ints(a)
	sort.Ints(b)

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestFarthestBasic(t *testing.T) {
	adjacent := map[int][]int{0: {1}, 1: {0, 2, 3}, 2: {1}, 3: {1}}
	node, path := farthest(0, adjacent, make([]bool, 4))
	switch node {
	case 0:
		t.Fatal("0 is not farthest from itself")

	case 1:
		expected := []int{0, 1}
		if !checkSlices(path, expected) {
			t.Fatalf("1 farthest expected %v, not %v", expected, path)
		}

	case 2:
		expected := []int{0, 1, 2}
		if !checkSlices(path, expected) {
			t.Fatalf("2 farthest expected %v, not %v", expected, path)
		}

	case 3:
		expected := []int{0, 1, 3}
		if !checkSlices(path, expected) {
			t.Fatalf("3 farthest expected %v, not %v", expected, path)
		}

	default:
		t.Fatalf("%d is not a valid node and %v is not a valid path", node, path)
	}

	node, path = farthest(1, adjacent, make([]bool, 4))
	switch node {
	case 0:
		expected := []int{0, 1}
		if !checkSlices(path, expected) {
			t.Fatalf("0 farthest expected %v, not %v", expected, path)
		}

	case 1:
		t.Fatal("1 is not farthest from itself")

	case 2:
		expected := []int{1, 2}
		if !checkSlices(path, expected) {
			t.Fatalf("2 farthest expected %v, not %v", expected, path)
		}

	case 3:
		expected := []int{1, 3}
		if !checkSlices(path, expected) {
			t.Fatalf("3 farthest expected %v, not %v", expected, path)
		}

	default:
		t.Fatalf("%d is not a valid node and %v is not a valid path", node, path)
	}

}

func TestBasic(t *testing.T) {
	ans := findMinHeightTrees(2, [][]int{{1, 0}})
	expected := []int{0, 1}
	if !checkSlices(ans, expected) {
		t.Fatalf("expected %v not %v", expected, ans)
	}
}

func TestProvided(t *testing.T) {
	ans := findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}})
	expected := []int{1}
	if !checkSlices(ans, expected) {
		t.Fatalf("expected %v not %v", expected, ans)
	}

	ans = findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}})
	expected = []int{3, 4}
	if !checkSlices(ans, expected) {
		t.Fatalf("expected %v not %v", expected, ans)
	}

}

func TestHelperHeightSlowBasic(t *testing.T) {
	root := (*Node)(nil)
	h := heightSlow(root, nil)
	ans := 0
	if h != ans {
		t.Fatalf("expected %d not %d", ans, h)
	}

	root = &Node{0, []*Node{{1, nil}, {2, nil}}}
	h = heightSlow(root, make([]bool, 3))
	ans = 2
	if h != ans {
		t.Fatalf("expected %d not %d", ans, h)
	}

	root = &Node{0, []*Node{{1, []*Node{{3, nil}}}, {2, nil}}}
	h = heightSlow(root, make([]bool, 4))
	ans = 3
	if h != ans {
		t.Fatalf("expected %d not %d", ans, h)
	}

	root = &Node{0, []*Node{{1, []*Node{{3, []*Node{{4, nil}}}}}, {2, []*Node{{5, []*Node{{6, nil}, {7, []*Node{{8, nil}}}}}}}}}
	h = heightSlow(root, make([]bool, 9))
	ans = 5
	if h != ans {
		t.Fatalf("expected %d not %d", ans, h)
	}
}

func TestHelperHeightSlowRotation(t *testing.T) {
	zero := Node{0, nil}
	one := Node{1, nil}
	two := Node{2, nil}
	three := Node{3, nil}
	four := Node{4, nil}
	five := Node{5, nil}
	six := Node{6, nil}

	zero.children = append(zero.children, &three)
	three.children = append(three.children, &zero)
	one.children = append(one.children, &three)
	three.children = append(three.children, &one)
	two.children = append(two.children, &three)
	three.children = append(three.children, &two)
	four.children = append(four.children, &three)
	three.children = append(three.children, &four)
	four.children = append(four.children, &five)
	five.children = append(five.children, &four)
	two.children = append(two.children, &six)
	six.children = append(six.children, &two)

	h := heightSlow(&zero, make([]bool, 7))
	if h != 4 {
		t.Fatalf("expected 0 root to be heightSlow 4 not %d", h)
	}

	h = heightSlow(&one, make([]bool, 7))
	if h != 4 {
		t.Fatalf("expected 1 root to be heightSlow 4 not %d", h)
	}

	h = heightSlow(&two, make([]bool, 7))
	if h != 4 {
		t.Fatalf("expected 2 root to be heightSlow 4 not %d", h)
	}

	h = heightSlow(&three, make([]bool, 7))
	if h != 3 {
		t.Fatalf("expected 3 root to be heightSlow 3 not %d", h)
	}

	h = heightSlow(&four, make([]bool, 7))
	if h != 4 {
		t.Fatalf("expected 4 root to be heightSlow 4 not %d", h)
	}

	h = heightSlow(&five, make([]bool, 7))
	if h != 5 {
		t.Fatalf("expected 5 root to be heightSlow 5 not %d", h)
	}

	h = heightSlow(&six, make([]bool, 7))
	if h != 5 {
		t.Fatalf("expected 7 root to be height 5 not %d", h)
	}
}
