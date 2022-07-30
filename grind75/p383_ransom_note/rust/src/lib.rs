use std::collections::hash_map::Entry::{Occupied, Vacant};
use std::collections::HashMap;

pub struct Solution;

impl Solution {
    pub fn can_construct_utf8(ransom_note: String, magazine: String) -> bool {
        let mut count: HashMap<char, i32> = HashMap::new();

        for elem in ransom_note.chars() {
            let entry = match count.entry(elem) {
                Occupied(a) => a.into_mut(),
                Vacant(a) => a.insert(0),
            };

            *entry -= 1;
        }

        for elem in magazine.chars() {
            let entry = match count.entry(elem) {
                Occupied(a) => a.into_mut(),
                Vacant(a) => a.insert(0),
            };

            *entry += 1;
        }

        for (_, val) in count {
            if val < 0 {
                return false;
            }
        }

        true
    }

    // Works well for ascii values.
    pub fn can_construct_ascii(ransom_note: String, magazine: String) -> bool {
        let mut count = [0; 26];

        for elem in ransom_note.chars() {
            count[elem as usize - 'a' as usize] -= 1;
        }

        for elem in magazine.chars() {
            count[elem as usize - 'a' as usize] += 1;
        }

        for value in count {
            if value < 0 {
                return false;
            }
        }

        true
    }

    // Using the match isn't any faster than casting the chars to their ascii value.
    pub fn can_construct(ransom_note: String, magazine: String) -> bool {
        let mut count = [0; 26];

        for elem in ransom_note.chars() {
            count[to_index(elem)] -= 1;
        }

        for elem in magazine.chars() {
            count[to_index(elem)] += 1;
        }

        for value in count {
            if value < 0 {
                return false;
            }
        }

        true
    }
}

fn to_index(letter: char) -> usize {
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
        _ => 999,
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(
            Solution::can_construct("a".to_string(), "b".to_string()),
            false
        );
        assert_eq!(
            Solution::can_construct("aa".to_string(), "ab".to_string()),
            false
        );
        assert_eq!(
            Solution::can_construct("aa".to_string(), "aab".to_string()),
            true
        );
    }

    #[test]
    fn basic() {
        assert_eq!(
            Solution::can_construct(
                "alloftheransomscanbetracedtoz".to_string(),
                "alloftheransomsanbetracedtoz".to_string()
            ),
            false
        );

        assert_eq!(
            Solution::can_construct(
                "alloftheransomscanbetracedtoz".to_string(),
                "alloftheransomscanbetracedtoz".to_string()
            ),
            true
        );

        assert_eq!(
            Solution::can_construct(
                "alloftheransomscanbetracedtoz".to_string(),
                "alloftheransomscanbetracedtozhomie".to_string()
            ),
            true
        );
    }
}
