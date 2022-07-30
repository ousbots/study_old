pub struct Solution;

impl Solution {
    // Create an array of n elements where output[i] = number of 1 digits in binary representation
    // of i.
    // Note:
    //   0 (0000) -> 0
    //   1 (0001) -> 1
    //   2 (0010) -> 1
    //   3 (0011) -> 2
    //   4 (0100) -> 1
    //   5 (0101) -> 2
    //   6 (0110) -> 2
    //   7 (0111) -> 3
    //   8 (1000) -> 1
    //   9 (1001) -> 2
    //  10 (1010) -> 2
    //  11 (1011) -> 3
    //  12 (1100) -> 2
    //  13 (1101) -> 3
    //  14 (1110) -> 3
    //  15 (1111) -> 4
    //
    // Which shows that the number of bits for i = i & (i -1) + 1. When the number "overflows" to
    // the next order of 2 it resets back to the "base" + 1 and in other cases it's the previous
    // value + 1.
    pub fn count_bits(n: i32) -> Vec<i32> {
        let mut count = vec![0];

        for index in 1..=n as usize {
            count.push(&count[index & (index - 1)] + 1);
        }

        count
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn provided_tests() {
        assert_eq!(Solution::count_bits(2), [0, 1, 1]);
        assert_eq!(Solution::count_bits(5), [0, 1, 1, 2, 1, 2]);
    }

    #[test]
    fn more_tests() {
        assert_eq!(Solution::count_bits(0), [0]);
        assert_eq!(Solution::count_bits(10), [0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2]);
    }
}
