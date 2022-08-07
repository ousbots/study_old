pub struct Solution;
impl Solution {
    // Returns an array of the number of seats booked for each flight from the given bookings. Each
    // element in the bookings vector is a vector where the first element is the starting flight,
    // the second element is the last flight, and the third element is the number of seats to be
    // booked. The flight numbers range from 1 to n.
    //
    // This solution uses the "difference algorithm" where the value of each interval is added to
    // the start position and subtracted from the position after the end position of a new vector.
    // Then the values vector is iterated over and the value of the previous position added to the
    // current one (i.e. vals[i] += vals[i-1]). This way, the value of each iterval is propagated
    // across the appropriate ranges.
    pub fn corp_flight_bookings(bookings: Vec<Vec<i32>>, n: i32) -> Vec<i32> {
        let mut counts = vec![0; usize::try_from(n + 1).unwrap()];

        for interval in bookings {
            counts[usize::try_from(interval[0] - 1).unwrap()] += interval[2];
            counts[usize::try_from(interval[1]).unwrap()] -= interval[2];
        }

        for i in 1..counts.len() {
            counts[i] += counts[i - 1];
        }

        counts[0..counts.len() - 1].to_vec()
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    struct TestCase {
        bookings: Vec<Vec<i32>>,
        n: i32,
        expected: Vec<i32>,
    }

    #[test]
    fn basic() {
        let tests = vec![
            TestCase {
                bookings: vec![vec![1, 2, 1]],
                n: 3,
                expected: vec![1, 1, 0],
            },
            TestCase {
                bookings: vec![
                    vec![2, 3, 10],
                    vec![3, 5, 5],
                    vec![7, 10, 5],
                    vec![9, 11, 1],
                ],
                n: 15,
                expected: vec![0, 10, 15, 5, 5, 0, 5, 5, 6, 6, 1, 0, 0, 0, 0],
            },
            TestCase {
                bookings: vec![vec![1, 2, 10], vec![2, 3, 20], vec![2, 5, 25]],
                n: 5,
                expected: vec![10, 55, 45, 25, 25],
            },
            TestCase {
                bookings: vec![vec![1, 2, 10], vec![2, 2, 15]],
                n: 2,
                expected: vec![10, 25],
            },
        ];

        for test in tests {
            assert_eq!(
                Solution::corp_flight_bookings(test.bookings.clone(), test.n),
                test.expected,
                "case bookings={:?} n={}",
                test.bookings,
                test.n
            );
        }
    }
}
