use std::collections::hash_map::Entry::{Occupied, Vacant};
use std::collections::HashMap;

pub struct Solution;

// Definition for singly-linked list.j
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }

    fn from_vec(input: Vec<i32>) -> Option<Box<ListNode>> {
        let mut root: Option<Box<ListNode>> = None;
        let mut current = &mut root;

        for num in input {
            match current {
                Some(ref mut node) => {
                    node.next = Some(Box::new(ListNode::new(num)));
                    current = &mut node.next;
                }
                None => {
                    *current = Some(Box::new(ListNode::new(num)));
                }
            }
        }

        root
    }
}

// Convenience function to convert a vector of vectors into a vector of linked lists for testing.
pub fn convert_input(input: Vec<Vec<i32>>) -> Vec<Option<Box<ListNode>>> {
    let mut converted = vec![];

    for elem in input {
        converted.push(ListNode::from_vec(elem));
    }

    converted
}

// Convenience function to convert a linked list into a vector for testing.
pub fn convert_output(output: Option<Box<ListNode>>) -> Vec<i32> {
    let mut converted = vec![];

    let mut current = output;
    while let Some(node) = current {
        converted.push(node.val);
        current = node.next;
    }

    converted
}

impl Solution {
    pub fn merge_k_lists(lists: Vec<Option<Box<ListNode>>>) -> Option<Box<ListNode>> {
        let mut count: HashMap<i32, i32> = HashMap::new();

        for list in lists {
            let mut current = &list;
            while let Some(node) = current {
                let entry = match count.entry(node.val) {
                    Occupied(entry) => entry.into_mut(),
                    Vacant(entry) => entry.insert(0),
                };
                *entry += 1;

                current = &node.next;
            }
        }

        let mut keys: Vec<i32> = count.keys().cloned().collect();
        keys.sort();

        let mut merged: Option<Box<ListNode>> = None;
        let mut current = &mut merged;

        for key in keys {
            for _ in 0..*count.get(&key).unwrap() {
                match current {
                    Some(ref mut node) => {
                        node.next = Some(Box::new(ListNode::new(key)));
                        current = &mut node.next;
                    }
                    None => *current = Some(Box::new(ListNode::new(key))),
                }
            }
        }

        merged
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn provided_tests() {
        assert_eq!(
            convert_output(Solution::merge_k_lists(convert_input(vec![
                vec![1, 4, 5],
                vec![1, 3, 4],
                vec![2, 6]
            ]))),
            [1, 1, 2, 3, 4, 4, 5, 6]
        );
        assert_eq!(convert_output(Solution::merge_k_lists(vec![])), []);
        assert_eq!(convert_output(Solution::merge_k_lists(vec![None])), []);
    }

    #[test]
    fn basic_tests() {
        assert_eq!(
            convert_output(Solution::merge_k_lists(convert_input(vec![
                vec![1, 1, 4, 5, 9],
                vec![1, 3, 4],
                vec![2, 10],
                vec![11, 13],
                vec![0],
            ]))),
            [0, 1, 1, 1, 2, 3, 4, 4, 5, 9, 10, 11, 13]
        );
    }
}
