pub struct Solution;
impl Solution {
    pub fn merge(intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut intervals = intervals;
        intervals.sort_unstable_by(|x, y| x[0].cmp(&y[0]));

        let mut output = vec![];
        let mut curr = intervals[0].clone();

        for i in 1..intervals.len() {
            let next = intervals[i].clone();

            if curr[1] < next[0] {
                output.push(curr);
                curr = next;
            } else if curr[1] < next[1] {
                curr[1] = next[1]
            }
        }
        output.push(curr);

        output
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        let input = [[1, 3], [2, 6], [8, 10], [15, 18]].map(|x| x.to_vec());
        let valid = [[1, 6], [8, 10], [15, 18]].map(|x| x.to_vec());
        assert_eq!(Solution::merge(input.to_vec()), valid);

        let input = [[1, 4], [4, 5]].map(|x| x.to_vec());
        let valid = [[1, 5]].map(|x| x.to_vec());
        assert_eq!(Solution::merge(input.to_vec()), valid);
    }

    #[test]
    fn basic() {
        let input = [[1, 3], [5, 6], [8, 10], [4, 7]].map(|x| x.to_vec());
        let valid = [[1, 3], [4, 7], [8, 10]].map(|x| x.to_vec());
        assert_eq!(Solution::merge(input.to_vec()), valid);

        let input = [[1, 3], [5, 6], [8, 10], [4, 9]].map(|x| x.to_vec());
        let valid = [[1, 3], [4, 10]].map(|x| x.to_vec());
        assert_eq!(Solution::merge(input.to_vec()), valid);

        let input = [[1, 3], [5, 6], [8, 10], [0, 8]].map(|x| x.to_vec());
        let valid = [[0, 10]].map(|x| x.to_vec());
        assert_eq!(Solution::merge(input.to_vec()), valid);

        let input = [[1, 3], [5, 6], [8, 10], [0, 11]].map(|x| x.to_vec());
        let valid = [[0, 11]].map(|x| x.to_vec());
        assert_eq!(Solution::merge(input.to_vec()), valid);
    }
}
