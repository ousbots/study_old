#[derive(Clone, Copy)]
pub struct Solution {
    bad_version: i32,
}

impl Solution {
    #[allow(non_snake_case)]
    fn isBadVersion(self, version: i32) -> bool {
        if version >= self.bad_version {
            return true;
        } else {
            return false;
        }
    }

    pub fn first_bad_version(&self, n: i32) -> i32 {
        let mut low = 0;
        let mut high = n;
        let mut good = 0;

        while low <= high {
            let middle = (high - low) / 2 + low;

            if self.isBadVersion(middle) {
                if middle - 1 == good {
                    return middle;
                }

                high = middle;
            } else {
                good = middle;
                low = middle + 1;
            }
        }

        0
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        let mut test = Solution { bad_version: 0 };

        test.bad_version = 4;
        assert_eq!(test.first_bad_version(5), 4);

        test.bad_version = 1;
        assert_eq!(test.first_bad_version(1), 1);
    }

    #[test]
    fn basic() {
        let mut test = Solution { bad_version: 0 };

        test.bad_version = 1;
        assert_eq!(test.first_bad_version(99), 1);

        test.bad_version = 99;
        assert_eq!(test.first_bad_version(99), 99);

        test.bad_version = 421;
        assert_eq!(test.first_bad_version(999), 421);

        test.bad_version = 1702766719;
        assert_eq!(test.first_bad_version(2126753390), 1702766719);
    }
}
