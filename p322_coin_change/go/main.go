package coin_change

import "sort"
import "math"

// The greedy algorithm doesn't always work for arbitrary coin combinations.
func coinChangeGreedy(coins []int, amount int) int {
	sort.Slice(coins, func(i, j int) bool { return coins[i] < coins[j] })
	count := 0

	for amount > 0 {
		if len(coins) == 0 {
			return -1
		}

		coin := coins[len(coins)-1]
		coins = coins[:len(coins)-1]

		for amount >= coin {
			amount -= coin
			count++
		}
	}

	return count
}

// Dynamic programming, bottom-up approach: builds a table of minimum number of coins from one to
// the amount, referencing previous values while calculating new ones.
func coinChange(coins []int, amount int) int {
	value := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		min := math.MaxInt32

		for _, coin := range coins {
			if i-coin >= 0 && value[i-coin] != -1 {
				temp := value[i-coin] + 1
				if temp < min {
					min = temp
				}
			}
		}

		if min == math.MaxInt32 {
			value[i] = -1
		} else {
			value[i] = min
		}
	}

	return value[amount]
}
