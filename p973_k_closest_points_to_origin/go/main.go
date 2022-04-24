package k_closest_points_to_origin

import (
	"math/rand"
	"sort"
)

type Point struct {
	dist  int
	point []int
}

func kClosestSlow(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0]*points[i][0]+points[i][1]*points[i][1] < points[j][0]*points[j][0]+points[j][1]*points[j][1]
	})

	return points[:k]
}

func kClosestFuncs(points [][]int, k int) [][]int {
	qSort(points, 0, len(points)-1)

	return points[:k]
}

func qSort(points [][]int, low, high int) {
	if low >= high || low < 0 {
		return
	}

	p := partition(points, low, high)

	qSort(points, low, p-1)
	qSort(points, p+1, high)
}

func partition(points [][]int, low, high int) int {
	pivotVal := points[high][0]*points[high][0] + points[high][1]*points[high][1]
	pivot := low - 1

	for i := low; i < high; i++ {
		val := points[i][0]*points[i][0] + points[i][1]*points[i][1]
		if val <= pivotVal {
			pivot += 1

			points[pivot], points[i] = points[i], points[pivot]
		}
	}

	pivot += 1
	points[pivot], points[high] = points[high], points[pivot]

	return pivot
}

// Iterative form of quicksort that exits early if possible ... still not very fast.
func kClosest(points [][]int, k int) [][]int {
	start := 0
	end := len(points) - 1

	for start <= end-1 {
		pivot := start + rand.Intn(end-start+1)
		pivotVal := points[pivot][0]*points[pivot][0] + points[pivot][1]*points[pivot][1]
		points[start], points[pivot] = points[pivot], points[start]

		low := start + 1
		for high := low; high <= end; high++ {
			val := points[high][0]*points[high][0] + points[high][1]*points[high][1]
			if val <= pivotVal {
				points[low], points[high] = points[high], points[low]
				low++
			}
		}

		points[low-1], points[start] = points[start], points[low-1]

		if low == k {
			return points[:k]
		} else if low > k {
			end = low - 2
		} else {
			start = low
		}
	}

	return points[:k]
}
