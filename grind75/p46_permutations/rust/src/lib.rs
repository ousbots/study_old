pub struct Solution;
impl Solution {
    pub fn permute(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut nums = nums;
        let mut perms = vec![];
        Solution::backtrack(&mut nums, &mut perms, 0);

        perms
    }

    fn backtrack(nums: &mut Vec<i32>, perms: &mut Vec<Vec<i32>>, start: usize) {
        if start == nums.len() {
            perms.push(nums.clone());
        }

        for i in start..nums.len() {
            (nums[start], nums[i]) = (nums[i], nums[start]);
            Solution::backtrack(nums, perms, start + 1);
            (nums[start], nums[i]) = (nums[i], nums[start]);
        }
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        let mut valid = vec![
            vec![1, 2, 3],
            vec![1, 3, 2],
            vec![2, 1, 3],
            vec![2, 3, 1],
            vec![3, 1, 2],
            vec![3, 2, 1],
        ];
        assert_eq!(Solution::permute(vec![1, 2, 3]).sort(), valid.sort());

        let mut valid = vec![vec![0, 1], vec![1, 0]];
        assert_eq!(Solution::permute(vec![0, 1]).sort(), valid.sort());

        let mut valid = vec![vec![1]];
        assert_eq!(Solution::permute(vec![1]).sort(), valid.sort());
    }
}
