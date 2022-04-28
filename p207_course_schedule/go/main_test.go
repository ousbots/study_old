package course_schedule

import "testing"

func TestProvided(t *testing.T) {
	if canFinish(2, [][]int{{0, 1}}) != true {
		t.Fatal("expected true")
	}

	if canFinish(2, [][]int{{1, 0}, {0, 1}}) != false {
		t.Fatal("expected false")
	}
}

func TestBasic(t *testing.T) {
	if canFinish(1, [][]int{}) != true {
		t.Fatal("expected true")
	}

	if canFinish(3, [][]int{{2, 1}, {1, 0}, {0, 2}}) != false {
		t.Fatal("expected false")
	}

	if canFinish(20, [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 1}, {13, 1}, {15, 1}, {17, 4}}) != false {
		t.Fatal("expected false")
	}

	if canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}) != true {
		t.Fatal("expected true")
	}
}
