package minimum_height_trees

// This finds the minimum height tree nodes from a given number of nodes and the edge list. This
// function creates an adjecency list, then finds the furthest node from an arbitrary node, the
// first node in this case, then finds the furthest node from that one, which is the diameter of
// the tree and the middle of the path in the diameter are the roots of the minimum height nodes.
//
// See https://en.wikipedia.org/wiki/Centered_tree for more information.
func findMinHeightTrees(n int, edges [][]int) []int {
	adjacent := make(map[int][]int, n)

	for _, edge := range edges {
		adjacent[edge[0]] = append(adjacent[edge[0]], edge[1])
		adjacent[edge[1]] = append(adjacent[edge[1]], edge[0])
	}

	a := 0
	b, _ := farthest(a, adjacent, make([]bool, n))
	_, path := farthest(b, adjacent, make([]bool, n))

	if len(path)%2 == 0 {
		return []int{path[(len(path)-1)/2], path[(len(path)-1)/2+1]}
	} else {
		return []int{path[(len(path)-1)/2]}
	}
}

// Finds the farthest away node from the given root via DFS. Returns a tuple of (node, path), which
// are the farthest away node and the path between the given rooth and that node. The boolean array
// seen should initially be all false and have a length of n.
func farthest(root int, adjacent map[int][]int, seen []bool) (int, []int) {
	seen[root] = true
	children := adjacent[root]

	deepest := root
	path := []int{}

	for _, child := range children {
		if !seen[child] {
			node, temp := farthest(child, adjacent, seen)

			if len(temp) > len(path) {
				deepest = node
				path = temp
			}
		}
	}

	return deepest, append(path, root)
}

//
// BAD, BAD, NOT GOOD.
//

type Node struct {
	value    int
	children []*Node
}

type Edge struct {
	first  int
	second int
}

func findMinHeightTreesSlow(n int, edges [][]int) []int {
	seen := make(map[Edge]struct{})
	nodes := make([]*Node, n)

	for _, edge := range edges {
		if len(edge) != 2 {
			continue
		}

		normal := Edge{edge[0], edge[1]}
		reverse := Edge{edge[1], edge[0]}

		if _, exists := seen[normal]; exists {
			continue
		}

		if _, exists := seen[reverse]; exists {
			continue
		}

		seen[normal] = struct{}{}
		seen[reverse] = struct{}{}

		first := nodes[edge[0]]
		if first == nil {
			first = &Node{edge[0], nil}
			nodes[edge[0]] = first
		}

		second := nodes[edge[1]]
		if second == nil {
			second = &Node{edge[1], nil}
			nodes[edge[1]] = second
		}

		first.children = append(first.children, second)
		second.children = append(second.children, first)
	}

	roots := []int(nil)
	smallest := n

	for root := 0; root < n; root++ {
		h := heightSlow(nodes[root], make([]bool, n))

		if h < smallest {
			roots = []int{root}
			smallest = h
		} else if h == smallest {
			roots = append(roots, root)
		}
	}

	return roots
}

func heightSlow(root *Node, seen []bool) int {
	if root == nil {
		return 0
	}

	seen[root.value] = true

	deepest := 0
	for _, child := range root.children {
		if seen[child.value] {
			continue
		}

		depth := heightSlow(child, seen)
		if depth > deepest {
			deepest = depth
		}
	}

	return deepest + 1
}
