pub struct Solution;
impl Solution {
    pub fn searchIntuition(nums: Vec<i32>, target: i32) -> i32 {
        let mut lo = 0;
        let mut hi = nums.len() - 1;

        while lo <= hi {
            let middle = (lo + hi) / 2;

            if target < nums[middle] {
                if nums[hi] < nums[middle] && target < nums[lo] {
                    lo = middle + 1;
                } else {
                    if middle > 0 {
                        hi = middle - 1;
                    } else {
                        break;
                    }
                }
            } else if target > nums[middle] {
                if nums[lo] > nums[middle] && target > nums[hi] {
                    if middle > 0 {
                        hi = middle - 1;
                    } else {
                        break;
                    }
                } else {
                    lo = middle + 1;
                }
            } else {
                return middle as i32;
            }
        }

        -1
    }

    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        let mut lo = 0;
        let mut hi = nums.len() - 1;

        while lo <= hi {
            let middle = (lo + hi) / 2;
            println!("{:?} {} {} {} ({})", nums, lo, middle, hi, target);

            if nums[middle] == target {
                return middle as i32;
            }

            // Rotation is in the lower half.
            if nums[lo] <= nums[middle] {
                if nums[middle] < target || target < nums[lo] {
                    lo = middle + 1;
                } else {
                    if middle != 0 {
                        hi = middle - 1;
                    } else {
                        break;
                    }
                }
            // Rotation is in the upper half.
            } else {
                if nums[middle] > target || target > nums[hi] {
                    if middle != 0 {
                        hi = middle - 1;
                    } else {
                        break;
                    }
                } else {
                    lo = middle + 1;
                }
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
        assert_eq!(Solution::search(vec![4, 5, 6, 7, 0, 1, 2], 0), 4);
        assert_eq!(Solution::search(vec![4, 5, 6, 7, 0, 1, 2], 3), -1);
        assert_eq!(Solution::search(vec![1], 0), -1);
    }

    #[test]
    fn basic() {
        assert_eq!(Solution::search(vec![0, 1, 2, 3, 4, 5, 6], 0), 0);
        assert_eq!(Solution::search(vec![0, 1, 2, 3, 4, 5, 6], 6), 6);
        assert_eq!(Solution::search(vec![0, 1, 2, 3, 4, 5, 6], 3), 3);
        assert_eq!(Solution::search(vec![0, 1, 2, 3, 4, 5, 6], 8), -1);
        assert_eq!(Solution::search(vec![0, 1, 2, 3, 4, 5, 6], -1), -1);
        assert_eq!(Solution::search(vec![1, 3, 5, 7, 9, -1], 1), 0);
        assert_eq!(Solution::search(vec![1, 3, 5, 7, 9, -1], -1), 5);
        assert_eq!(Solution::search(vec![1, 3, 5, 7, 9, -1], 9), 4);
        assert_eq!(Solution::search(vec![9, -1, 1, 3, 5, 7], 9), 0);
        assert_eq!(Solution::search(vec![9, -1, 1, 3, 5, 7], 7), 5);
        assert_eq!(Solution::search(vec![9, -1, 1, 3, 5, 7], -1), 1);
        assert_eq!(Solution::search(vec![9, -1, 1, 3, 5, 7], 5), 4);
        assert_eq!(Solution::search(vec![9, -1, 1, 3, 5, 7], 11), -1);
        assert_eq!(Solution::search(vec![9, -1, 1, 3, 5, 7], -2), -1);
        assert_eq!(Solution::search(vec![9, -1, 1, 3, 5, 7], 2), -1);
        assert_eq!(Solution::search(vec![0], 0), 0);
        assert_eq!(Solution::search(vec![0], 1), -1);
    }
}
