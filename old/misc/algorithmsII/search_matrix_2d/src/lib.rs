pub struct Solution;

impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        for row in matrix {
            if row.len() > 0 && row[0] <= target && row[row.len() - 1] >= target {
                return binary_search_exists(&row, target);
            }
        }

        false
    }
}

fn binary_search_exists(data: &[i32], target: i32) -> bool {
    let mut start = 0;
    let mut end = data.len() - 1;

    while end >= start {
        if data[start] > target || data[end] < target {
            return false;
        }

        let half = (start + end) / 2;

        if data[half] == target {
            return true;
        }

        if data[half] > target {
            // This check is used to avoid a usize overflow.
            if half == 0 {
                return false;
            }
            end = half - 1;
        } else {
            start = half + 1;
        }
    }

    false
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::search_matrix(
                vec![vec![1, 3, 5, 7], vec![10, 11, 16, 20], vec![23, 30, 34, 60]],
                3
            ),
            true
        );
        assert_eq!(
            Solution::search_matrix(vec![vec![1, 1, 1], vec![1, 1]], 1),
            true
        );
        assert_eq!(Solution::search_matrix(vec![vec![1]], 1), true);
        assert_eq!(Solution::search_matrix(vec![], 1), false);
        assert_eq!(Solution::search_matrix(vec![vec![], vec![]], 1), false);
    }
}
