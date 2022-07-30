package fruit_into_baskets

// Finds the maximimum number of fruits that can be picked following a few rules. The elements in
// the fruits slice are the type of the tree at that position. The rules for picking are:
//   * Two baskets that can each hold only one type of fruit.
//   * Starting at any tree, move to the right while picking exactly one fruit from each tree.
//   * Stop if a tree is reached that cannot be picked due to the first rule.
//
// The solution uses a sliding window where the front of the window moves forward every cycle,
// updating the count map. If the map has more than two elements in it, the rear of the window is
// moved forward, updating the count map in reverse until there are only two elements left. The
// max difference between the two window edges is maintained and returned.
func totalFruit(fruits []int) int {
	counts := make(map[int]int)
	max := 0

	j := 0
	for i := range fruits {
		counts[fruits[i]]++

		for len(counts) > 2 {
			counts[fruits[j]]--
			if counts[fruits[j]] == 0 {
				delete(counts, fruits[j])
			}

			j++
		}

		total := i - j + 1
		if total > max {
			max = total
		}
	}

	return max
}

// First poor attempt at a sliding window.
func totalFruitFirstBad(fruits []int) int {
	if len(fruits) <= 2 {
		return len(fruits)
	}

	bt1 := fruits[0]
	bc1 := 1
	bt2 := 0
	bc2 := 0

	p1 := 1
	p2 := 0

	// Determine the first two basket types
	for fruits[p1] == bt1 {
		bc1++
		p1++
	}
	bt2 = fruits[p1]
	bc2++

	max := p1 - p2 + 1
	p1++

	for p1 < len(fruits) {
		switch fruits[p1] {
		case bt1:
			bc1++

		case bt2:
			bc2++

		default:
			for p2 < p1 {
				if fruits[p2] == bt1 {
					bc1--
				} else {
					bc2--
				}
				p2++

				if bc1 == 0 {
					bt1 = fruits[p1]
					bc1 = 1
					break
				}

				if bc2 == 0 {
					bt2 = fruits[p1]
					bc2 = 1
					break
				}
			}
		}
		p1++

		if p1-p2+1 > max {
			max = p1 - p2 + 1
		}
	}

	return max
}
