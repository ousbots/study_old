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
    pub fn invert_tree(mut root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        Self::swap(&mut root);

        root
    }

    pub fn swap(root: &mut Option<Rc<RefCell<TreeNode>>>) {
        if let Some(node) = root {
            let left = std::mem::replace(&mut node.borrow_mut().left, None);
            let right = std::mem::replace(&mut node.borrow_mut().right, left);
            node.borrow_mut().left = right;

            Self::swap(&mut node.borrow_mut().left);
            Self::swap(&mut node.borrow_mut().right);
        }
    }
}

#[cfg(test)]
mod tests {
    use super::{Solution, TreeNode};
    use std::cell::RefCell;
    use std::rc::Rc;

    #[test]
    fn provided() {
        let input = Some(Rc::new(RefCell::new(TreeNode {
            val: 4,
            left: Some(Rc::new(RefCell::new(TreeNode {
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
            }))),
            right: Some(Rc::new(RefCell::new(TreeNode {
                val: 7,
                left: Some(Rc::new(RefCell::new(TreeNode {
                    val: 6,
                    left: None,
                    right: None,
                }))),
                right: Some(Rc::new(RefCell::new(TreeNode {
                    val: 9,
                    left: None,
                    right: None,
                }))),
            }))),
        })));

        let ans = Some(Rc::new(RefCell::new(TreeNode {
            val: 4,
            left: Some(Rc::new(RefCell::new(TreeNode {
                val: 7,
                left: Some(Rc::new(RefCell::new(TreeNode {
                    val: 9,
                    left: None,
                    right: None,
                }))),
                right: Some(Rc::new(RefCell::new(TreeNode {
                    val: 6,
                    left: None,
                    right: None,
                }))),
            }))),
            right: Some(Rc::new(RefCell::new(TreeNode {
                val: 2,
                left: Some(Rc::new(RefCell::new(TreeNode {
                    val: 3,
                    left: None,
                    right: None,
                }))),
                right: Some(Rc::new(RefCell::new(TreeNode {
                    val: 1,
                    left: None,
                    right: None,
                }))),
            }))),
        })));

        let output = Solution::invert_tree(input);
        assert_eq!(check(output, ans), true);
    }

    fn check(first: Option<Rc<RefCell<TreeNode>>>, second: Option<Rc<RefCell<TreeNode>>>) -> bool {
        if first.is_none() && second.is_none() {
            return true;
        }

        if let Some(first) = first {
            if let Some(second) = second {
                let f = first.borrow();
                let s = second.borrow();

                println!("checking {} == {}", f.val, s.val);

                return f.val == s.val
                    && check(f.left.clone(), s.left.clone())
                    && check(f.right.clone(), s.right.clone());
            }
        }

        false
    }
}
