pub struct Solution;
impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        let mut first = 0;
        let mut second = height.len() - 1;
        let mut max = 0;

        while second > first {
            max = max.max(height[first].min(height[second]) * (second - first) as i32);

            if height[second] > height[first] {
                first += 1;
            } else {
                second -= 1;
            }
        }

        max
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn basic() {
        assert_eq!(Solution::max_area(vec![1, 1]), 1);
        assert_eq!(Solution::max_area(vec![1, 2, 3]), 2);
        assert_eq!(Solution::max_area(vec![3, 2, 1]), 2);
        assert_eq!(Solution::max_area(vec![1, 1, 1, 1, 9, 9, 1, 1]), 9);
        assert_eq!(Solution::max_area(vec![1, 1, 1, 1, 6, 6, 1, 1]), 7);
        assert_eq!(Solution::max_area(vec![1, 1, 1, 1, 6, 1, 1, 1]), 7);
    }
}
