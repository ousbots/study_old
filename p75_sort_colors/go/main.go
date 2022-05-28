package p75_sort_colors

func sortColors(nums []int) {
	red := 0
	white := 0

	for _, num := range nums {
		switch num {
		case 0:
			red++
		case 1:
			white++
		}
	}

	for i := range nums {
		if i < red+white {
			if i < red {
				nums[i] = 0
			} else {
				nums[i] = 1
			}
		} else {
			nums[i] = 2
		}
	}
}
