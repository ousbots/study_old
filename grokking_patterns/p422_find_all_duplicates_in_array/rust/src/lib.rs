pub struct Solution;
impl Solution {
    // Returns a vector of the duplicate integers in the nums vector. The integers in nums range
    // from 1..n and each integer is seen at most twice.
    //
    // This solution uses a modified cyclic sort. Because the integers are in a sequence that is
    // the length of the array, there isn't any need to count the correct position of any given
    // integer. Instead, the array can be looped through and for each element if it is not in the
    // right position, swap it to the correct position (nums[i - 1]) and continue swapping until
    // the current position has an integer whose correct position is already filled. After the
    // vector is sorted, it is iterated through again and any position that has an incorrect
    // integer is swapped to the next front position. Then the subarry of integers swapped to the
    // front is returned.
    pub fn find_duplicates(nums: Vec<i32>) -> Vec<i32> {
        let mut nums = nums;

        let mut index;
        for start in 0..nums.len() {
            index = usize::try_from(nums[start] - 1).unwrap();
            while nums[index] != nums[start] {
                (nums[index], nums[start]) = (nums[start], nums[index]);
                index = usize::try_from(nums[start] - 1).unwrap();
            }
        }

        let mut count = 0;
        for i in 0..nums.len() {
            if nums[i] != i32::try_from(i + 1).unwrap() {
                nums[count] = nums[i];
                count += 1;
            }
        }

        nums[0..count].to_vec()
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
                nums: vec![4, 3, 2, 7, 8, 2, 3, 1],
                expected: vec![2, 3],
            },
            TestCase {
                nums: vec![1, 1, 2],
                expected: vec![1],
            },
            TestCase {
                nums: vec![2, 2],
                expected: vec![2],
            },
            TestCase {
                nums: vec![1],
                expected: vec![],
            },
        ];

        for test in tests {
            let mut ans = Solution::find_duplicates(test.nums.clone());
            ans.sort();
            assert_eq!(ans, test.expected, "case {:?}", test.nums);
        }
    }
}
