package best_time_to_buy_and_sell_stock

import "testing"

func TestProvided(t *testing.T) {
	ans := maxProfit([]int{7, 1, 5, 3, 6, 4})
	if ans != 5 {
		t.Fatalf("expected 5 not %d", ans)
	}

	ans = maxProfit([]int{7, 6, 4, 3, 1})
	if ans != 0 {
		t.Fatalf("expected 0 not %d", ans)
	}
}

func TestBasic(t *testing.T) {
	ans := maxProfit([]int{1, 5, 2, 7, 3, 9})
	if ans != 8 {
		t.Fatalf("expected 8 not %d", ans)
	}
}
