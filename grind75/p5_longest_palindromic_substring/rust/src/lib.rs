pub struct Solution;
impl Solution {
    pub fn longest_palindrome(s: String) -> String {
        let s: Vec<char> = s.chars().collect();
        let mut start = 0;
        let mut end = 0;

        for i in 0..s.len() {
            let length = Solution::find(&s, i, i).max(Solution::find(&s, i, i + 1));

            if length > end - start {
                end = i + (length / 2) + 1;

                if length % 2 == 0 {
                    start = i - ((length / 2) - 1);
                } else {
                    start = i - (length / 2);
                }
            }
        }

        s[start..end].iter().collect()
    }

    fn find(s: &Vec<char>, start: usize, end: usize) -> usize {
        let mut length = 0;
        let mut start = start;
        let mut end = end;

        while end < s.len() && s[start] == s[end] {
            length = end - start + 1;

            if start == 0 {
                break;
            }
            start -= 1;
            end += 1;
        }

        return length;
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn basic() {
        assert_eq!(Solution::longest_palindrome(String::from("a")), "a");
        assert_eq!(Solution::longest_palindrome(String::from("aab")), "aa");
        assert_eq!(Solution::longest_palindrome(String::from("aaba")), "aba");
        assert_eq!(
            Solution::longest_palindrome(String::from("aabaaaba")),
            "abaaaba"
        );
    }
}
