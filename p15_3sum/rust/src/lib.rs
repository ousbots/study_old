pub struct Solution;

impl Solution {
    pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut sets: Vec<Vec<i32>> = vec![];
        if nums.len() < 3 {
            return sets;
        }

        let mut nums = nums;
        nums.sort_unstable();

        for i in 0..nums.len() - 2 {
            let mut low = i + 1;
            let mut high = nums.len() - 1;

            while low < high {
                let sum = nums[i] + nums[low] + nums[high];
                if sum == 0 {
                    // Check for duplicates by looking back at any previous entries that have
                    // the same starting number.
                    let mut found = false;
                    for check in (0..sets.len()).rev() {
                        if sets[check][0] == nums[i] {
                            if sets[check][1] == nums[low] && sets[check][2] == nums[high] {
                                found = true;
                                break;
                            }
                        }
                    }

                    if !found {
                        sets.push(vec![nums[i], nums[low], nums[high]]);
                    }

                    low += 1;
                    high -= 1;
                } else if sum < 0 {
                    low += 1;
                } else {
                    high -= 1;
                }
            }
        }

        sets
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        let mut ans = Solution::three_sum(vec![-1, 0, 1, 2, -1, -4]);
        for elem in &mut ans {
            elem.sort();
        }
        let mut expect = vec![vec![-1, -1, 2], vec![-1, 0, 1]];
        for elem in &mut expect {
            elem.sort()
        }
        assert_eq!(ans, expect);

        let empty: Vec<Vec<i32>> = Vec::new();
        assert_eq!(Solution::three_sum(vec![]), empty);
        assert_eq!(Solution::three_sum(vec![0]), empty);
    }

    #[test]
    fn basic() {
        assert_eq!(Solution::three_sum(vec![0, 0, 0]), vec![vec![0, 0, 0]]);
        assert_eq!(Solution::three_sum(vec![0, 0, 0, 0]), vec![vec![0, 0, 0]]);

        let mut ans = Solution::three_sum(vec![1, -1, 0]);
        for elem in &mut ans {
            elem.sort();
        }
        let mut expect = vec![vec![1, -1, 0]];
        for elem in &mut expect {
            elem.sort()
        }
        assert_eq!(ans, expect);

        let mut ans = Solution::three_sum(vec![1, 1, -1, 1, -1, -1, 2]);
        for elem in &mut ans {
            elem.sort();
        }
        let mut expect = vec![vec![2, -1, -1]];
        for elem in &mut expect {
            elem.sort()
        }
        assert_eq!(ans, expect);

        let mut ans = Solution::three_sum(vec![
            34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56,
            94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55,
            -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49,
        ]);
        for elem in &mut ans {
            elem.sort();
        }
        let mut expect = vec![
            vec![-82, -11, 93],
            vec![-82, 13, 69],
            vec![-82, 17, 65],
            vec![-82, 21, 61],
            vec![-82, 26, 56],
            vec![-82, 33, 49],
            vec![-82, 34, 48],
            vec![-82, 36, 46],
            vec![-70, -14, 84],
            vec![-70, -6, 76],
            vec![-70, 1, 69],
            vec![-70, 13, 57],
            vec![-70, 15, 55],
            vec![-70, 21, 49],
            vec![-70, 34, 36],
            vec![-66, -11, 77],
            vec![-66, -3, 69],
            vec![-66, 1, 65],
            vec![-66, 10, 56],
            vec![-66, 17, 49],
            vec![-49, -6, 55],
            vec![-49, -3, 52],
            vec![-49, 1, 48],
            vec![-49, 2, 47],
            vec![-49, 13, 36],
            vec![-49, 15, 34],
            vec![-49, 21, 28],
            vec![-43, -14, 57],
            vec![-43, -6, 49],
            vec![-43, -3, 46],
            vec![-43, 10, 33],
            vec![-43, 12, 31],
            vec![-43, 15, 28],
            vec![-43, 17, 26],
            vec![-29, -14, 43],
            vec![-29, 1, 28],
            vec![-29, 12, 17],
            vec![-14, -3, 17],
            vec![-14, 1, 13],
            vec![-14, 2, 12],
            vec![-11, -6, 17],
            vec![-11, 1, 10],
            vec![-3, 1, 2],
        ];
        for elem in &mut expect {
            elem.sort()
        }
        assert_eq!(ans, expect);
    }
}
