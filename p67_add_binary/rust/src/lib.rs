pub struct Solution;
impl Solution {
    pub fn add_binary(a: String, b: String) -> String {
        let a: Vec<char> = a.chars().collect();
        let b: Vec<char> = b.chars().collect();
        let mut sum = vec![];

        let mut carry = false;
        let mut i = a.len() as i64;
        let mut j = b.len() as i64;

        while i > 0 || j > 0 {
            i -= 1;
            j -= 1;

            let x = if i >= 0 { a[i as usize] } else { '0' };
            let y = if j >= 0 { b[j as usize] } else { '0' };

            if x == '1' && y == '1' {
                if carry {
                    sum.push('1');
                } else {
                    sum.push('0');
                    carry = true;
                }
            } else if x == '1' || y == '1' {
                if carry {
                    sum.push('0');
                } else {
                    sum.push('1');
                }
            } else {
                if carry {
                    sum.push('1');
                    carry = false;
                } else {
                    sum.push('0');
                }
            }
        }

        if carry {
            sum.push('1')
        }

        sum.reverse();

        sum.iter().cloned().collect::<String>()
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(
            Solution::add_binary("11".to_string(), "1".to_string()),
            "100"
        );
        assert_eq!(
            Solution::add_binary("1010".to_string(), "1011".to_string()),
            "10101"
        );
    }

    #[test]
    fn basic() {
        assert_eq!(Solution::add_binary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101".to_string(), "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011".to_string()), "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000")
    }
}
