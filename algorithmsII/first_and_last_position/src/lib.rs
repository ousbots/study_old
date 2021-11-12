use std::convert::TryFrom;

pub struct Solution;

impl Solution {
    pub fn search_range(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let found_pos = find(&nums, target);

        if found_pos.is_none() {
            return vec![-1, -1];
        }

        let mut first = found_pos.unwrap();
        let mut last = found_pos.unwrap();

        while first >= 1 {
            if nums[first - 1] == target {
                first -= 1;
            } else {
                break;
            }
        }

        while last + 1 < nums.len() {
            if nums[last + 1] == target {
                last += 1;
            } else {
                break;
            }
        }

        return vec![i32::try_from(first).unwrap(), i32::try_from(last).unwrap()];
    }
}

fn find(nums: &[i32], target: i32) -> Option<usize> {
    if nums.len() == 0 {
        return None;
    }

    let mut start: usize = 0;
    let mut end: usize = nums.len() - 1;

    while start <= end {
        let half: usize = (end + start) / 2;

        if nums[half] > target {
            if half == 0 {
                break;
            }

            end = half - 1;
            continue;
        }

        if nums[half] < target {
            start = half + 1;
            continue;
        }

        if nums[half] == target {
            return Some(half);
        }
    }

    None
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn valid_tests() {
        assert_eq!(
            Solution::search_range(vec![5, 7, 7, 8, 8, 10], 8),
            vec![3, 4]
        );
        assert_eq!(
            Solution::search_range(vec![1, 1, 1, 1, 1, 1], 1),
            vec![0, 5]
        );
        assert_eq!(
            Solution::search_range(vec![-1, -1, -1, -1, -1, -1], -1),
            vec![0, 5]
        );
        assert_eq!(Solution::search_range(vec![-1], -1), vec![0, 0]);
        assert_eq!(Solution::search_range(vec![1], 1), vec![0, 0]);
    }

    #[test]
    fn invalid_tests() {
        assert_eq!(
            Solution::search_range(vec![5, 7, 7, 8, 8, 10], 6),
            vec![-1, -1]
        );
        assert_eq!(
            Solution::search_range(vec![5, 7, 7, 8, 8, 10], -1),
            vec![-1, -1]
        );
        assert_eq!(Solution::search_range(vec![-3], -1), vec![-1, -1]);
        assert_eq!(Solution::search_range(vec![1], 0), vec![-1, -1]);
    }

    #[test]
    fn empty_tests() {
        assert_eq!(Solution::search_range(vec![], 9), vec![-1, -1]);
        assert_eq!(Solution::search_range(vec![], -9), vec![-1, -1]);
    }
}
