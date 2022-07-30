pub struct Solution;

impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        if prices.len() < 2 {
            return 0;
        }

        let mut buy = 0;
        let mut sell = 1;
        let mut max = 0;

        while sell < prices.len() {
            let price = prices[sell] - prices[buy];

            if price > 0 {
                if price > max {
                    max = price;
                }
            } else {
                buy = sell;
            }

            sell += 1;
        }

        max
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(Solution::max_profit(vec![7, 1, 5, 3, 6, 4]), 5);
        assert_eq!(Solution::max_profit(vec![7, 6, 4, 3, 1]), 0);
    }

    #[test]
    fn basic() {
        assert_eq!(Solution::max_profit(vec![1, 4, 2, 6, 3, 9]), 8);
        assert_eq!(Solution::max_profit(vec![1]), 0);
        assert_eq!(Solution::max_profit(vec![1, 4, 2, 6, 0, 5]), 5);
    }
}
