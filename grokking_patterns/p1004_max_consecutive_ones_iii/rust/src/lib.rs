use std::collections::VecDeque;

pub struct Solution;
impl Solution {
    // Returns the maximum number of consecutive 1's in the nums vector if at most k 0's can be
    // flipeed. The elements of the nums vector are either 1 or 0.
    pub fn longest_ones(nums: Vec<i32>, k: i32) -> i32 {
        let mut max = 0;
        let mut back = 0;
        let mut flipped = 0;
        let mut positions: VecDeque<usize> = VecDeque::new();

        for front in 0..nums.len() {
            if nums[front] == 0 {
                flipped += 1;
                positions.push_back(front);

                while flipped > k && back <= front {
                    if back == positions[0] {
                        positions.pop_front();
                        flipped -= 1;
                    }

                    back += 1;
                }
            }

            if front >= back {
                let length = i32::try_from(front - back + 1).unwrap();
                if length > max {
                    max = length;
                }
            }
        }

        max
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    struct TestCase {
        nums: Vec<i32>,
        k: i32,
        output: i32,
    }

    #[test]
    fn basic() {
        let tests = vec![
            TestCase {
                nums: vec![1],
                k: 1,
                output: 1,
            },
            TestCase {
                nums: vec![0],
                k: 0,
                output: 0,
            },
            TestCase {
                nums: vec![0],
                k: 1,
                output: 1,
            },
            TestCase {
                nums: vec![0, 1, 0],
                k: 1,
                output: 2,
            },
            TestCase {
                nums: vec![
                    0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0,
                ],
                k: 2,
                output: 14,
            },
            TestCase {
                nums: vec![1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0],
                k: 2,
                output: 6,
            },
            TestCase {
                nums: vec![0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1],
                k: 3,
                output: 10,
            },
        ];

        for test in tests {
            assert_eq!(
                Solution::longest_ones(test.nums.clone(), test.k),
                test.output,
                "testing nums = {:?} k = {}",
                test.nums,
                test.k
            );
        }
    }
}
