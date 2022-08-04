use std::collections::HashMap;

pub struct Solution;
impl Solution {
    // Returns all tuples of four elements from nums whose sum equals the target. The tuples must
    // be unique.
    //
    // This solution uses a two pointers variant. There are two "fixed" points that iterate over
    // the vector and every loop a front and back pointer are created to search for the optimal
    // elements. If the sum is less than the target, the back pointer is incremented, if the sum
    // is greater than the target, the front pointer is decremented. A hashmap is used to check the
    // uniqueness of each matching set. Even if the set isn't unique, a matching set still requires
    // one of the pointers to be moved.
    pub fn four_sum(nums: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        let target = i64::from(target);
        if nums.len() < 4 {
            return vec![];
        }

        let mut nums = nums;
        nums.sort();

        let mut sets = vec![];
        let mut seen: HashMap<[i32; 4], bool> = HashMap::new();

        for i in 0..nums.len() - 2 {
            for j in i + 1..nums.len() - 1 {
                let mut back = j + 1;
                let mut front = nums.len() - 1;

                while back < front {
                    let sum = i64::from(nums[i])
                        + i64::from(nums[back])
                        + i64::from(nums[front])
                        + i64::from(nums[j]);

                    if sum < target {
                        back += 1;
                    } else if sum > target {
                        front -= 1;
                    } else {
                        let set = [nums[i], nums[back], nums[front], nums[j]];
                        if !seen.contains_key(&set) {
                            sets.push(set.to_vec());
                            seen.insert(set, true);
                        }
                        back += 1;
                    }
                }
            }
        }

        sets
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    struct TestCase {
        nums: Vec<i32>,
        target: i32,
        expected: Vec<Vec<i32>>,
    }

    #[test]
    fn basic() {
        let tests = vec![
            TestCase {
                nums: vec![1000000000, 1000000000, 1000000000, 1000000000],
                target: -294967296,
                expected: vec![],
            },
            TestCase {
                nums: vec![0],
                target: 0,
                expected: vec![],
            },
            TestCase {
                nums: vec![1, 1, 1, 1, 1, 1],
                target: 4,
                expected: vec![vec![1, 1, 1, 1]],
            },
            TestCase {
                nums: vec![1, 0, -1, 0, -2, 2],
                target: 0,
                expected: vec![vec![-2, 0, 0, 2], vec![-1, 0, 0, 1], vec![-2, -1, 1, 2]],
            },
            TestCase {
                nums: vec![1, 0, -1, 0, -2, 2],
                target: 0,
                expected: vec![vec![-2, 0, 0, 2], vec![-1, 0, 0, 1], vec![-2, -1, 1, 2]],
            },
        ];

        for mut test in tests {
            let mut ans = Solution::four_sum(test.nums.clone(), test.target);
            for i in 0..ans.len() {
                ans[i].sort();
            }
            ans.sort();

            for i in 0..test.expected.len() {
                test.expected[i].sort();
            }
            test.expected.sort();

            assert_eq!(
                ans, test.expected,
                "case nums={:?} target={}",
                test.nums, test.target
            )
        }
    }
}
