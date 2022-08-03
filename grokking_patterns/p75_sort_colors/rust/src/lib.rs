pub struct Solution;
impl Solution {
    // Sorts the nums vector in place. Only the numbers 0, 1, and 2 are in the vector.
    //
    // This is a two pointers solution where there is a front and back pointer, the array is then
    // iterated through and if the current position is 0 then it is swapped with the back pointer
    // and the back pointer incremented, if the position is 2 then it is swapped with the front
    // pointer and the front point decremented. After a swap, the index position should not be
    // incremented so that the value that was swapped in can be checked. Because of the limits of
    // usize, a skip boolean is used instead of decrementing the position.
    pub fn sort_colors(nums: &mut Vec<i32>) {
        let mut back = 0;
        let mut front = nums.len() - 1;

        let mut skip = false;
        let mut i = 0;
        while i <= front {
            if nums[i] == 0 && i != back {
                (nums[back], nums[i]) = (nums[i], nums[back]);
                back += 1;
                skip = true;
            } else if nums[i] == 2 && i != front {
                (nums[front], nums[i]) = (nums[i], nums[front]);
                front -= 1;
                skip = true;
            }

            if skip {
                skip = false;
            } else {
                i += 1;
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    struct TestCase {
        nums: Vec<i32>,
        expected: Vec<i32>,
    }

    #[test]
    fn basic() {
        let tests = vec![
            TestCase {
                nums: vec![1],
                expected: vec![1],
            },
            TestCase {
                nums: vec![2, 1, 0],
                expected: vec![0, 1, 2],
            },
            TestCase {
                nums: vec![2, 0, 1],
                expected: vec![0, 1, 2],
            },
            TestCase {
                nums: vec![1, 0, 2],
                expected: vec![0, 1, 2],
            },
            TestCase {
                nums: vec![1, 2, 0],
                expected: vec![0, 1, 2],
            },
            TestCase {
                nums: vec![0, 2, 1],
                expected: vec![0, 1, 2],
            },
            TestCase {
                nums: vec![2, 0, 2, 1, 1, 0],
                expected: vec![0, 0, 1, 1, 2, 2],
            },
        ];

        for test in tests {
            let mut input = test.nums.clone();
            Solution::sort_colors(&mut input);

            assert_eq!(
                input, test.expected,
                "test case {:?} was {:?} not {:?}",
                test.nums, input, test.expected
            );
        }
    }
}
