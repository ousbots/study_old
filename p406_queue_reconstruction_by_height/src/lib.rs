use std::collections::hash_map::Entry::{Occupied, Vacant};
use std::collections::HashMap;

pub struct Solution;

impl Solution {
    pub fn reconstruct_queue_fast(people: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut people = people;
        let length = people.len();

        people.sort_by(|a, b| {
            if a[0] != b[0] {
                b[0].partial_cmp(&a[0]).unwrap()
            } else {
                a[1].partial_cmp(&b[1]).unwrap()
            }
        });

        println!("people {:?}", people);

        let mut queue = vec![];

        for person in people {
            queue.insert(person[1] as usize, person);
        }

        println!("queue {:?}", queue);

        queue
    }

    pub fn reconstruct_queue(people: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut behind: HashMap<i32, Vec<Vec<i32>>> = HashMap::new();

        for person in people {
            if person.len() != 2 {
                panic!("bad data {:?}", person);
            }

            let entry = match behind.entry(person[1]) {
                Occupied(entry) => entry.into_mut(),
                Vacant(entry) => entry.insert(vec![]),
            };

            entry.push(person);
        }

        for (_, people) in &mut behind {
            people.sort_by(|a, b| b[0].cmp(&a[0]));
        }

        let mut queue: Vec<Vec<i32>> = vec![];

        'outer: while !behind.is_empty() {
            for in_front in (0..=queue.len() as i32).rev() {
                if let Some(people) = behind.get_mut(&in_front) {
                    for index in (0..people.len()).rev() {
                        let person = &people[index];

                        let mut count = 0;
                        for elem in &queue {
                            if elem[0] >= person[0] {
                                count += 1;
                            }
                        }

                        if count == person[1] {
                            queue.push(person.clone());
                            people.remove(index);
                            if people.is_empty() {
                                behind.remove(&in_front);
                            }
                            continue 'outer;
                        }
                    }
                }
            }
        }

        queue
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn provided_tests() {
        assert_eq!(
            Solution::reconstruct_queue(vec![
                vec![7, 0],
                vec![4, 4],
                vec![7, 1],
                vec![5, 0],
                vec![6, 1],
                vec![5, 2]
            ]),
            vec![[5, 0], [7, 0], [5, 2], [6, 1], [4, 4], [7, 1]]
        );

        assert_eq!(
            Solution::reconstruct_queue_fast(vec![
                vec![7, 0],
                vec![4, 4],
                vec![7, 1],
                vec![5, 0],
                vec![6, 1],
                vec![5, 2]
            ]),
            vec![[5, 0], [7, 0], [5, 2], [6, 1], [4, 4], [7, 1]]
        );

        assert_eq!(
            Solution::reconstruct_queue(vec![
                vec![6, 0],
                vec![5, 0],
                vec![4, 0],
                vec![3, 2],
                vec![2, 2],
                vec![1, 4]
            ]),
            [[4, 0], [5, 0], [2, 2], [3, 2], [1, 4], [6, 0]]
        );

        assert_eq!(
            Solution::reconstruct_queue_fast(vec![
                vec![6, 0],
                vec![5, 0],
                vec![4, 0],
                vec![3, 2],
                vec![2, 2],
                vec![1, 4]
            ]),
            [[4, 0], [5, 0], [2, 2], [3, 2], [1, 4], [6, 0]]
        );
    }

    #[test]
    fn basic_tests() {
        assert_eq!(Solution::reconstruct_queue(vec![vec![0, 0]]), [[0, 0]]);
        assert_eq!(Solution::reconstruct_queue_fast(vec![vec![0, 0]]), [[0, 0]]);
    }
}
