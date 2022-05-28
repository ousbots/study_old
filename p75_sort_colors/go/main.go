package p75_sort_colors

func sortColors(nums []int) {
	white := 0
	blue := 0
	index := 0

	for _, num := range nums {
		switch num {
		case 0:
			nums[index] = 0
			index++
		case 1:
			white++
		case 2:
			blue++
		}
	}

	for i := index; i < index+white; i++ {
		nums[i] = 1
	}

	for i := index + white; i < len(nums); i++ {
		nums[i] = 2
	}
}
