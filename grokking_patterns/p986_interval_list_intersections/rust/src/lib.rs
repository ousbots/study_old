pub struct Solution;
impl Solution {
    // Determines the union (overlapping sections) of the two given vectors of intervals. Each
    // set of intervals is disjoint and ordered.
    //
    // This solution uses two "pointers" to track which interval is the current one for each list.
    // Each current interval is checked for overlap and any overlap is recorded. Whichever
    // interval's end is less than the other is incremented.
    pub fn interval_intersection(
        first_list: Vec<Vec<i32>>,
        second_list: Vec<Vec<i32>>,
    ) -> Vec<Vec<i32>> {
        let mut overlaps: Vec<Vec<i32>> = Vec::new();
        let (mut first, mut second, f_end, s_end) = (0, 0, first_list.len(), second_list.len());

        while first < f_end && second < s_end {
            let front = first_list[first][0].max(second_list[second][0]);
            let back = first_list[first][1].min(second_list[second][1]);

            if front <= back {
                overlaps.push(vec![front, back]);
            }

            if first_list[first][1] < second_list[second][1] {
                first += 1;
            } else {
                second += 1;
            }
        }

        overlaps
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    struct TestCase {
        first_list: Vec<Vec<i32>>,
        second_list: Vec<Vec<i32>>,
        expected: Vec<Vec<i32>>,
    }

    #[test]
    fn basic() {
        let tests = vec![
            TestCase {
                first_list: vec![vec![1, 5], vec![7, 9], vec![10, 13]],
                second_list: vec![vec![2, 3], vec![5, 5], vec![6, 11], vec![13, 15]],
                expected: vec![
                    vec![2, 3],
                    vec![5, 5],
                    vec![7, 9],
                    vec![10, 11],
                    vec![13, 13],
                ],
            },
            TestCase {
                first_list: vec![vec![0, 2], vec![5, 10], vec![13, 23], vec![24, 25]],
                second_list: vec![vec![1, 5], vec![8, 12], vec![15, 24], vec![25, 26]],
                expected: vec![
                    vec![1, 2],
                    vec![5, 5],
                    vec![8, 10],
                    vec![15, 23],
                    vec![24, 24],
                    vec![25, 25],
                ],
            },
            TestCase {
                first_list: vec![vec![1, 3], vec![5, 9]],
                second_list: vec![],
                expected: vec![],
            },
            TestCase {
                first_list: vec![vec![1, 3], vec![5, 9]],
                second_list: vec![vec![0, 0], vec![4, 4], vec![10, 20]],
                expected: vec![],
            },
        ];

        for test in tests {
            assert_eq!(
                Solution::interval_intersection(test.first_list.clone(), test.second_list.clone()),
                test.expected,
                "case first_list={:?} second_list={:?}",
                test.first_list,
                test.second_list
            );
        }
    }
}
