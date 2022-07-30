pub struct Solution;

impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        if let Some(x) = find(&nums, target) {
            return x as i32;
        } else {
            return -1;
        }
    }
}

// The unique property of the rotated sorted array is that one of the halves is sorted and
// the other is not. The binary search is modified to check which half is sorted and switch
// to the other half if the target isn't in the sorted half.
fn find(nums: &[i32], target: i32) -> Option<usize> {
    if nums.len() == 0 {
        return None;
    }

    let mut start: usize = 0;
    let mut end: usize = nums.len() - 1;

    while start <= end {
        let half: usize = (end + start) / 2;

        if nums[half] == target {
            return Some(half);
        }

        if nums[start] <= nums[half] {
            if nums[start] > target || nums[half] < target {
                start = half + 1;
                continue;
            }

            if half == 0 {
                break;
            }

            end = half - 1;
            continue;
        }

        if nums[half] <= nums[end] {
            if nums[half] > target || nums[end] < target {
                if half == 0 {
                    break;
                }

                end = half - 1;
                continue;
            }

            start = half + 1;
            continue;
        }
    }

    None
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn basic_tests() {
        assert_eq!(Solution::search(vec![4, 5, 6, 7, 0, 1, 2], 0), 4);
        assert_eq!(Solution::search(vec![0, 1, 2, 4, 5, 6, 7], 0), 0);
        assert_eq!(Solution::search(vec![0, 1, 2, 4, 5, 6, 7], 4), 3);
        assert_eq!(Solution::search(vec![1, 2, 4, 5, 6, 7, 0], 0), 6);
        assert_eq!(Solution::search(vec![1, 2, 4, 5, 6, 7, 0], 1), 0);
        assert_eq!(Solution::search(vec![7, 0, 1, 2, 4, 5, 6], 6), 6);
        assert_eq!(Solution::search(vec![7, 0, 1, 2, 4, 5, 6], 7), 0);
        assert_eq!(Solution::search(vec![7, 0, 1, 2, 4, 5, 6], 0), 1);
        assert_eq!(Solution::search(vec![4, 5, 6, 7, 0, 1, 2], 3), -1);
        assert_eq!(Solution::search(vec![1], 0), -1);
        assert_eq!(Solution::search(vec![], 0), -1);
    }
}
