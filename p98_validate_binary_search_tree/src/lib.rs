use std::cell::RefCell;
use std::rc::Rc;

pub struct Solution;

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

impl Solution {
    // Checks that the submitted tree is a binary search tree.
    pub fn is_valid_bst(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        fn validate(
            root: Option<Rc<RefCell<TreeNode>>>,
            min: Option<Rc<RefCell<TreeNode>>>,
            max: Option<Rc<RefCell<TreeNode>>>,
        ) -> bool {
            match root {
                None => return true,
                Some(node) => {
                    if let Some(left) = min.clone() {
                        if left.borrow().val >= node.borrow().val {
                            return false;
                        }
                    }

                    if let Some(right) = max.clone() {
                        if right.borrow().val <= node.borrow().val {
                            return false;
                        }
                    }

                    return validate(node.borrow().left.clone(), min, Some(node.clone()))
                        && validate(node.borrow().right.clone(), Some(node.clone()), max);
                }
            }
        }

        validate(root, None, None)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn provided_tests() {
        let root = Some(Rc::new(RefCell::new(TreeNode {
            val: 2,
            left: Some(Rc::new(RefCell::new(TreeNode {
                val: 1,
                left: None,
                right: None,
            }))),
            right: Some(Rc::new(RefCell::new(TreeNode {
                val: 3,
                left: None,
                right: None,
            }))),
        })));
        assert_eq!(Solution::is_valid_bst(root), true);

        let root = Some(Rc::new(RefCell::new(TreeNode {
            val: 5,
            left: Some(Rc::new(RefCell::new(TreeNode {
                val: 1,
                left: None,
                right: None,
            }))),
            right: Some(Rc::new(RefCell::new(TreeNode {
                val: 4,
                left: Some(Rc::new(RefCell::new(TreeNode {
                    val: 3,
                    left: None,
                    right: None,
                }))),
                right: Some(Rc::new(RefCell::new(TreeNode {
                    val: 6,
                    left: None,
                    right: None,
                }))),
            }))),
        })));
        assert_eq!(Solution::is_valid_bst(root), false);
    }

    #[test]
    fn extra_tests() {
        assert_eq!(Solution::is_valid_bst(None), true);

        let root = Some(Rc::new(RefCell::new(TreeNode {
            val: 2,
            left: Some(Rc::new(RefCell::new(TreeNode {
                val: 2,
                left: None,
                right: None,
            }))),
            right: Some(Rc::new(RefCell::new(TreeNode {
                val: 2,
                left: None,
                right: None,
            }))),
        })));
        assert_eq!(Solution::is_valid_bst(root), false);

        let root = Some(Rc::new(RefCell::new(TreeNode {
            val: 5,
            left: Some(Rc::new(RefCell::new(TreeNode {
                val: 4,
                left: None,
                right: None,
            }))),
            right: Some(Rc::new(RefCell::new(TreeNode {
                val: 6,
                left: Some(Rc::new(RefCell::new(TreeNode {
                    val: 3,
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
        assert_eq!(Solution::is_valid_bst(root), false);
    }
}
