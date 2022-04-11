pub struct Solution;

impl Solution {
    pub fn count_range_sum(nums: Vec<i32>, lower: i32, upper: i32) -> i32 {
        let length = nums.len();
        let mut sums: Vec<i64> = vec![0; length + 1];

        for index in 0..nums.len() {
            sums[index + 1] = sums[index] + i64::from(nums[index]);
        }

        count_merge_sort(&mut sums, 0, length + 1, lower, upper)
    }
}

fn count_merge_sort(sums: &mut Vec<i64>, start: usize, end: usize, lower: i32, upper: i32) -> i32 {
    if end - start <= 1 {
        return 0;
    }

    let middle = (start + end) / 2;

    let mut count = count_merge_sort(sums, start, middle, lower, upper)
        + count_merge_sort(sums, middle, end, lower, upper);

    let mut j = middle;
    let mut k = middle;

    for index in start..middle {
        while j < end && sums[j] - sums[index] < i64::from(lower) {
            j += 1;
        }

        while k < end && sums[k] - sums[index] <= i64::from(upper) {
            k += 1;
        }

        count += (k - j) as i32;
    }

    sums[start..end].sort_unstable();

    count
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(Solution::count_range_sum(vec![-2, 5, -1], -2, 2), 3);
        assert_eq!(Solution::count_range_sum(vec![0], 0, 0), 1);
    }
}
