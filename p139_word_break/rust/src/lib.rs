use std::collections::HashMap;

pub struct Solution;
impl Solution {
    pub fn word_break(s: String, word_dict: Vec<String>) -> bool {
        let mut words = HashMap::<String, bool>::with_capacity(word_dict.len());
        for word in word_dict {
            words.insert(word, true);
        }

        let mut found = HashMap::<usize, bool>::with_capacity(s.len() + 1);
        found.insert(0, true);

        for i in 1..=s.len() {
            for j in 0..=i {
                if words.contains_key(&s[j..i]) && found.contains_key(&j) {
                    found.insert(i, true);
                }
            }
        }

        if found.contains_key(&s.len()) {
            return true;
        }

        false
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(
            Solution::word_break(
                String::from("leetcode"),
                ["leet", "code"].map(|e| e.to_string()).to_vec()
            ),
            true
        );

        assert_eq!(
            Solution::word_break(
                String::from("applepenapple"),
                ["apple", "pen"].map(|e| e.to_string()).to_vec()
            ),
            true
        );

        assert_eq!(
            Solution::word_break(
                String::from("catsandog"),
                ["cats", "dog", "sand", "and", "cat"]
                    .map(|e| e.to_string())
                    .to_vec()
            ),
            false
        );
    }

    #[test]
    fn basic() {
        assert_eq!(
            Solution::word_break(String::from("a"), ["b"].map(|e| e.to_string()).to_vec()),
            false
        );

        assert_eq!(
            Solution::word_break(String::from("a"), ["a"].map(|e| e.to_string()).to_vec()),
            true
        );

        assert_eq!(
            Solution::word_break(
                String::from("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaac"),
                ["a", "aa", "aaa"].map(|e| e.to_string()).to_vec()
            ),
            false
        );

        assert_eq!(
            Solution::word_break(
                String::from("abcdefg"),
                ["abcde", "defg", "a", "ab", "c"]
                    .map(|e| e.to_string())
                    .to_vec()
            ),
            true
        );

        assert_eq!(
            Solution::word_break(
                String::from("abcdefg"),
                ["abcde", "efg", "a", "ab", "c", "de", "abc", "cdef"]
                    .map(|e| e.to_string())
                    .to_vec()
            ),
            false
        );
    }
}
