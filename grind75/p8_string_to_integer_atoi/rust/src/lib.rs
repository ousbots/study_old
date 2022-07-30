pub struct Solution;
impl Solution {
    pub fn my_atoi(s: String) -> i32 {
        let s: Vec<char> = s.chars().collect();

        let mut index = 0;
        for i in 0..s.len() {
            if s[i] != ' ' {
                break;
            }
            index += 1;
        }

        let mut factor = 1;
        if index < s.len() {
            if s[index] == '+' {
                index += 1;
            } else if s[index] == '-' {
                factor = -1;
                index += 1;
            }
        }

        let max = i64::from(i32::MAX);
        let min = i64::from(i32::MIN);

        let mut value = i64::from(0);
        for elem in s[index..s.len()].iter() {
            let digit = match elem {
                '0' => 0,
                '1' => 1,
                '2' => 2,
                '3' => 3,
                '4' => 4,
                '5' => 5,
                '6' => 6,
                '7' => 7,
                '8' => 8,
                '9' => 9,
                _ => break,
            };

            value = (value * 10) + digit;
            if value > max {
                break;
            }
        }

        value = value * factor;

        if value > max {
            value = max;
        } else if value < min {
            value = min;
        }

        value as i32
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn basic() {
        assert_eq!(Solution::my_atoi(String::from("")), 0);
        assert_eq!(Solution::my_atoi(String::from("1")), 1);
        assert_eq!(Solution::my_atoi(String::from("  1")), 1);
        assert_eq!(Solution::my_atoi(String::from("-1")), -1);
        assert_eq!(Solution::my_atoi(String::from(" -1")), -1);
        assert_eq!(Solution::my_atoi(String::from("+1")), 1);
        assert_eq!(Solution::my_atoi(String::from(" +1")), 1);
        assert_eq!(Solution::my_atoi(String::from(" +-1")), 0);
        assert_eq!(Solution::my_atoi(String::from(" +1 1")), 1);
        assert_eq!(Solution::my_atoi(String::from(" +1 a")), 1);
        assert_eq!(Solution::my_atoi(String::from(" +1a")), 1);
        assert_eq!(Solution::my_atoi(String::from("  a1")), 0);
        assert_eq!(Solution::my_atoi(String::from("99999")), 99999);
        assert_eq!(Solution::my_atoi(String::from("-2147483649")), -2147483648);
        assert_eq!(Solution::my_atoi(String::from("2147483648")), 2147483647);
        assert_eq!(
            Solution::my_atoi(String::from("9223372036800400")),
            2147483647
        );
    }
}
