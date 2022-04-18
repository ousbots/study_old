use std::collections::hash_map::Entry::{Occupied, Vacant};
use std::collections::HashMap;

pub struct Solution;

impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        if s.len() != t.len() {
            return false;
        }

        let s: Vec<char> = s.chars().collect();
        let t: Vec<char> = t.chars().collect();

        let mut count: HashMap<char, i32> = HashMap::new();

        for index in 0..s.len() {
            let add = match count.entry(s[index]) {
                Occupied(entry) => entry.into_mut(),
                Vacant(entry) => entry.insert(0),
            };
            *add += 1;

            let sub = match count.entry(t[index]) {
                Occupied(entry) => entry.into_mut(),
                Vacant(entry) => entry.insert(0),
            };
            *sub -= 1;
        }

        for (_, val) in count {
            if val != 0 {
                return false;
            }
        }

        true
    }

    // This only works for ascii alpha strings with a 26 character alphabet and isn't any faster
    // than the hashmap solution on leetcode.
    pub fn is_anagram_ascii(s: String, t: String) -> bool {
        if s.len() != t.len() {
            return false;
        }

        let s: Vec<char> = s.chars().collect();
        let t: Vec<char> = t.chars().collect();

        let mut count: [i32; 26] = [0; 26];

        fn translate(letter: char) -> usize {
            match letter {
                'a' => 0,
                'b' => 1,
                'c' => 2,
                'd' => 3,
                'e' => 4,
                'f' => 5,
                'g' => 6,
                'h' => 7,
                'i' => 8,
                'j' => 9,
                'k' => 10,
                'l' => 11,
                'm' => 12,
                'n' => 13,
                'o' => 14,
                'p' => 15,
                'q' => 16,
                'r' => 17,
                's' => 18,
                't' => 19,
                'u' => 20,
                'v' => 21,
                'w' => 22,
                'x' => 23,
                'y' => 24,
                'z' => 25,
                _ => 9999,
            }
        }

        for index in 0..s.len() {
            count[translate(s[index])] += 1;
            count[translate(t[index])] -= 1;
        }

        for val in count {
            if val != 0 {
                return false;
            }
        }

        true
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(
            Solution::is_anagram(String::from("anagram"), String::from("nagaram")),
            true
        );
        assert_eq!(
            Solution::is_anagram(String::from("rat"), String::from("car")),
            false
        );
    }

    #[test]
    fn basic() {
        assert_eq!(
            Solution::is_anagram(String::from("a"), String::from("a")),
            true
        );
        assert_eq!(
            Solution::is_anagram(String::from("a"), String::from("b")),
            false
        );
        assert_eq!(
            Solution::is_anagram(String::from("aaaa"), String::from("aaab")),
            false
        );
        assert_eq!(
            Solution::is_anagram(String::from("aabaa"), String::from("aaaab")),
            true
        );
        assert_eq!(
            Solution::is_anagram(String::from("aabaa"), String::from("aabaa")),
            true
        );
        assert_eq!(
            Solution::is_anagram(String::from("aabaaa"), String::from("aabaab")),
            false
        );
    }
}
