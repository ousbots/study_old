pub struct Solution;

impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mut open: Vec<char> = vec![];

        for letter in s.chars() {
            match letter {
                '(' => open.push(letter),
                '[' => open.push(letter),
                '{' => open.push(letter),

                ')' => {
                    if open.len() == 0 {
                        return false;
                    }
                    let last = open.pop().unwrap();
                    if last != '(' {
                        return false;
                    }
                }
                ']' => {
                    if open.len() == 0 {
                        return false;
                    }
                    let last = open.pop().unwrap();
                    if last != '[' {
                        return false;
                    }
                }
                '}' => {
                    if open.len() == 0 {
                        return false;
                    }
                    let last = open.pop().unwrap();
                    if last != '{' {
                        return false;
                    }
                }

                _ => return false,
            }
        }

        if open.len() == 0 {
            return true;
        }

        println!("length {}", open.len());
        false
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(Solution::is_valid("()".to_string()), true);
        assert_eq!(Solution::is_valid("()[]{}".to_string()), true);
        assert_eq!(Solution::is_valid("(]".to_string()), false);
    }

    #[test]
    fn basic() {
        assert_eq!(Solution::is_valid("(".to_string()), false);
        assert_eq!(Solution::is_valid("[".to_string()), false);
        assert_eq!(Solution::is_valid("{".to_string()), false);

        assert_eq!(Solution::is_valid("()".to_string()), true);
        assert_eq!(Solution::is_valid("[]".to_string()), true);
        assert_eq!(Solution::is_valid("{}".to_string()), true);

        assert_eq!(Solution::is_valid("([{}])[]{}".to_string()), true);

        assert_eq!(
            Solution::is_valid("(([]{[()]}){(([[(())]]))})".to_string()),
            true
        );

        assert_eq!(
            Solution::is_valid("(([]{[()]}){(([[(()])]]))})".to_string()),
            false
        );
    }
}
