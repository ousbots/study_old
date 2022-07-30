use std::cell::RefCell;
use std::rc::Rc;

pub struct Solution;

#[derive(Clone, Debug)]
struct Node {
    val: i32,
    count: i32,
    left: Option<Rc<RefCell<Node>>>,
    right: Option<Rc<RefCell<Node>>>,
}

impl Node {
    // Insert the value into the BST.
    fn insert(&mut self, value: i32) -> i32 {
        if value <= self.val {
            self.count += 1;
            if let Some(ref mut left) = self.left {
                return left.borrow_mut().insert(value);
            } else {
                self.left = Some(Rc::new(RefCell::new(Node {
                    val: value,
                    count: 1,
                    left: None,
                    right: None,
                })));

                return 0;
            }
        } else {
            if let Some(ref mut right) = self.right {
                return self.count + right.borrow_mut().insert(value);
            } else {
                self.right = Some(Rc::new(RefCell::new(Node {
                    val: value,
                    count: 1,
                    left: None,
                    right: None,
                })));

                return self.count;
            }
        }
    }
}

impl Solution {
    // Generates an array such that result[i] is the count of all elements less
    // than nums[i] for i+1 to nums.len().
    pub fn count_smaller(nums: Vec<i32>) -> Vec<i32> {
        if nums.len() == 0 {
            return vec![];
        }

        let mut bst = Box::new(Node {
            val: nums[nums.len() - 1],
            count: 1,
            left: None,
            right: None,
        });
        let mut count = vec![0];

        for index in (0..nums.len() - 1).rev() {
            count.push(bst.insert(nums[index]));
        }
        count.reverse();

        count
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn provided_tests() {
        assert_eq!(Solution::count_smaller(vec![5, 2, 6, 1]), [2, 1, 1, 0]);
        assert_eq!(Solution::count_smaller(vec![-1]), [0]);
        assert_eq!(Solution::count_smaller(vec![-1, -1]), [0, 0]);
    }

    #[test]
    fn more_tests() {
        assert_eq!(
            Solution::count_smaller(vec![5, 4, 3, 2, 1]),
            [4, 3, 2, 1, 0]
        );
        assert_eq!(
            Solution::count_smaller(vec![1, 2, 3, 4, 5]),
            [0, 0, 0, 0, 0]
        );
        assert_eq!(Solution::count_smaller(vec![1, 0, -1]), [2, 1, 0]);
        assert_eq!(Solution::count_smaller(vec![2, 0, 1]), [2, 0, 0]);

        assert_eq!(
            Solution::count_smaller(vec![
                26, 78, 27, 100, 33, 67, 90, 23, 66, 5, 38, 7, 35, 23, 52, 22, 83, 51, 98, 69, 81,
                32, 78, 28, 94, 13, 2, 97, 3, 76, 99, 51, 9, 21, 84, 66, 65, 36, 100, 41
            ]),
            [
                10, 27, 10, 35, 12, 22, 28, 8, 19, 2, 12, 2, 9, 6, 12, 5, 17, 9, 19, 12, 14, 6, 12,
                5, 12, 3, 0, 10, 0, 7, 8, 4, 0, 0, 4, 3, 2, 0, 1, 0
            ]
        );
    }
}
