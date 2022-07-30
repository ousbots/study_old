pub struct Solution;
impl Solution {
    pub fn coin_change(coins: Vec<i32>, amount: i32) -> i32 {
        let mut mins: Vec<i32> = vec![0; amount as usize + 1];

        for i in 1..=amount {
            let mut min_coins = amount + 1;

            for coin in &coins {
                let last = i - coin;
                if last >= 0 && mins[last as usize] != -1 {
                    min_coins = min_coins.min(mins[last as usize] + 1);
                }
            }

            if min_coins == amount + 1 {
                mins[i as usize] = -1;
            } else {
                mins[i as usize] = min_coins;
            }
        }

        mins[amount as usize]
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(Solution::coin_change(vec![1, 2, 5], 11), 3);
        assert_eq!(Solution::coin_change(vec![2], 3), -1);
        assert_eq!(Solution::coin_change(vec![1], 0), 0);
    }

    #[test]
    fn basic() {
        assert_eq!(Solution::coin_change(vec![1, 3, 4], 6), 2);
        assert_eq!(Solution::coin_change(vec![4, 3, 1], 6), 2);
    }
}
