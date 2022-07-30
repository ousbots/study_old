pub struct Solution;

impl Solution {
    pub fn min_subarray_len(target: i32, nums: Vec<i32>) -> i32 {
        let mut size = 0;

        let mut pos_l = 0;
        let mut pos_r = 0;
        let mut count = if nums.len() > 0 { nums[0] } else { 0 };

        while pos_l < nums.len() && pos_r < nums.len() {
            if count >= target {
                let len = pos_r - pos_l + 1;
                if size == 0 || len < size {
                    size = len;
                }

                if pos_l < pos_r {
                    count -= nums[pos_l];
                    pos_l += 1;
                } else {
                    pos_r += 1;
                    if pos_r == nums.len() {
                        break;
                    }
                    count += nums[pos_r];
                }
            } else {
                pos_r += 1;
                if pos_r == nums.len() {
                    break;
                }
                count += nums[pos_r];
            }
        }

        size as i32
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn basic_test() {
        assert_eq!(Solution::min_subarray_len(0, vec![]), 0);
        assert_eq!(Solution::min_subarray_len(3, vec![1, 1, 1, 1]), 3);
        assert_eq!(Solution::min_subarray_len(3, vec![1, 1, 1, 2]), 2);
        assert_eq!(Solution::min_subarray_len(3, vec![1, 1, 2, 3]), 1);
        assert_eq!(Solution::min_subarray_len(7, vec![2, 3, 1, 2, 4, 3]), 2);
        assert_eq!(Solution::min_subarray_len(4, vec![1, 4, 4]), 1);
        assert_eq!(
            Solution::min_subarray_len(11, vec![1, 1, 1, 1, 1, 1, 1, 1]),
            0
        );
    }
}
