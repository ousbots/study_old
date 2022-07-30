use std::cmp::Reverse;
use std::collections::BinaryHeap;

pub struct Solution;

// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    // Create a linked list from a vector.
    pub fn new(values: Vec<i32>) -> Option<Box<ListNode>> {
        if values.len() == 0 {
            return None;
        }

        let mut root = None;

        for index in (0..values.len()).rev() {
            let temp = ListNode {
                val: values[index],
                next: root,
            };
            root = Some(Box::new(temp));
        }

        root
    }

    // Convert the linked list to a vector.
    pub fn to_vec(self) -> Vec<i32> {
        let mut results = vec![];
        let mut current = Some(Box::new(self));

        while let Some(node) = current {
            results.push(node.val);
            current = node.next;
        }

        results
    }
}

// Calculate the length of the linked list.
pub fn len(head: Option<Box<ListNode>>) -> i32 {
    let mut len = 0;

    match head {
        None => return len,
        Some(node) => {
            let mut current = *node;
            loop {
                len += 1;

                let temp = current;
                match temp.next {
                    Some(node) => current = *node,
                    None => break,
                }
            }
        }
    }

    len
}

// Split the linked list at the given position.
pub fn split(head: Option<Box<ListNode>>) -> (Option<Box<ListNode>>, Option<Box<ListNode>>) {
    match head {
        Some(root) => {
            let mut fast = Some(root.clone());
            let mut base = root.clone();
            let mut slow = &mut base;

            while fast != None {
                if let Some(node) = fast.clone() {
                    fast = node.next;

                    if let Some(node) = fast.clone() {
                        if node.next != None {
                            fast = node.next;

                            if let Some(ref mut next) = slow.next {
                                slow = next;
                            }
                        }
                    }
                }
            }

            let second = slow.next.clone();
            slow.next = None;

            (Some(base), second)
        }

        None => (None, None),
    }
}

// Merge sort for linked lists.
pub fn merge_sort(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    match head {
        Some(root) => {
            if root.next == None {
                return Some(root);
            }

            let (first, second) = split(Some(root.clone()));
            let mut first = merge_sort(first);
            let mut second = merge_sort(second);

            let mut base = ListNode { val: 0, next: None };
            let mut current = &mut base;

            while first != None || second != None {
                let first_val = match first.clone() {
                    Some(node) => node.val,
                    None => i32::MAX,
                };

                let second_val = match second.clone() {
                    Some(node) => node.val,
                    None => i32::MAX,
                };

                if first_val < second_val {
                    if let Some(node) = first {
                        first = node.next;
                    }
                    current.next = Some(Box::new(ListNode {
                        val: first_val,
                        next: None,
                    }));
                } else {
                    if let Some(node) = second {
                        second = node.next;
                    }
                    current.next = Some(Box::new(ListNode {
                        val: second_val,
                        next: None,
                    }));
                }

                if let Some(ref mut next) = current.next {
                    current = next;
                }
            }

            base.next
        }
        None => None,
    }
}

impl Solution {
    pub fn sort_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        merge_sort(head)
    }

    // Alternate implementation that doesn't use merge sort and is much faster.
    pub fn sort_list2(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let current = &mut head;
        let mut heap = BinaryHeap::new();

        while let Some(ref mut node) = current {
            heap.push(Reverse(node.val));
            *current = node.next.take();
        }

        let mut root: Option<Box<ListNode>> = None;
        let mut current: &mut Option<Box<ListNode>> = &mut None;

        while !heap.is_empty() {
            let temp = Some(Box::new(ListNode {
                val: heap.pop().unwrap().0,
                next: None,
            }));

            match current {
                Some(node) => {
                    node.next = temp;
                    current = &mut node.next;
                }
                None => {
                    root = temp;
                    current = &mut root;
                }
            }
        }

        root
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn provided_tests() {
        assert_eq!(
            Solution::sort_list(ListNode::new(vec![4, 2, 1, 3]))
                .unwrap()
                .to_vec(),
            vec![1, 2, 3, 4]
        );

        assert_eq!(
            Solution::sort_list(ListNode::new(vec![-1, 5, 3, 4, 0]))
                .unwrap()
                .to_vec(),
            vec![-1, 0, 3, 4, 5]
        );

        assert_eq!(
            Solution::sort_list2(ListNode::new(vec![4, 2, 1, 3]))
                .unwrap()
                .to_vec(),
            vec![1, 2, 3, 4]
        );

        assert_eq!(
            Solution::sort_list2(ListNode::new(vec![-1, 5, 3, 4, 0]))
                .unwrap()
                .to_vec(),
            vec![-1, 0, 3, 4, 5]
        );
    }
}
