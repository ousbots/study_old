pub struct Solution;
impl Solution {
    // Returns the merged intervals after inserting the new interval. The sets in intervals are all
    // disjoint and sorted by intervals[i][1].
    //
    // This solution first inserts the new_interval into the intervals vector and then sorts the
    // vector by the first elements of each interval. Each interval is then checked against its
    // neighbor intervals for any overlap and if there is any, the intervals are merged. This is
    // done by comparing each interval to the previous one and merging if needed.
    pub fn insert(intervals: Vec<Vec<i32>>, new_interval: Vec<i32>) -> Vec<Vec<i32>> {
        let mut intervals = intervals;
        intervals.push(new_interval);
        intervals.sort();

        let mut merged = vec![intervals[0].clone()];
        let mut m_index = 0;

        for interval in intervals {
            if !Self::overlap(&merged[m_index], &interval) {
                merged.push(interval.clone());
                m_index += 1;
            } else {
                let start = merged[m_index][0].min(interval[0]);
                let end = merged[m_index][1].max(interval[1]);
                (merged[m_index][0], merged[m_index][1]) = (start, end)
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
        new_interval: Vec<i32>,
        expected: Vec<Vec<i32>>,
    }

    #[test]
    fn basic() {
        let tests = vec![
            TestCase {
                intervals: vec![vec![0, 1]],
                new_interval: vec![2, 3],
                expected: vec![vec![0, 1], vec![2, 3]],
            },
            TestCase {
                intervals: vec![vec![0, 1]],
                new_interval: vec![1, 3],
                expected: vec![vec![0, 3]],
            },
            TestCase {
                intervals: vec![vec![0, 1], vec![3, 4], vec![6, 7], vec![9, 10]],
                new_interval: vec![2, 8],
                expected: vec![vec![0, 1], vec![2, 8], vec![9, 10]],
            },
            TestCase {
                intervals: vec![vec![1, 3], vec![6, 9]],
                new_interval: vec![2, 5],
                expected: vec![vec![1, 5], vec![6, 9]],
            },
            TestCase {
                intervals: vec![vec![1, 3]],
                new_interval: vec![0, 0],
                expected: vec![vec![0, 0], vec![1, 3]],
            },
        ];

        for test in tests {
            assert_eq!(
                Solution::insert(test.intervals.clone(), test.new_interval.clone()),
                test.expected,
                "case interval={:?} new_interval={:?}",
                test.intervals,
                test.new_interval
            );
        }
    }
}
