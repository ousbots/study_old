use std::cell::RefCell;
use std::collections::VecDeque;
use std::rc::Rc;

// Binary tree node definition.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

pub struct Solution;
impl Solution {
    pub fn level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        let mut results: Vec<Vec<i32>> = vec![];
        let mut next: VecDeque<Option<Rc<RefCell<TreeNode>>>> = VecDeque::from([root]);

        while next.len() != 0 {
            let mut values: Vec<i32> = vec![];
            let len = next.len();

            for _ in 0..len {
                if let Some(elem) = next.pop_front().unwrap() {
                    let node = elem.borrow();
                    values.push(node.val);
                    next.push_back(node.left.clone());
                    next.push_back(node.right.clone());
                }
            }

            results.push(values)
        }
        results.pop();

        results
    }

    pub fn level_order_recursive(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        let mut results: Vec<Vec<i32>> = vec![];

        fn build(root: Option<Rc<RefCell<TreeNode>>>, results: &mut Vec<Vec<i32>>, depth: usize) {
            if let Some(root) = root {
                if results.len() == depth {
                    results.push(vec![]);
                }

                results[depth].push(root.borrow().val);
                build(root.borrow().left.clone(), results, depth + 1);
                build(root.borrow().right.clone(), results, depth + 1);
            }
        }
        build(root, &mut results, 0);

        results
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
        assert_eq!(
            Solution::level_order(input),
            vec![vec![3], vec![9, 20], vec![15, 7]]
        );

        let input = Some(Rc::new(RefCell::new(TreeNode {
            val: 1,
            left: None,
            right: None,
        })));
        assert_eq!(Solution::level_order(input), vec![vec![1]]);

        let empty: Vec<Vec<i32>> = vec![];
        assert_eq!(Solution::level_order(None), empty);
    }

    #[test]
    fn basic() {
        let input = Some(Rc::new(RefCell::new(TreeNode {
            val: 1,
            left: Some(Rc::new(RefCell::new(TreeNode {
                val: 2,
                left: Some(Rc::new(RefCell::new(TreeNode {
                    val: 3,
                    left: None,
                    right: None,
                }))),
                right: None,
            }))),
            right: Some(Rc::new(RefCell::new(TreeNode {
                val: 4,
                left: Some(Rc::new(RefCell::new(TreeNode {
                    val: 5,
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
        assert_eq!(
            Solution::level_order(input),
            vec![vec![1], vec![2, 4], vec![3, 5, 6]]
        );

        let input = Some(Rc::new(RefCell::new(TreeNode {
            val: 1,
            left: Some(Rc::new(RefCell::new(TreeNode {
                val: 2,
                left: None,
                right: Some(Rc::new(RefCell::new(TreeNode {
                    val: 3,
                    left: Some(Rc::new(RefCell::new(TreeNode {
                        val: 4,
                        left: Some(Rc::new(RefCell::new(TreeNode {
                            val: 5,
                            left: None,
                            right: Some(Rc::new(RefCell::new(TreeNode {
                                val: 6,
                                left: None,
                                right: None,
                            }))),
                        }))),
                        right: None,
                    }))),
                    right: None,
                }))),
            }))),
            right: None,
        })));
        assert_eq!(
            Solution::level_order(input),
            vec![vec![1], vec![2], vec![3], vec![4], vec![5], vec![6]]
        )
    }
}
