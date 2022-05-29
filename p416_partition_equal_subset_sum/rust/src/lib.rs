pub struct Solution;
impl Solution {
    pub fn can_partition(nums: Vec<i32>) -> bool {
        let sum = nums.iter().sum::<i32>();
        let sum = sum as usize;

        if sum % 2 == 1 {
            return false;
        }

        let mut dp = vec![false; sum + 1];
        dp[0] = true;

        for num in nums {
            for j in (0..=sum).rev() {
                if dp[j] {
                    dp[j + num as usize] = true
                }
            }
        }

        dp[sum / 2]
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn basic() {
        assert_eq!(Solution::can_partition(vec![1]), false);
        assert_eq!(Solution::can_partition(vec![1, 1]), true);
        assert_eq!(Solution::can_partition(vec![1, 3, 1, 1]), true);
        assert_eq!(Solution::can_partition(vec![3, 1, 1, 1, 1]), false);
    }
}
