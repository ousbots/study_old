package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	l := 0
	r := 1
	max := 0

	for r < len(prices) {
		profit := prices[r] - prices[l]
		if profit > 0 {
			if profit > max {
				max = profit
			}
		} else {
			l = r
		}

		r++
	}

	return max
}
