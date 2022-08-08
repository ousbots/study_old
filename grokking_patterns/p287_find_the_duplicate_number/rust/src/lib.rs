pub struct Solution;
impl Solution {
    // Returns the duplicate number from the vector of integers. The integers range from 1..n and
    // the vector has a length of n+1. There is one integer that has at least one duplicate.
    //
    // Because each integer points to a valid position in the array, following the integers to
    // their position in the array (nums[i] -> nums[nums[i]]) will lead to a cycle with the
    // duplicated integers. I'm not exactly sure why, but the start of the cycle is always
    // one of the duplicated integers.
    pub fn find_duplicate(nums: Vec<i32>) -> i32 {
        let mut fast = 0;
        let mut slow = 0;

        loop {
            slow = nums[slow as usize];
            fast = nums[nums[fast as usize] as usize];

            if slow == fast {
                break;
            }
        }

        slow = 0;
        while slow != fast {
            slow = nums[slow as usize];
            fast = nums[fast as usize];
        }

        slow
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    struct TestCase {
        nums: Vec<i32>,
        expected: i32,
    }

    #[test]
    fn basic() {
        let tests = vec![
            TestCase {
                nums: vec![1, 3, 4, 2, 2],
                expected: 2,
            },
            TestCase {
                nums: vec![3, 1, 3, 4, 2],
                expected: 3,
            },
            TestCase {
                nums: vec![2, 5, 9, 6, 9, 3, 8, 9, 7, 1],
                expected: 9,
            },
        ];

        for test in tests {
            assert_eq!(
                Solution::find_duplicate(test.nums.clone()),
                test.expected,
                "case {:?}",
                test.nums
            );
        }
    }
}
