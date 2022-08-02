use std::collections::HashMap;

pub struct Solution;
impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        let s: Vec<char> = s.chars().collect();
        let mut counts: HashMap<char, i32> = HashMap::new();
        let mut max = 0;
        let mut back = 0;

        for front in 0..s.len() {
            let count = counts.entry(s[front]).or_insert(0);
            *count += 1;

            if *count > 1 {
                while back < front {
                    let temp = counts.get_mut(&s[back]).unwrap();
                    *temp -= 1;
                    back += 1;

                    if *temp == 1 {
                        break;
                    }
                }
            } else {
                let length = front - back + 1;
                if length > max {
                    max = length;
                }
            }
        }

        i32::try_from(max).unwrap()
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    struct TestCase {
        input: String,
        output: i32,
    }

    #[test]
    fn basic() {
        let tests = vec![
            TestCase {
                input: String::from("a"),
                output: 1,
            },
            TestCase {
                input: String::from("aa"),
                output: 1,
            },
            TestCase {
                input: String::from("abab"),
                output: 2,
            },
            TestCase {
                input: String::from("ababcabcdabcabcbd"),
                output: 4,
            },
            TestCase {
                input: String::from("abbcdb"),
                output: 3,
            },
        ];

        for test in tests {
            assert_eq!(
                Solution::length_of_longest_substring(test.input),
                test.output
            );
        }
    }
}
