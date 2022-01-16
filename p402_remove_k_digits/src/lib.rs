pub struct Solution;

// Remove the digit that generates the lowest number. This is always when nums[i] > nums[i+1].
fn remove_digit(nums: &mut Vec<u32>) {
    for index in 0..nums.len() - 1 {
        if nums[index] > nums[index + 1] {
            nums.remove(index);
            return;
        }
    }

    nums.pop();
}

impl Solution {
    // Return the lowest number from removing k digits from the given string.
    pub fn remove_kdigits(num: String, k: i32) -> String {
        let mut digits: Vec<u32> = num.chars().map(|x| x.to_digit(10).unwrap()).collect();

        if k >= digits.len() as i32 {
            return "0".to_string();
        }

        for _ in 0..k {
            remove_digit(&mut digits);
        }

        let mut start = 0;
        while digits[start] == 0 {
            if start == digits.len() - 1 {
                break;
            }

            start += 1;
        }

        digits[start..]
            .iter()
            .map(|num| num.to_string())
            .collect::<String>()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn provided_tests() {
        assert_eq!(Solution::remove_kdigits("1432219".to_string(), 3), "1219");
        assert_eq!(Solution::remove_kdigits("10200".to_string(), 1), "200");
        assert_eq!(Solution::remove_kdigits("10".to_string(), 2), "0");
    }

    #[test]
    fn more_test() {
        assert_eq!(Solution::remove_kdigits("0".to_string(), 1), "0");
        assert_eq!(Solution::remove_kdigits("99".to_string(), 2), "0");
        assert_eq!(Solution::remove_kdigits("".to_string(), 1), "0");
    }
}
