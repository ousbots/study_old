pub struct Solution;
impl Solution {
    pub fn sort_colors(nums: &mut Vec<i32>) {
        let mut white = 0;
        let mut index = 0;

        for i in 0..nums.len() {
            match nums[i] {
                0 => {
                    nums[index] = 0;
                    index += 1;
                }
                1 => white += 1,
                _ => continue,
            }
        }

        for i in index..index + white {
            nums[i] = 1;
        }

        for i in index + white..nums.len() {
            nums[i] = 2;
        }
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        let mut input = vec![2, 0, 2, 1, 1, 0];
        let valid = vec![0, 0, 1, 1, 2, 2];
        Solution::sort_colors(&mut input);
        assert_eq!(input, valid);

        let mut input = vec![2, 0, 1];
        let valid = vec![0, 1, 2];
        Solution::sort_colors(&mut input);
        assert_eq!(input, valid);
    }
}
