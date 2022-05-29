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
    pub fn right_side_view(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        if let Some(root) = root {
            let mut right = Solution::right_side_view(root.borrow().right.clone());
            let left = Solution::right_side_view(root.borrow().left.clone());

            if left.len() > right.len() {
                for i in right.len()..left.len() {
                    right.push(left[i]);
                }
            }

            let mut output = vec![root.borrow().val];
            output.append(&mut right);

            return output;
        } else {
            return vec![];
        }
    }
}

#[cfg(test)]
mod tests {
    use super::{Solution, TreeNode};
    use std::cell::RefCell;
    use std::rc::Rc;

    #[test]
    fn basic() {
        let valid = vec![];
        assert_eq!(Solution::right_side_view(None), valid);

        let input = Some(Rc::new(RefCell::new(TreeNode {
            val: 1,
            left: None,
            right: None,
        })));
        let valid = vec![1];
        assert_eq!(Solution::right_side_view(input), valid);

        let input = Some(Rc::new(RefCell::new(TreeNode {
            val: 1,
            left: None,
            right: Some(Rc::new(RefCell::new(TreeNode {
                val: 2,
                left: None,
                right: None,
            }))),
        })));
        let valid = vec![1, 2];
        assert_eq!(Solution::right_side_view(input), valid);

        let input = Some(Rc::new(RefCell::new(TreeNode {
            val: 1,
            left: Some(Rc::new(RefCell::new(TreeNode {
                val: 2,
                left: None,
                right: None,
            }))),
            right: None,
        })));
        let valid = vec![1, 2];
        assert_eq!(Solution::right_side_view(input), valid);

        let input = Some(Rc::new(RefCell::new(TreeNode {
            val: 1,
            left: Some(Rc::new(RefCell::new(TreeNode {
                val: 2,
                left: Some(Rc::new(RefCell::new(TreeNode {
                    val: 4,
                    left: Some(Rc::new(RefCell::new(TreeNode {
                        val: 5,
                        left: None,
                        right: None,
                    }))),
                    right: None,
                }))),
                right: Some(Rc::new(RefCell::new(TreeNode {
                    val: 3,
                    left: None,
                    right: None,
                }))),
            }))),
            right: Some(Rc::new(RefCell::new(TreeNode {
                val: 9,
                left: None,
                right: None,
            }))),
        })));
        let valid = vec![1, 9, 3, 5];
        assert_eq!(Solution::right_side_view(input), valid);
    }
}
