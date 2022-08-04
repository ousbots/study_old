pub struct Solution;
impl Solution {
    // Returns the number of contiguous subarrays whose product is less than k.
    //
    // This solution uses two pointers where the rear pointer iterates over the nums array and for
    // every iteration the front pointer iterates forward until the product is greater than k. This
    // is not the most efficient solution, it's used to study the two pointer pattern.
    pub fn num_subarray_product_less_than_k(nums: Vec<i32>, k: i32) -> i32 {
        let mut count = 0;
        let mut back = 0;

        while back < nums.len() {
            let mut product = nums[back];
            if product < k {
                count += 1;
            }

            let mut front = back + 1;
            while product < k && front < nums.len() {
                product *= nums[front];
                if product < k {
                    count += 1;
                }

                front += 1;
            }

            back += 1;
        }

        count
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    struct TestCase {
        nums: Vec<i32>,
        k: i32,
        expected: i32,
    }

    #[test]
    fn basic() {
        let tests = vec![
            TestCase {
                nums: vec![10, 5, 2, 6],
                k: 100,
                expected: 8,
            },
            TestCase {
                nums: vec![1, 2, 3],
                k: 0,
                expected: 0,
            },
            TestCase {
                nums: vec![0, 1, 2, 3, 4, 5],
                k: 1,
                expected: 6,
            },
        ];

        for test in tests {
            assert_eq!(
                Solution::num_subarray_product_less_than_k(test.nums.clone(), test.k),
                test.expected
            );
        }
    }
}
