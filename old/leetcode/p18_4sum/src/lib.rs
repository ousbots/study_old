use std::collections::HashSet;
use std::hash::Hash;

pub struct Solution;

// Generate all combinations of k elements.
pub fn k_combos<T: Clone + Copy + Hash + PartialEq + Ord>(elements: &[T], k: usize) -> Vec<Vec<T>> {
    let mut combos: HashSet<Vec<T>> = HashSet::new();
    find_combinations(elements, &mut combos, &mut vec![], 0, k);

    combos.into_iter().collect::<Vec<Vec<T>>>()
}

// Recursively generate all k-combinations with backtracing.
pub fn find_combinations<T: Clone + Copy + Hash + PartialEq + Ord>(
    nums: &[T],
    combos: &mut HashSet<Vec<T>>,
    current: &mut Vec<T>,
    begin: usize,
    k: usize,
) {
    if current.len() == k {
        combos.insert(current.clone());
        return;
    }

    for i in begin..nums.len() {
        current.push(nums[i]);
        find_combinations(nums, combos, current, i + 1, k);
        current.pop();
    }
}

impl Solution {
    // Generate all combinations of four elements form the given numbers that sum to the target.
    // Note: Due to performance, this has been changed to use an algorithm specific to
    // generating valid 4sums instead of generating all k-combos of length 4.
    pub fn four_sum(nums: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        if nums.len() < 4 {
            return vec![];
        }

        let mut nums = nums;
        nums.sort();

        let mut combos: Vec<Vec<i32>> = vec![];
        let len = nums.len();

        for i in 0..(len - 3) {
            if i > 0 && nums[i] == nums[i - 1] {
                continue;
            }

            for j in (i + 1)..(len - 2) {
                if j > i + 1 && nums[j] == nums[j - 1] {
                    continue;
                }

                let mut low = j + 1;
                let mut high = len - 1;

                while low < high {
                    if low > j + 1 && nums[low] == nums[low - 1] {
                        low += 1;
                        continue;
                    }

                    if nums[i] + nums[j] + nums[low] + nums[high] > target {
                        high -= 1;
                    } else if nums[i] + nums[j] + nums[low] + nums[high] < target {
                        low += 1;
                    } else {
                        combos.push(vec![nums[i], nums[j], nums[low], nums[high]]);
                        low += 1;
                        high -= 1;
                    }
                }
            }
        }

        combos
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn provided_tests() {
        let mut result = Solution::four_sum(vec![1, 0, -1, 0, -2, 2], 0);
        let mut answer: Vec<Vec<i32>> =
            vec![vec![-2, -1, 1, 2], vec![-2, 0, 0, 2], vec![-1, 0, 0, 1]]
                .into_iter()
                .map(|mut x| {
                    x.sort();
                    x
                })
                .collect();

        result.sort();
        answer.sort();
        assert_eq!(result, answer);

        let mut result = Solution::four_sum(vec![2, 2, 2, 2, 2], 8);
        let mut answer: Vec<Vec<i32>> = vec![vec![2, 2, 2, 2]]
            .into_iter()
            .map(|mut x| {
                x.sort();
                x
            })
            .collect();

        result.sort();
        answer.sort();
        assert_eq!(result, answer);
    }

    #[test]
    fn combination_tests() {
        let mut data = k_combos(&vec![1], 1);
        data.sort();
        assert_eq!(data, [[1]]);

        let mut data = k_combos(&vec![1, 2], 1);
        data.sort();
        assert_eq!(data, [[1], [2]]);

        let mut data = k_combos(&vec![1, 2, 3], 2);
        data.sort();
        assert_eq!(data, [[1, 2], [1, 3], [2, 3]]);
    }

    #[test]
    fn more_tests() {
        assert_eq!(Solution::four_sum(vec![1], 2), Vec::<Vec<i32>>::new());

        let mut result = Solution::four_sum(vec![-2, -1, -1, 1, 1, 2, 2], 0);
        let mut answer: Vec<Vec<i32>> = vec![vec![-2, -1, 1, 2], vec![-1, -1, 1, 1]]
            .into_iter()
            .map(|mut x| {
                x.sort();
                x
            })
            .collect();

        result.sort();
        answer.sort();
        assert_eq!(result, answer);

        let mut result = Solution::four_sum(vec![-5, 5, 4, -3, 0, 0, 4, -2], 4);
        let mut answer: Vec<Vec<i32>> = vec![vec![-5, 0, 4, 5], vec![-3, -2, 4, 5]]
            .into_iter()
            .map(|mut x| {
                x.sort();
                x
            })
            .collect();

        result.sort();
        answer.sort();
        assert_eq!(result, answer);

        let result = Solution::four_sum(
            vec![
                -500, -481, -480, -469, -437, -423, -408, -403, -397, -381, -379, -377, -353, -347,
                -337, -327, -313, -307, -299, -278, -265, -258, -235, -227, -225, -193, -192, -177,
                -176, -173, -170, -164, -162, -157, -147, -118, -115, -83, -64, -46, -36, -35, -11,
                0, 0, 33, 40, 51, 54, 74, 93, 101, 104, 105, 112, 112, 116, 129, 133, 146, 152,
                157, 158, 166, 177, 183, 186, 220, 263, 273, 320, 328, 332, 356, 357, 363, 372,
                397, 399, 420, 422, 429, 433, 451, 464, 484, 485, 498, 499,
            ],
            2139,
        );
        assert_eq!(result.len() == 0, true);
    }
}
