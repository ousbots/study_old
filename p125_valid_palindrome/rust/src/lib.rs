pub struct Solution;

impl Solution {
    pub fn is_palindrome_slower(s: String) -> bool {
        let clean = s
            .to_lowercase()
            .chars()
            .filter(|c| c.is_ascii_alphanumeric())
            .collect::<Vec<char>>();

        if clean.len() <= 1 {
            return true;
        }

        let mut left = 0;
        let mut right = clean.len() - 1;

        while left <= right {
            if clean[left] != clean[right] {
                return false;
            }

            left += 1;
            right -= 1;
        }

        true
    }

    pub fn is_palindrome(s: String) -> bool {
        if s.len() <= 1 {
            return true;
        }

        let letters: Vec<char> = s.chars().collect();
        let mut left = 0;
        let mut right = letters.len() - 1;

        while left <= right {
            while left < letters.len() - 1 && !letters[left].is_alphanumeric() {
                left += 1;
            }

            while right > 0 && !letters[right].is_alphanumeric() {
                right -= 1;
            }

            if left == letters.len() || right == 0 {
                break;
            }

            if letters[left].to_ascii_lowercase() != letters[right].to_ascii_lowercase() {
                return false;
            }

            left += 1;
            right -= 1;
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
            Solution::is_palindrome("A man, a plan, a canal: Panama".to_string()),
            true
        );
        assert_eq!(Solution::is_palindrome("race a car".to_string()), false);
        assert_eq!(Solution::is_palindrome(" ".to_string()), true);
    }

    #[test]
    fn basic() {
        assert_eq!(Solution::is_palindrome("abcdedcba".to_string()), true);
        assert_eq!(Solution::is_palindrome("abcdeedcba".to_string()), true);
        assert_eq!(Solution::is_palindrome("abcdefdcba".to_string()), false);
        assert_eq!(Solution::is_palindrome("a.".to_string()), true);
        assert_eq!(Solution::is_palindrome(".,".to_string()), true);
        assert_eq!(Solution::is_palindrome("a.,".to_string()), true);
    }
}
