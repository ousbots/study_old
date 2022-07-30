package permutations

import (
	"sort"
	"testing"
)

func TestBasic(t *testing.T) {
	valid := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	if ans := permute([]int{1, 2, 3}); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = [][]int{{0, 1}, {1, 0}}
	if ans := permute([]int{0, 1}); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = [][]int{{1}}
	if ans := permute([]int{1}); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func comp(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for _, aa := range a {
		sort.Slice(aa, func(i, j int) bool { return aa[i] < aa[j] })

		found := false
		index := 0

	inner:
		for _, bb := range b {
			if len(aa) != len(bb) {
				continue
			}

			sort.Slice(bb, func(i, j int) bool { return bb[i] < bb[j] })

			for i := range aa {
				if aa[i] != bb[i] {
					b[index] = bb
					index++
					continue inner
				}
			}

			found = true
		}

		if !found {
			return false
		}
	}

	return true
}
