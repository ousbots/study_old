pub struct Solution;
impl Solution {
    pub fn subsets(nums: Vec<i32>) -> Vec<Vec<i32>> {
        Solution::generate(&nums, vec![], 0)
    }

    fn generate(nums: &Vec<i32>, set: Vec<i32>, index: usize) -> Vec<Vec<i32>> {
        if index >= nums.len() {
            return vec![set];
        }

        let mut b = set.clone();
        b.push(nums[index]);

        let mut next = Solution::generate(nums, b, index + 1);
        next.append(&mut Solution::generate(nums, set, index + 1));

        next
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn basic() {
        let mut result = Solution::subsets(vec![1]);
        let mut valid = vec![vec![], vec![1]];
        result.sort_unstable();
        valid.sort_unstable();
        assert_eq!(result, valid);

        let mut result = Solution::subsets(vec![1, 2]);
        let mut valid = vec![vec![], vec![1], vec![2], vec![1, 2]];
        result.sort_unstable();
        valid.sort_unstable();
        assert_eq!(result, valid);

        let mut result = Solution::subsets(vec![1, 2, 3, 4, 5]);
        let mut valid = vec![
            vec![],
            vec![1],
            vec![2],
            vec![3],
            vec![4],
            vec![5],
            vec![1, 2],
            vec![1, 3],
            vec![1, 4],
            vec![1, 5],
            vec![2, 3],
            vec![2, 4],
            vec![2, 5],
            vec![3, 4],
            vec![3, 5],
            vec![4, 5],
            vec![1, 2, 3],
            vec![1, 2, 4],
            vec![1, 2, 5],
            vec![1, 3, 4],
            vec![1, 3, 5],
            vec![1, 4, 5],
            vec![2, 3, 4],
            vec![2, 3, 5],
            vec![2, 4, 5],
            vec![3, 4, 5],
            vec![1, 2, 3, 4],
            vec![1, 2, 3, 5],
            vec![1, 2, 4, 5],
            vec![1, 3, 4, 5],
            vec![2, 3, 4, 5],
            vec![1, 2, 3, 4, 5],
        ];
        result.sort_unstable();
        valid.sort_unstable();
        assert_eq!(result, valid);
    }
}
