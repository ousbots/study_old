pub struct Solution;
impl Solution {
    // Returns a vector of length two that contains the missing and the duplicate integers if nums
    // is supposed to range from 1..n.
    //
    // This solution uses a cyclic sort to swap the integers into the correct positions in constant
    // memory, then iterates through the sorted list until a wrong element is found (i.e.
    // nums[i] != i+1). The wrong element is both the duplicate and the missing integers (i.e.
    // nums[i] and i+1).
    pub fn find_error_nums(nums: &mut Vec<i32>) -> Vec<i32> {
        let mut temp: usize;
        for i in 0..nums.len() {
            while nums[nums[i] as usize - 1] != nums[i] {
                temp = nums[i] as usize - 1;
                (nums[i], nums[temp]) = (nums[temp], nums[i]);
            }
        }

        for i in 0..nums.len() {
            if nums[i] != (i + 1) as i32 {
                return vec![nums[i], (i + 1) as i32];
            }
        }

        return vec![];
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
                nums: vec![1, 2, 2, 4],
                expected: vec![2, 3],
            },
            TestCase {
                nums: vec![1, 1],
                expected: vec![1, 2],
            },
        ];

        for test in tests {
            assert_eq!(
                Solution::find_error_nums(&mut test.nums.clone()),
                test.expected,
                "case {:?}",
                test.nums
            );
        }
    }
}
