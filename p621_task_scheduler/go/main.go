package task_scheduler

import "sort"
import "fmt"

// Calculate the number of cycles needed to complete all given tasks (uppercase letters) when there
// needs to be n cycles before the same task can be repeated again.
//
// The task(s) with the highest frequency need tasks-1 * n number of idle cycles, which can be
// filled in by any remaining tasks. Then the final count is any remaining idle cycles + the number
// of tasks.
func leastInterval(tasks []byte, n int) int {
	var count [26]int

	for _, task := range tasks {
		count[task-'A']++
	}
	sort.Ints(count[:])

	highestFreq := count[25] - 1
	idleCount := highestFreq * n
	for i := 24; i >= 0; i-- {
		if count[i] < highestFreq {
			idleCount -= count[i]
		} else {
			idleCount -= highestFreq
		}
	}

	if idleCount > 0 {
		return len(tasks) + idleCount
	} else {
		return len(tasks)
	}
}

// This was close, but just kind of junk.
func leastIntervalCrap(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	counts := make(map[byte]int)
	length := len(tasks)

	for _, task := range tasks {
		counts[task]++
	}

	names := []byte{}
	count := []int{}
	for key, val := range counts {
		names = append(names, key)
		count = append(count, val)
	}

	sort.Slice(count, func(a, b int) bool { return count[b] < count[a] })
	positions := make(map[int][]int)
	for i := range count {
		positions[count[i]] = append(positions[count[i]], i)
	}

	for key, value := range counts {
		position := positions[value][0]
		positions[value] = positions[value][1:]
		if len(positions[value]) == 0 {
			delete(positions, value)
		}

		names[position] = key
	}

	order := []byte{}
	i := 0
	offset := 0
	for completed := 0; completed < length; {
		if i >= len(names) {
			if i >= n-offset {
				for ; i >= n-offset; i-- {
					order = append(order, 0)
				}
			}

			offset = 0
			i = 0
		}

		if i >= n-offset {
			i = 0
			offset = 0
		}

		order = append(order, names[i])
		completed++

		counts[names[i]]--
		if counts[names[i]] == 0 {
			names = append(names[:i], names[i+1:]...)
			offset++
		} else {
			i++
		}
	}

	fmt.Printf("order %v\n", order)

	return len(order)
}
