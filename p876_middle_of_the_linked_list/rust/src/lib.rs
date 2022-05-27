// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

pub struct Solution;
impl Solution {
    pub fn middle_node(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut middle = head.clone();
        let mut end = head.clone();

        while let Some(node) = end {
            if node.next.is_none() {
                break;
            }

            middle = if let Some(node) = middle {
                node.next
            } else {
                None
            };

            end = if let Some(next) = node.next {
                next.next
            } else {
                None
            };
        }

        middle
    }

    pub fn middle_node_count(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut i = head.clone();
        let mut len = 0;
        while let Some(node) = i {
            len += 1;
            i = node.next;
        }

        len = len / 2;

        let mut middle = head.clone();
        while len != 0 {
            len -= 1;
            middle = middle.unwrap().next;
        }

        middle
    }
}

#[cfg(test)]
mod tests {
    use super::{ListNode, Solution};

    #[test]
    fn basic() {
        let input = Some(Box::new(ListNode { val: 1, next: None }));
        let valid = Some(Box::new(ListNode { val: 1, next: None }));
        assert_eq!(Solution::middle_node(input), valid);

        let input = Some(Box::new(ListNode {
            val: 1,
            next: Some(Box::new(ListNode { val: 2, next: None })),
        }));
        let valid = Some(Box::new(ListNode { val: 2, next: None }));
        assert_eq!(Solution::middle_node(input), valid);

        let input = Some(Box::new(ListNode {
            val: 1,
            next: Some(Box::new(ListNode {
                val: 2,
                next: Some(Box::new(ListNode { val: 3, next: None })),
            })),
        }));
        let valid = Some(Box::new(ListNode {
            val: 2,
            next: Some(Box::new(ListNode { val: 3, next: None })),
        }));
        assert_eq!(Solution::middle_node(input), valid);

        let input = Some(Box::new(ListNode {
            val: 1,
            next: Some(Box::new(ListNode {
                val: 2,
                next: Some(Box::new(ListNode {
                    val: 3,
                    next: Some(Box::new(ListNode { val: 4, next: None })),
                })),
            })),
        }));
        let valid = Some(Box::new(ListNode {
            val: 3,
            next: Some(Box::new(ListNode { val: 4, next: None })),
        }));
        assert_eq!(Solution::middle_node(input), valid);
    }
}
