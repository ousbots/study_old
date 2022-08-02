pub struct Solution;
impl Solution {
    // Returns the sum of the three elements of nums that is closest to the target.
    pub fn three_sum_closest(nums: Vec<i32>, target: i32) -> i32 {
        let mut nums = nums;
        nums.sort();

        // Since len(nums) <= 1000 and nums[i] <= 1000, then 10^6+1 is higher than any sum.
        let mut closest = 1_000_001;

        for fixed in 0..nums.len() {
            let mut lower = fixed + 1;
            let mut upper = nums.len() - 1;

            while lower < upper {
                let sum = nums[fixed] + nums[lower] + nums[upper];

                if sum < target {
                    lower += 1;
                } else if sum > target {
                    upper -= 1;
                } else {
                    return sum;
                }

                if (target - sum).abs() < (target - closest).abs() {
                    closest = sum;
                }
            }
        }

        closest
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    struct TestCase {
        nums: Vec<i32>,
        target: i32,
        output: i32,
    }

    #[test]
    fn basic() {
        let tests = vec![
            TestCase {
                nums: vec![0, 0, 0],
                target: 1,
                output: 0,
            },
            TestCase {
                nums: vec![0, 0, 1],
                target: 1,
                output: 1,
            },
            TestCase {
                nums: vec![1, 0, 0, 1],
                target: 0,
                output: 1,
            },
            TestCase {
                nums: vec![1, 0, 0, 1],
                target: 1,
                output: 1,
            },
            TestCase {
                nums: vec![1, 0, 0, 1],
                target: 2,
                output: 2,
            },
            TestCase {
                nums: vec![10, -1, -1, 5, 1, 1],
                target: 7,
                output: 7,
            },
            TestCase {
                nums: vec![-1, 2, 1, -4],
                target: 1,
                output: 2,
            },
        ];

        for test in tests {
            assert_eq!(
                Solution::three_sum_closest(test.nums.clone(), test.target),
                test.output,
                "testing nums: {:?} target: {}",
                test.nums,
                test.target
            )
        }
    }
}
