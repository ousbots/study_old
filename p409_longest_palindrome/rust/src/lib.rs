use std::collections::hash_map::Entry::{Occupied, Vacant};
use std::collections::HashMap;

pub struct Solution;
impl Solution {
    pub fn longest_palindrome(s: String) -> i32 {
        let mut letters: HashMap<char, i32> = HashMap::new();
        let mut count = 0;

        s.chars().for_each(|x| {
            let entry = match letters.entry(x) {
                Occupied(e) => e.into_mut(),
                Vacant(e) => e.insert(0),
            };
            *entry += 1;
            if *entry == 2 {
                count += 2;
                *entry = 0;
            }
        });

        for (_, val) in letters {
            if val == 1 {
                count += 1;
                break;
            }
        }

        count
    }

    pub fn longest_palindrome_ascii(s: String) -> i32 {
        let mut letters = [0; 52];
        let mut count = 0;

        s.chars().for_each(|x| {
            let ch = if x.is_ascii_uppercase() {
                &mut letters[x as usize - 'A' as usize + 26]
            } else {
                &mut letters[x as usize - 'a' as usize]
            };

            *ch += 1;
            if *ch == 2 {
                count += 2;
                *ch = 0;
            }
        });

        for val in letters {
            if val != 0 {
                count += 1;
                break;
            }
        }

        count
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(Solution::longest_palindrome("abccccdd".to_string()), 7);
        assert_eq!(Solution::longest_palindrome("a".to_string()), 1);
        assert_eq!(Solution::longest_palindrome("bb".to_string()), 2);
    }

    #[test]
    fn basic() {
        assert_eq!(Solution::longest_palindrome("aA".to_string()), 1);
        assert_eq!(Solution::longest_palindrome("aabbccc".to_string()), 7);
        assert_eq!(Solution::longest_palindrome("abcabcAB".to_string()), 7);
    }
}
