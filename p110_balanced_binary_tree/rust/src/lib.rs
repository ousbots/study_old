use std::cell::RefCell;
use std::rc::Rc;

// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

pub struct Solution;

impl Solution {
    pub fn is_balanced(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        distance(root) != -1
    }
}

// Surprisingly, the borrowing in this function is much slower than the cloning in distance() below.
fn _distance_slow(root: &Option<Rc<RefCell<TreeNode>>>) -> i32 {
    if let Some(node) = root {
        let left = _distance_slow(&node.borrow().left);
        if left == -1 {
            return -1;
        }

        let right = _distance_slow(&node.borrow().right);
        if right == -1 {
            return -1;
        }

        let total = left - right;
        if total < -1 || total > 1 {
            return -1;
        }

        if left > right {
            return left + 1;
        } else {
            return right + 1;
        }
    } else {
        return 0;
    }
}

// Much faster to clone than borrow, I'm not sure why.
fn distance(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    if let Some(node) = root {
        let left = distance(node.borrow().left.clone());
        if left == -1 {
            return -1;
        }

        let right = distance(node.borrow().right.clone());
        if right == -1 {
            return -1;
        }

        let total = (left - right).abs();
        if total > 1 {
            return -1;
        }

        return left.max(right) + 1;
    } else {
        return 0;
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;
    use super::TreeNode;
    use std::cell::RefCell;
    use std::rc::Rc;

    #[test]
    fn provided() {
        let input = Some(Rc::new(RefCell::new(TreeNode {
            val: 3,
            left: Some(Rc::new(RefCell::new(TreeNode {
                val: 9,
                left: None,
                right: None,
            }))),
            right: Some(Rc::new(RefCell::new(TreeNode {
                val: 20,
                left: Some(Rc::new(RefCell::new(TreeNode {
                    val: 15,
                    left: None,
                    right: None,
                }))),
                right: Some(Rc::new(RefCell::new(TreeNode {
                    val: 7,
                    left: None,
                    right: None,
                }))),
            }))),
        })));
        assert_eq!(Solution::is_balanced(input), true);

        let input = Some(Rc::new(RefCell::new(TreeNode {
            val: 1,
            left: Some(Rc::new(RefCell::new(TreeNode {
                val: 2,
                left: Some(Rc::new(RefCell::new(TreeNode {
                    val: 3,
                    left: Some(Rc::new(RefCell::new(TreeNode {
                        val: 4,
                        left: None,
                        right: None,
                    }))),
                    right: Some(Rc::new(RefCell::new(TreeNode {
                        val: 4,
                        left: None,
                        right: None,
                    }))),
                }))),
                right: Some(Rc::new(RefCell::new(TreeNode {
                    val: 3,
                    left: None,
                    right: None,
                }))),
            }))),
            right: Some(Rc::new(RefCell::new(TreeNode {
                val: 2,
                left: None,
                right: None,
            }))),
        })));
        assert_eq!(Solution::is_balanced(input), false);

        let input = None;
        assert_eq!(Solution::is_balanced(input), true);
    }
}
