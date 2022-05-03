pub struct Solution;
impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {
        let n = n as usize;
        let mut values: Vec<i32> = vec![1; n + 1];

        for i in 2..=n {
            values[i] = values[i - 1] + values[i - 2];
        }

        values[n]
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn basic() {
        assert_eq!(Solution::climb_stairs(2), 2);
        assert_eq!(Solution::climb_stairs(3), 3);
    }
}
