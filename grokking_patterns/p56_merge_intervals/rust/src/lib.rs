pub struct Solution;
impl Solution {
    pub fn merge(intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut intervals = intervals;
        intervals.sort();
        let mut merged: Vec<Vec<i32>> = vec![intervals[0].clone()];
        let mut m_index = 0;

        for i in 1..intervals.len() {
            if !Self::overlap(&merged[m_index], &intervals[i]) {
                merged.push(intervals[i].clone());
                m_index += 1;
            } else {
                let start = merged[m_index][0].min(intervals[i][0]);
                let end = merged[m_index][1].max(intervals[i][1]);

                (merged[m_index][0], merged[m_index][1]) = (start, end);
            }
        }

        merged
    }

    fn overlap(a: &Vec<i32>, b: &Vec<i32>) -> bool {
        let start = a[0].max(b[0]);
        let end = a[1].min(b[1]);

        end - start >= 0
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    struct TestCase {
        intervals: Vec<Vec<i32>>,
        expected: Vec<Vec<i32>>,
    }

    #[test]
    fn basic() {
        let tests = vec![
            TestCase {
                intervals: vec![vec![1, 2]],
                expected: vec![vec![1, 2]],
            },
            TestCase {
                intervals: vec![vec![1, 2], vec![3, 4]],
                expected: vec![vec![1, 2], vec![3, 4]],
            },
            TestCase {
                intervals: vec![vec![1, 2], vec![2, 4]],
                expected: vec![vec![1, 4]],
            },
            TestCase {
                intervals: vec![vec![1, 2], vec![0, 0]],
                expected: vec![vec![0, 0], vec![1, 2]],
            },
            TestCase {
                intervals: vec![vec![1, 2], vec![2, 2]],
                expected: vec![vec![1, 2]],
            },
        ];

        for test in tests {
            assert_eq!(
                Solution::merge(test.intervals.clone()),
                test.expected,
                "case {:?}",
                test.intervals,
            )
        }
    }
}
