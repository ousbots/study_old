use std::collections::HashMap;

pub struct Solution;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut seen: HashMap<i32, usize> = HashMap::new();

        for index in 0..nums.len() {
            if seen.contains_key(&(target - nums[index])) {
                return vec![index as i32, seen[&(target - nums[index])] as i32];
            }

            seen.insert(nums[index], index);
        }

        vec![]
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        let mut ans = Solution::two_sum(vec![2, 7, 11, 15], 9);
        ans.sort();
        assert_eq!(ans, vec![0, 1]);

        let mut ans = Solution::two_sum(vec![3, 2, 4], 6);
        ans.sort();
        assert_eq!(ans, vec![1, 2]);

        let mut ans = Solution::two_sum(vec![3, 3], 6);
        ans.sort();
        assert_eq!(ans, vec![0, 1]);
    }

    #[test]
    fn basic() {
        let mut ans = Solution::two_sum(vec![0, 0], 0);
        ans.sort();
        assert_eq!(ans, vec![0, 1]);

        let mut ans = Solution::two_sum(vec![1, 5, 1, 5], 10);
        ans.sort();
        assert_eq!(ans, vec![1, 3]);

        let mut ans = Solution::two_sum(vec![1, 11, 1, 0, 5, -1], 10);
        ans.sort();
        assert_eq!(ans, vec![1, 5]);
    }
}
