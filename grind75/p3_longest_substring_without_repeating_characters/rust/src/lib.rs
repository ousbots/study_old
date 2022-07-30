use std::collections::hash_map::Entry::{Occupied, Vacant};
use std::collections::HashMap;

pub struct Solution;

impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        let letters: Vec<char> = s.chars().collect();
        let mut count: HashMap<char, i32> = HashMap::new();
        let mut back = 0;
        let mut max = 0;

        for (index, letter) in letters.iter().enumerate() {
            let entry = match count.entry(*letter) {
                Occupied(elem) => elem.into_mut(),
                Vacant(elem) => elem.insert(0),
            };

            *entry += 1;

            if *entry > 1 {
                while back <= index {
                    let tmp = match count.entry(letters[back]) {
                        Occupied(elem) => elem.into_mut(),
                        Vacant(_) => panic!("vacant element while moving back"),
                    };
                    back += 1;

                    *tmp -= 1;
                    if *tmp == 1 {
                        break;
                    }
                }
            } else {
                let length = index - back + 1;
                if length > max {
                    max = length;
                }
            }
        }

        max as i32
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(
            Solution::length_of_longest_substring("abcabcbb".to_string()),
            3
        );
        assert_eq!(
            Solution::length_of_longest_substring("bbbbb".to_string()),
            1
        );
        assert_eq!(
            Solution::length_of_longest_substring("pwwkew".to_string()),
            3
        );
    }

    #[test]
    fn basic() {
        assert_eq!(Solution::length_of_longest_substring(" ".to_string()), 1);
        assert_eq!(
            Solution::length_of_longest_substring("abcdefghijklmnopqrstuvwxyzabc".to_string()),
            26
        );
        assert_eq!(
            Solution::length_of_longest_substring("abcdefghijklmnoapqrstubvwxyz".to_string()),
            26
        );
    }
}
