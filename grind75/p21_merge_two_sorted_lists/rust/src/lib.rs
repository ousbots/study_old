use std::mem;

// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

pub struct Solution;

impl Solution {
    pub fn merge_two_lists(
        mut list1: Option<Box<ListNode>>,
        mut list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        //let mut head = ListNode { val: 0, next: None };
        let mut head = None;
        let mut current = &mut head;

        let l1 = &mut list1;
        let l2 = &mut list2;

        while l1.is_some() && l2.is_some() {
            if l1.as_ref().unwrap().val < l2.as_ref().unwrap().val {
                mem::swap(current, l1);
                mem::swap(l1, &mut current.as_mut().unwrap().next);
            } else {
                mem::swap(current, l2);
                mem::swap(l2, &mut current.as_mut().unwrap().next);
            }

            current = &mut current.as_mut().unwrap().next;
        }

        mem::swap(
            current,
            if list1.is_none() {
                &mut list2
            } else {
                &mut list1
            },
        );

        head
    }
}

pub fn compare(list1: Option<Box<ListNode>>, list2: Option<Box<ListNode>>) -> bool {
    let mut l1 = list1;
    let mut l2 = list2;

    while l1.is_some() && l2.is_some() {
        if l1.as_ref().unwrap().val != l2.as_ref().unwrap().val {
            return false;
        }

        l1 = l1.unwrap().next;
        l2 = l2.unwrap().next;
    }

    if l1.is_none() && l2.is_none() {
        return true;
    }

    false
}

#[cfg(test)]
mod tests {
    use super::{compare, ListNode, Solution};

    #[test]
    fn provided() {
        let list1 = Some(Box::new(ListNode {
            val: 1,
            next: Some(Box::new(ListNode {
                val: 2,
                next: Some(Box::new(ListNode { val: 4, next: None })),
            })),
        }));

        let list2 = Some(Box::new(ListNode {
            val: 1,
            next: Some(Box::new(ListNode {
                val: 3,
                next: Some(Box::new(ListNode { val: 4, next: None })),
            })),
        }));

        let ans = Some(Box::new(ListNode {
            val: 1,
            next: Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode {
                    val: 2,
                    next: Some(Box::new(ListNode {
                        val: 3,
                        next: Some(Box::new(ListNode {
                            val: 4,
                            next: Some(Box::new(ListNode { val: 4, next: None })),
                        })),
                    })),
                })),
            })),
        }));

        let result = Solution::merge_two_lists(list1, list2);
        assert_eq!(compare(result, ans), true);

        let list1 = None;
        let list2 = None;
        let ans = None;
        let result = Solution::merge_two_lists(list1, list2);
        assert_eq!(compare(result, ans), true);

        let list1 = None;
        let list2 = Some(Box::new(ListNode { val: 0, next: None }));
        let ans = Some(Box::new(ListNode { val: 0, next: None }));
        let result = Solution::merge_two_lists(list1, list2);
        assert_eq!(compare(result, ans), true);
    }
}
