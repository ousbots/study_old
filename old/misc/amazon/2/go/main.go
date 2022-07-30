package amazon

func findMaxProducts(products []int32) int64 {
	return find(products, int32(len(products)-1), 0)
}

func find(products []int32, index, start int32) int64 {
	if index < 0 {
		return 0
	}

	total := int64(0)
	diff := int64(start)
	for i := index; i >= 0; i-- {
		if int64(products[i]) < diff {
			break
		}

		total += diff
		diff += 1
	}

	first := find(products, index, start+1)
	second := find(products, index-1, start)

	max := total
	if first > max {
		max = first
	}
	if second > max {
		max = second
	}

	return max
}