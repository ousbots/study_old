pub struct Solution;

impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        let mut low: i32 = 0;
        let mut high: i32 = nums.len() as i32 - 1;

        while low <= high {
            let middle = (low + high) / 2;

            if nums[middle as usize] < target {
                low = middle + 1;
            } else if nums[middle as usize] > target {
                high = middle - 1;
            } else {
                return middle as i32;
            }
        }

        -1
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(Solution::search(vec![-1, 0, 3, 5, 9, 12], 9), 4);
        assert_eq!(Solution::search(vec![-1, 0, 3, 5, 9, 12], 2), -1);
    }

    #[test]
    fn basic() {
        assert_eq!(Solution::search(vec![-1], 1), -1);
        assert_eq!(Solution::search(vec![-1], -1), 0);
        assert_eq!(Solution::search(vec![-1, 1], -1), 0);
        assert_eq!(Solution::search(vec![-1, 1], 1), 1);
        assert_eq!(Solution::search(vec![-1, 1], 2), -1);
        assert_eq!(Solution::search(vec![5], -5), -1);
    }
}
