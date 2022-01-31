use std::cmp::Reverse;
use std::collections::BinaryHeap;

pub struct Solution;

impl Solution {
    pub fn k_smallest_pairs(nums1: Vec<i32>, nums2: Vec<i32>, k: i32) -> Vec<Vec<i32>> {
        let mut heap: BinaryHeap<Reverse<(i32, usize, usize)>> = BinaryHeap::from(
            (0..std::cmp::min(k as usize, nums1.len()))
                .map(|i| Reverse((nums1[i] + nums2[0], i, 0)))
                .collect::<Vec<Reverse<(i32, usize, usize)>>>(),
        );

        let mut pairs = vec![];
        while pairs.len() < k as usize {
            if let Some(element) = heap.pop() {
                let (_, first, second) = element.0;
                pairs.push(vec![nums1[first], nums2[second]]);

                if second < nums2.len() - 1 {
                    heap.push(Reverse((
                        nums1[first] + nums2[second + 1],
                        first,
                        second + 1,
                    )));
                }
            } else {
                break;
            }
        }

        println!("nums1 {:?} nums2 {:?}", nums1, nums2);
        println!("pairs {:?}", pairs);

        pairs
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn provided_tests() {
        assert_eq!(
            Solution::k_smallest_pairs(vec![1, 7, 11], vec![2, 4, 6], 3),
            [[1, 2], [1, 4], [1, 6]]
        );

        assert_eq!(
            Solution::k_smallest_pairs(vec![1, 1, 2], vec![1, 2, 3], 2),
            [[1, 1], [1, 1]]
        );

        assert_eq!(
            Solution::k_smallest_pairs(vec![1, 2], vec![3], 3),
            [[1, 3], [2, 3]]
        );
    }

    #[test]
    fn basic_tests() {
        assert_eq!(Solution::k_smallest_pairs(vec![-1], vec![1], 2), [[-1, 1]]);

        let mut solution = Solution::k_smallest_pairs(vec![1, 1, 2], vec![1, 2, 3], 10);
        let mut answer = [
            [1, 1],
            [1, 1],
            [2, 1],
            [1, 2],
            [1, 2],
            [2, 2],
            [1, 3],
            [1, 3],
            [2, 3],
        ];

        solution.sort();
        answer.sort();
        assert_eq!(solution, answer);
    }
}
