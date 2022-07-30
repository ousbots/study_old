pub struct Solution;
impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let mut majority = 0;
        let mut count = 0;

        for num in nums {
            if count == 0 {
                majority = num;
                count += 1;
            } else if num == majority {
                count += 1;
            } else {
                count -= 1;
            }
        }

        majority
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(Solution::majority_element(vec![3, 2, 3]), 3);
        assert_eq!(Solution::majority_element(vec![2, 2, 1, 1, 1, 2, 2]), 2);
    }

    #[test]
    fn basic() {
        assert_eq!(Solution::majority_element(vec![1]), 1);
    }
}
