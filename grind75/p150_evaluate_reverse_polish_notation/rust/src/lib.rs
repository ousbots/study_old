pub struct Solution;
impl Solution {
    // Using .pop().unwrap() and .push() were much slower than direct indexing. Pre-allocation
    // is required, but didn't help much when using pop() and push().
    pub fn eval_rpn(tokens: Vec<String>) -> i32 {
        let mut stack: Vec<i32> = vec![0; tokens.len()];
        let mut index = 0;

        for token in tokens {
            match token.as_str() {
                "+" => {
                    index -= 1;
                    let value = stack[index - 1] + stack[index];
                    stack[index - 1] = value;
                }

                "-" => {
                    index -= 1;
                    let value = stack[index - 1] - stack[index];
                    stack[index - 1] = value;
                }
                "*" => {
                    index -= 1;
                    let value = stack[index - 1] * stack[index];
                    stack[index - 1] = value;
                }
                "/" => {
                    index -= 1;
                    let value = stack[index - 1] / stack[index];
                    stack[index - 1] = value;
                }
                _ => {
                    stack[index] = token.parse().unwrap();
                    index += 1;
                }
            }
        }

        stack[0]
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(
            Solution::eval_rpn(vec![
                "2".to_string(),
                "1".to_string(),
                "+".to_string(),
                "3".to_string(),
                "*".to_string()
            ]),
            9
        );

        assert_eq!(
            Solution::eval_rpn(vec![
                "4".to_string(),
                "13".to_string(),
                "5".to_string(),
                "/".to_string(),
                "+".to_string()
            ]),
            6
        );

        assert_eq!(
            Solution::eval_rpn(vec![
                "10".to_string(),
                "6".to_string(),
                "9".to_string(),
                "3".to_string(),
                "+".to_string(),
                "-11".to_string(),
                "*".to_string(),
                "/".to_string(),
                "*".to_string(),
                "17".to_string(),
                "+".to_string(),
                "5".to_string(),
                "+".to_string(),
            ]),
            22
        );
    }
}
