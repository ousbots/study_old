pub struct Solution;

impl Solution {
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        let mut max = nums[0];
        let mut total = 0;

        for num in nums {
            total += num;

            if total > 0 {
                if total > max {
                    max = total;
                }
            } else {
                total = 0;
                if num > max {
                    max = num;
                }
            }
        }

        max
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(
            Solution::max_sub_array(vec![-2, 1, -3, 4, -1, 2, 1, -5, 4]),
            6
        );

        assert_eq!(Solution::max_sub_array(vec![5, 4, -1, 7, 8]), 23);

        assert_eq!(Solution::max_sub_array(vec![1]), 1);
    }

    #[test]
    fn basic() {
        assert_eq!(Solution::max_sub_array(vec![-3, -2, -1]), -1);
        assert_eq!(Solution::max_sub_array(vec![-2, -1, -3]), -1);
        assert_eq!(Solution::max_sub_array(vec![-1, -3, -2]), -1);
        assert_eq!(Solution::max_sub_array(vec![-1]), -1);
        assert_eq!(Solution::max_sub_array(vec![1, -2, 5]), 5);
        assert_eq!(Solution::max_sub_array(vec![5, -2, 1]), 5);
        assert_eq!(Solution::max_sub_array(vec![-2, 5, -1]), 5);
        assert_eq!(Solution::max_sub_array(vec![-2, 5, 1]), 6);
    }

    #[test]
    fn extra() {
        assert_eq!(
            Solution::max_sub_array(vec![
                -6, -1, 5, 4, 1, -8, 6, 7, -3, 6, 0, -6, -7, 8, -8, -4, 1
            ]),
            18
        );
    }
}
