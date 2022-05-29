package p416_partition_equal_subset_sum

import "testing"

func TestProvided(t *testing.T) {
	if ans := canPartition([]int{1, 5, 11, 5}); ans != true {
		t.Fatal("expected true")
	}

	if ans := canPartition([]int{1, 2, 3, 5}); ans != false {
		t.Fatal("expected false")
	}
}

func TestBasic(t *testing.T) {
	if ans := canPartition([]int{1}); ans != false {
		t.Fatal("expected false")
	}

	if ans := canPartition([]int{1, 1}); ans != true {
		t.Fatal("expected true")
	}

	if ans := canPartition([]int{1, 2, 5}); ans != false {
		t.Fatal("expected false")
	}

	if ans := canPartition([]int{1, 1, 1, 3}); ans != true {
		t.Fatal("expected true")
	}

	if ans := canPartition([]int{1, 3, 1, 1, 1, 3, 1, 1}); ans != true {
		t.Fatal("expected true")
	}
}
