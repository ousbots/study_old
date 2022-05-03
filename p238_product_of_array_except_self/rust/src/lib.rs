pub struct Solution;

impl Solution {
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        let mut total: i32 = 1;
        let mut zeros: i32 = 0;
        let mut ret = nums.clone();

        for num in nums {
            if num == 0 {
                zeros += 1;
            } else {
                total *= num;
            }
        }

        for i in 0..ret.len() {
            if zeros > 1 {
                ret[i] = 0;
            } else if zeros == 1 {
                if ret[i] != 0 {
                    ret[i] = 0;
                } else {
                    ret[i] = total;
                }
            } else {
                ret[i] = total / ret[i];
            }
        }

        ret
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(
            Solution::product_except_self(vec![1, 2, 3, 4]),
            vec![24, 12, 8, 6]
        );

        assert_eq!(
            Solution::product_except_self(vec![-1, 1, 0, -3, 3]),
            vec![0, 0, 9, 0, 0]
        );
    }

    #[test]
    fn basic() {
        assert_eq!(Solution::product_except_self(vec![0, 0]), vec![0, 0]);
        assert_eq!(Solution::product_except_self(vec![1, 1]), vec![1, 1]);
        assert_eq!(Solution::product_except_self(vec![1, 10]), vec![10, 1]);
        assert_eq!(
            Solution::product_except_self(vec![1, 0, 2, 3]),
            vec![0, 6, 0, 0]
        );
        assert_eq!(
            Solution::product_except_self(vec![1, 0, 2, 0, 3]),
            vec![0, 0, 0, 0, 0]
        );
    }
}
