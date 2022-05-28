use std::cell::RefCell;
use std::rc::Rc;

// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

pub struct Solution;
impl Solution {
    pub fn lowest_common_ancestor(
        root: Option<Rc<RefCell<TreeNode>>>,
        p: Option<Rc<RefCell<TreeNode>>>,
        q: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        if let Some(root) = root {
            if root.borrow().val == p.as_ref().unwrap().borrow().val
                || root.borrow().val == q.as_ref().unwrap().borrow().val
            {
                return Some(root);
            }

            let left =
                Solution::lowest_common_ancestor(root.borrow().left.clone(), p.clone(), q.clone());
            let right =
                Solution::lowest_common_ancestor(root.borrow().right.clone(), p.clone(), q.clone());

            if left.is_none() && right.is_none() {
                return None;
            } else if left.is_some() && right.is_some() {
                return Some(root);
            } else if left.is_some() {
                return left;
            } else {
                return right;
            }
        }

        None
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
            val: 3,
            left: Some(Rc::new(RefCell::new(TreeNode {
                val: 5,
                left: Some(Rc::new(RefCell::new(TreeNode {
                    val: 6,
                    left: None,
                    right: None,
                }))),
                right: Some(Rc::new(RefCell::new(TreeNode {
                    val: 2,
                    left: Some(Rc::new(RefCell::new(TreeNode {
                        val: 7,
                        left: None,
                        right: None,
                    }))),
                    right: Some(Rc::new(RefCell::new(TreeNode {
                        val: 4,
                        left: None,
                        right: None,
                    }))),
                }))),
            }))),
            right: Some(Rc::new(RefCell::new(TreeNode {
                val: 1,
                left: Some(Rc::new(RefCell::new(TreeNode {
                    val: 0,
                    left: None,
                    right: None,
                }))),
                right: Some(Rc::new(RefCell::new(TreeNode {
                    val: 8,
                    left: None,
                    right: None,
                }))),
            }))),
        })));
        assert_eq!(
            Solution::lowest_common_ancestor(
                input.clone(),
                Some(Rc::new(RefCell::new(TreeNode {
                    val: 5,
                    left: None,
                    right: None
                }))),
                Some(Rc::new(RefCell::new(TreeNode {
                    val: 1,
                    left: None,
                    right: None
                })))
            )
            .unwrap()
            .borrow()
            .val,
            3
        );

        assert_eq!(
            Solution::lowest_common_ancestor(
                input.clone(),
                Some(Rc::new(RefCell::new(TreeNode {
                    val: 5,
                    left: None,
                    right: None
                }))),
                Some(Rc::new(RefCell::new(TreeNode {
                    val: 4,
                    left: None,
                    right: None
                })))
            )
            .unwrap()
            .borrow()
            .val,
            5
        );

        let input = Some(Rc::new(RefCell::new(TreeNode {
            val: 1,
            left: Some(Rc::new(RefCell::new(TreeNode {
                val: 2,
                left: None,
                right: None,
            }))),
            right: None,
        })));
        assert_eq!(
            Solution::lowest_common_ancestor(
                input,
                Some(Rc::new(RefCell::new(TreeNode {
                    val: 1,
                    left: None,
                    right: None
                }))),
                Some(Rc::new(RefCell::new(TreeNode {
                    val: 2,
                    left: None,
                    right: None
                })))
            )
            .unwrap()
            .borrow()
            .val,
            1
        );
    }
}
