// Definition for singly-linked list.
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
}

pub struct Solution;

impl Solution {
    pub fn rotate_right(head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
        let mut head = match head {
            Some(node) => node,
            None => return None,
        };

        if head.next.is_none() {
            return Some(head);
        }

        let rotation = {
            let mut count = 1;
            let mut node = &head;

            while let Some(ref next) = node.next {
                count += 1;
                node = next;
            }

            if k == 0 || k == count || (k % count) == 0 {
                return Some(head);
            }

            count - (k % count)
        };

        let mut temp = &mut head;
        for _ in 1..rotation {
            if let Some(ref mut next) = temp.next {
                temp = next;
            }
        }

        let mut new_head = temp.next.clone().unwrap();
        temp.next = None;

        let mut tail = {
            let mut temp = &mut new_head;
            while let Some(ref mut next) = temp.next {
                temp = next;
            }
            temp
        };

        tail.next = Some(head);

        Some(new_head)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn provided_tests() {
        assert_eq!(
            Solution::rotate_right(
                Some(Box::new(ListNode {
                    val: 1,
                    next: Some(Box::new(ListNode {
                        val: 2,
                        next: Some(Box::new(ListNode {
                            val: 3,
                            next: Some(Box::new(ListNode {
                                val: 4,
                                next: Some(Box::new(ListNode { val: 5, next: None })),
                            }))
                        }))
                    }))
                })),
                2
            ),
            Some(Box::new(ListNode {
                val: 4,
                next: Some(Box::new(ListNode {
                    val: 5,
                    next: Some(Box::new(ListNode {
                        val: 1,
                        next: Some(Box::new(ListNode {
                            val: 2,
                            next: Some(Box::new(ListNode { val: 3, next: None })),
                        }))
                    }))
                }))
            }))
        );

        assert_eq!(
            Solution::rotate_right(
                Some(Box::new(ListNode {
                    val: 0,
                    next: Some(Box::new(ListNode {
                        val: 1,
                        next: Some(Box::new(ListNode { val: 2, next: None })),
                    }))
                })),
                4
            ),
            Some(Box::new(ListNode {
                val: 2,
                next: Some(Box::new(ListNode {
                    val: 0,
                    next: Some(Box::new(ListNode { val: 1, next: None })),
                }))
            }))
        );
    }

    #[test]
    fn more_tests() {
        assert_eq!(
            Solution::rotate_right(
                Some(Box::new(ListNode {
                    val: 1,
                    next: Some(Box::new(ListNode { val: 2, next: None }))
                })),
                0
            ),
            Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode { val: 2, next: None }))
            }))
        );

        assert_eq!(
            Solution::rotate_right(
                Some(Box::new(ListNode {
                    val: 1,
                    next: Some(Box::new(ListNode { val: 2, next: None }))
                })),
                2
            ),
            Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode { val: 2, next: None }))
            }))
        );
    }
}
