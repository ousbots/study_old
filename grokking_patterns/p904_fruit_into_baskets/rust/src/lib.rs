use std::collections::HashMap;

pub struct Solution;

impl Solution {
    // Finds the maximimum number of fruits that can be picked following a few rules. The elements in
    // the fruits vector are the type of the tree at that position. The rules for picking are:
    //   * Two baskets that can each hold only one type of fruit.
    //   * Starting at any tree, move to the right while picking exactly one fruit from each tree.
    //   * Stop if a tree is reached that cannot be picked due to the first rule.
    //
    // The solution uses a sliding window where the front of the window moves forward every cycle,
    // updating the count map. If the map has more than two elements in it, the rear of the window is
    // moved forward, updating the count map in reverse until there are only two elements left. The
    // max difference between the two window edges is maintained and returned.
    pub fn total_fruit(fruits: Vec<i32>) -> i32 {
        let mut counts: HashMap<i32, i32> = HashMap::new();
        let mut max: i32 = 0;

        let mut j = 0;
        for i in 0..fruits.len() {
            counts.entry(fruits[i]).and_modify(|e| *e += 1).or_insert(1);

            while counts.len() > 2 {
                counts.entry(fruits[j]).and_modify(|e| *e -= 1).or_insert(0);
                counts.retain(|_, v| *v != 0);
                j += 1;
            }

            let size = i32::try_from(i - j + 1).unwrap();
            if size > max {
                max = size;
            }
        }

        max
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    struct TestCase {
        fruits: Vec<i32>,
        expected: i32,
    }

    #[test]
    fn provided() {
        let cases = vec![
            TestCase {
                fruits: vec![1, 2, 1],
                expected: 3,
            },
            TestCase {
                fruits: vec![0, 1, 2, 2],
                expected: 3,
            },
            TestCase {
                fruits: vec![1, 2, 3, 2, 2],
                expected: 4,
            },
        ];

        for test in cases {
            assert_eq!(Solution::total_fruit(test.fruits), test.expected);
        }
    }
}
