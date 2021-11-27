use std::collections::HashMap;

pub struct Solution;

impl Solution {
    pub fn permute_unique(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut collection: HashMap<Vec<i32>, bool> = HashMap::new();
        collection.insert(nums.clone(), true);
        Solution::permute(nums.len(), &mut nums.clone(), &mut collection);

        collection.keys().map(|k| k.to_vec()).collect()
    }

    // Heap's permutation algorithm.
    fn permute(k: usize, mut nums: &mut Vec<i32>, collection: &mut HashMap<Vec<i32>, bool>) {
        if k == 1 {
            collection.insert(nums.clone(), true);
            return;
        }

        Solution::permute(k - 1, nums, collection);

        for i in 0..k - 1 {
            if k % 2 == 0 {
                Solution::swap(i, k - 1, &mut nums);
            } else {
                Solution::swap(0, k - 1, &mut nums);
            }

            Solution::permute(k - 1, nums, collection)
        }
    }

    fn swap(first: usize, second: usize, nums: &mut Vec<i32>) {
        let temp = nums[first];
        nums[first] = nums[second];
        nums[second] = temp;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn basic_test() {
        assert_eq!(Solution::permute_unique(vec![1]), vec![vec![1]]);
        assert_eq!(Solution::permute_unique(vec![1, 1]), vec![vec![1, 1]]);

        let mut results = Solution::permute_unique(vec![1, 1, 2]);
        results.sort();

        let mut expected = vec![vec![1, 1, 2], vec![2, 1, 1], vec![1, 2, 1]];
        expected.sort();

        assert_eq!(results, expected);

        results = Solution::permute_unique(vec![1, 2, 3]);
        results.sort();

        expected = vec![
            vec![1, 2, 3],
            vec![2, 1, 3],
            vec![3, 1, 2],
            vec![1, 3, 2],
            vec![2, 3, 1],
            vec![3, 2, 1],
        ];
        expected.sort();

        assert_eq!(results, expected);

        results = Solution::permute_unique(vec![0, 1, 0, 0, 9]);
        results.sort();

        expected = vec![
            vec![0, 1, 0, 0, 9],
            vec![1, 0, 0, 0, 9],
            vec![0, 0, 1, 0, 9],
            vec![0, 0, 0, 1, 9],
            vec![9, 0, 0, 0, 1],
            vec![0, 9, 0, 0, 1],
            vec![0, 0, 9, 0, 1],
            vec![0, 0, 0, 9, 1],
            vec![1, 0, 0, 9, 0],
            vec![0, 1, 0, 9, 0],
            vec![0, 0, 1, 9, 0],
            vec![9, 0, 1, 0, 0],
            vec![0, 9, 1, 0, 0],
            vec![1, 9, 0, 0, 0],
            vec![9, 1, 0, 0, 0],
            vec![0, 1, 9, 0, 0],
            vec![1, 0, 9, 0, 0],
            vec![9, 0, 0, 1, 0],
            vec![0, 9, 0, 1, 0],
            vec![0, 0, 9, 1, 0],
        ];
        expected.sort();

        assert_eq!(results, expected);
    }
}
