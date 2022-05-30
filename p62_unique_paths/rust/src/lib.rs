pub struct Solution;
impl Solution {
    pub fn unique_paths_fib(m: i32, n: i32) -> i32 {
        if m == 1 || n == 1 {
            return 1;
        }

        if m == 2 {
            return n;
        }

        if n == 2 {
            return m;
        }

        Solution::unique_paths(m - 1, n) + Solution::unique_paths(m, n - 1)
    }

    // Dynamic programming solution builds number of moves O(m + n + n*m) -> O(m*n)
    pub fn unique_paths(m: i32, n: i32) -> i32 {
        let mut count = Vec::<Vec<i32>>::with_capacity(m as usize);

        for i in 0..m as usize {
            count.push(vec![0; n as usize]);
            count[i][0] = 1;
        }

        for i in 0..count[0].len() {
            count[0][i] = 1;
        }

        for i in 1..count.len() {
            for j in 1..count[0].len() {
                count[i][j] = count[i - 1][j] + count[i][j - 1];
            }
        }

        count[m as usize - 1][n as usize - 1]
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;
    #[test]
    fn basic() {
        assert_eq!(Solution::unique_paths(1, 1), 1);
        assert_eq!(Solution::unique_paths(1, 2), 1);
        assert_eq!(Solution::unique_paths(2, 1), 1);
        assert_eq!(Solution::unique_paths(2, 3), 3);
        assert_eq!(Solution::unique_paths(3, 2), 3);
        assert_eq!(Solution::unique_paths(4, 4), 20);
        assert_eq!(Solution::unique_paths(3, 7), 28);
    }
}
