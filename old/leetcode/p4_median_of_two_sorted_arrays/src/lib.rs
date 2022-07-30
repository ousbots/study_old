pub struct Solution;

impl Solution {
    pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
        let mut nums1 = nums1;
        let mut nums2 = nums2;
        if nums1.len() > nums2.len() {
            let tmp = nums1;
            nums1 = nums2;
            nums2 = tmp;
        }

        let length = nums1.len() + nums2.len();
        let mut left = 0;
        let mut right = nums1.len();

        while left <= right {
            let i = (left + right) / 2;
            let j = length / 2 - i;

            if i < nums1.len() && nums1[i] < nums2[j - 1] {
                left = i + 1;
            } else if i > 0 && nums1[i - 1] > nums2[j] {
                right = i - 1;
            } else {
                let min = if i == nums1.len() {
                    nums2[j]
                } else if j == nums2.len() {
                    nums1[i]
                } else {
                    nums1[i].min(nums2[j])
                };

                if length % 2 == 0 {
                    let max = if i == 0 {
                        nums2[j - 1]
                    } else if j == 0 {
                        nums1[i - 1]
                    } else {
                        nums1[i - 1].max(nums2[j - 1])
                    };
                    return (max + min) as f64 / 2.;
                } else {
                    return min as f64;
                }
            }
        }

        0.
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided_tests() {
        assert_eq!(
            Solution::find_median_sorted_arrays(vec![1, 3], vec![2]),
            2.0
        );

        assert_eq!(
            Solution::find_median_sorted_arrays(vec![1, 2], vec![3, 4]),
            2.5
        );
    }

    #[test]
    fn edge_case_tests() {
        assert_eq!(Solution::find_median_sorted_arrays(vec![], vec![]), 0.);
        assert_eq!(Solution::find_median_sorted_arrays(vec![2], vec![]), 2.);
        assert_eq!(Solution::find_median_sorted_arrays(vec![2], vec![]), 2.);

        assert_eq!(
            Solution::find_median_sorted_arrays(vec![1, 1], vec![9, 9, 9, 9, 9]),
            9.0
        );

        assert_eq!(
            Solution::find_median_sorted_arrays(vec![1, 1, 1, 1, 1], vec![0]),
            1.0
        );
    }

    #[test]
    fn basic_tests() {
        assert_eq!(Solution::find_median_sorted_arrays(vec![], vec![3, 4]), 2.5);
    }
}
