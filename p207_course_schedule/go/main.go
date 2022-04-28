package course_schedule

// Kahns algorithm
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int, numCourses)
	edges := make(map[int][]int, len(prerequisites))
	entry := []int{}

	for _, pre := range prerequisites {
		edges[pre[0]] = append(edges[pre[0]], pre[1])
		graph[pre[1]] = append(graph[pre[1]], pre[0])
	}

	for i := 0; i < numCourses; i++ {
		if _, exists := edges[i]; !exists {
			entry = append(entry, i)
		}
	}

	for len(entry) != 0 {
		node := entry[len(entry)-1]
		entry = entry[:len(entry)-1]

		for _, elem := range graph[node] {
			for i, val := range edges[elem] {
				if val == node {
					edges[elem] = append(edges[elem][:i], edges[elem][i+1:]...)
					break
				}
			}

			if len(edges[elem]) == 0 {
				entry = append(entry, elem)
			}
		}

		delete(graph, node)
	}

	for _, edges := range graph {
		if len(edges) != 0 {
			return false
		}
	}

	return true
}

type State int

const (
	None State = iota
	Seen
	Start
)

// DFS recursive
func canFinishRecursive(numCourses int, prerequisites [][]int) bool {
	state := make([]State, numCourses)
	courses := make([][]int, numCourses)

	for _, pre := range prerequisites {
		courses[pre[0]] = append(courses[pre[0]], pre[1])
	}

	for i := 0; i < numCourses; i++ {
		if !search(courses, state, i) {
			return false
		}
	}

	return true
}

func search(courses [][]int, state []State, name int) bool {
	if state[name] == Start {
		return false
	}

	if state[name] == Seen {
		return true
	}

	state[name] = Start
	for _, pre := range courses[name] {
		if !search(courses, state, pre) {
			return false
		}
	}

	state[name] = Seen
	return true
}
