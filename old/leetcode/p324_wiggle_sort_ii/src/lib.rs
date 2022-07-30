pub struct Solution;

// Quicksort partition using the Lomuto scheme.
fn partition(nums: &mut Vec<i32>, start: usize, end: usize) -> usize {
    let pivot_val = nums[end];
    let mut pivot_index = start;

    for current in start..end {
        if nums[current] <= pivot_val {
            nums.swap(pivot_index, current);
            pivot_index += 1;
        }
    }
    nums.swap(pivot_index, end);

    pivot_index
}

// Quick select finds the kth smallest element in the given array using the quick sort algorithm
// modified to only continue in the partition that contains the kth element.
fn quick_select(nums: &mut Vec<i32>, kth: usize, start: usize, end: usize) -> i32 {
    if start == end {
        return nums[kth];
    }

    let pivot_index = partition(nums, start, end);

    if pivot_index > kth {
        return quick_select(nums, kth, start, pivot_index - 1);
    } else if pivot_index < kth {
        return quick_select(nums, kth, pivot_index + 1, end);
    } else {
        return nums[pivot_index];
    }
}

// Maps the index of a three way sort to wiggle sort.
fn index_map(index: usize, length: usize) -> usize {
    (1 + 2 * index) % (length | 1)
}

impl Solution {
    // Wiggle sort the vector of numbers: nums[n-1] < nums[n] < nums[n+1].
    // Note: this is a non-optimal algorithm that doesn't work in special cases.
    pub fn wiggle_sort(nums: &mut Vec<i32>) {
        let median = quick_select(nums, ((nums.len() + 1) / 2) - 1, 0, nums.len() - 1);

        let mut index = 0;
        let mut start = 0;
        let length = nums.len();
        let mut end = length - 1;

        while start <= end {
            if nums[index_map(start, length)] > median {
                nums.swap(index_map(index, length), index_map(start, length));
                index += 1;
                start += 1;
            } else if nums[index_map(start, nums.len())] < median {
                nums.swap(index_map(start, length), index_map(end, length));
                end -= 1;
            } else {
                start += 1;
            }
        }
    }
}

// Check that the numbers are wiggle sorted: nums[n-1] < nums[n] < nums[n+1].
pub fn check_wiggle(nums: Vec<i32>) -> bool {
    for index in 0..nums.len() - 1 {
        if index % 2 == 0 {
            if nums[index] >= nums[index + 1] {
                return false;
            }
        } else {
            if nums[index] <= nums[index + 1] {
                return false;
            }
        }
    }

    true
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn provided_tests() {
        let mut data = vec![1, 5, 1, 1, 6, 4];
        Solution::wiggle_sort(&mut data);
        assert_eq!(check_wiggle(data), true);

        data = vec![1, 3, 2, 2, 3, 1];
        Solution::wiggle_sort(&mut data);
        assert_eq!(check_wiggle(data), true);
    }

    #[test]
    fn more_tests() {
        let mut data = vec![1, 1, 2, 2, 2, 1];
        Solution::wiggle_sort(&mut data);
        assert_eq!(check_wiggle(data), true);

        let mut data = vec![1, 1, 2, 1, 2, 2, 1];
        Solution::wiggle_sort(&mut data);
        assert_eq!(check_wiggle(data), true);

        let mut data = vec![4, 5, 5, 6];
        Solution::wiggle_sort(&mut data);
        assert_eq!(check_wiggle(data), true);

        let mut data = vec![1];
        Solution::wiggle_sort(&mut data);
        assert_eq!(check_wiggle(data), true);

        let mut data = vec![4, 5, 5, 5, 5, 6, 6, 6];
        Solution::wiggle_sort(&mut data);
        assert_eq!(check_wiggle(data), true);
    }
}
