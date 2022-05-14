pub struct Solution;
impl Solution {
    pub fn combination_sum(candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        let mut result = vec![];
        Solution::backtrack(&mut result, &mut vec![], &candidates, 0, target);

        result
    }

    pub fn backtrack(
        result: &mut Vec<Vec<i32>>,
        sum: &mut Vec<i32>,
        vals: &Vec<i32>,
        index: usize,
        total: i32,
    ) {
        if total == 0 {
            result.push(sum.clone());
            return;
        }

        if total < 0 || index == vals.len() {
            return;
        }

        for i in index..vals.len() {
            sum.push(vals[i]);
            Solution::backtrack(result, sum, vals, i, total - vals[i]);
            sum.pop();
        }
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(
            Solution::combination_sum(vec![2, 3, 6, 7], 7).sort(),
            vec![vec![2, 2, 3], vec![7]].sort()
        );

        assert_eq!(
            Solution::combination_sum(vec![2, 3, 5], 8).sort(),
            vec![vec![2, 2, 2, 2], vec![2, 3, 3], vec![3, 5]].sort()
        );

        let empty: Vec<Vec<i32>> = vec![];
        assert_eq!(Solution::combination_sum(vec![2], 1), empty);

        let mut valid = vec![
            vec![1, 1, 1, 1, 1, 1, 1, 1, 1],
            vec![1, 1, 1, 1, 1, 1, 1, 2],
            vec![1, 1, 1, 1, 1, 1, 3],
            vec![1, 1, 1, 1, 1, 2, 2],
            vec![1, 1, 1, 1, 2, 3],
            vec![1, 1, 1, 1, 5],
            vec![1, 1, 1, 2, 2, 2],
            vec![1, 1, 1, 3, 3],
            vec![1, 1, 1, 6],
            vec![1, 1, 2, 2, 3],
            vec![1, 1, 2, 5],
            vec![1, 1, 7],
            vec![1, 2, 2, 2, 2],
            vec![1, 2, 3, 3],
            vec![1, 2, 6],
            vec![1, 3, 5],
            vec![2, 2, 2, 3],
            vec![2, 2, 5],
            vec![2, 7],
            vec![3, 3, 3],
            vec![3, 6],
        ];

        assert_eq!(
            Solution::combination_sum(vec![2, 7, 6, 3, 5, 1], 9).sort(),
            valid.sort()
        );
    }
}
