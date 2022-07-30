package main

import ("testing")

func TestProvided(t *testing.T) {
	ans := countRangeSum([]int{-2, 5, -1}, -2, 2)
	if ans != 3 {
		t.Fatalf("%d incorrect, expected 3", ans)
	}
	
	ans = countRangeSum([]int{0}, 0, 0)
	if ans != 1 {
		t.Fatalf("%d incorrect, expected 1", ans)
	}

	ans = countRangeSum([]int{0, 0}, 0, 0)
	if ans != 3 {
		t.Fatalf("%d incorrect, expected 3", ans)
	}
}
