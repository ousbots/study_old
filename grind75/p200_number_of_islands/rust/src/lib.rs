pub struct Solution;
impl Solution {
    pub fn num_islands(grid: Vec<Vec<char>>) -> i32 {
        let mut count = 0;
        let mut grid = grid;

        for y in 0..grid.len() {
            for x in 0..grid[y].len() {
                if grid[y][x] == '1' {
                    count += 1;
                    Solution::sink(&mut grid, y, x);
                }
            }
        }

        count
    }

    fn sink(grid: &mut Vec<Vec<char>>, y: usize, x: usize) {
        let mut visit = vec![(y, x)];

        while let Some(check) = visit.pop() {
            if grid[check.0][check.1] == '1' {
                grid[check.0][check.1] = '0';

                if check.0 > 0 {
                    visit.push((check.0 - 1, check.1));
                }

                if check.0 < grid.len() - 1 {
                    visit.push((check.0 + 1, check.1));
                }

                if check.1 > 0 {
                    visit.push((check.0, check.1 - 1));
                }

                if check.1 < grid[check.0].len() - 1 {
                    visit.push((check.0, check.1 + 1))
                }
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        let input = vec![
            vec!['1', '1', '1', '1', '0'],
            vec!['1', '1', '0', '1', '0'],
            vec!['1', '1', '0', '0', '0'],
            vec!['0', '0', '0', '0', '0'],
        ];
        assert_eq!(Solution::num_islands(input), 1);

        let input = vec![
            vec!['1', '1', '0', '0', '0'],
            vec!['1', '1', '0', '0', '0'],
            vec!['0', '0', '1', '0', '0'],
            vec!['0', '0', '0', '1', '1'],
        ];
        assert_eq!(Solution::num_islands(input), 3);
    }

    #[test]
    fn basic() {
        let input = vec![
            vec!['1', '1', '1', '1', '1', '1', '1'],
            vec!['1', '0', '0', '0', '0', '0', '1'],
            vec!['1', '0', '1', '1', '1', '0', '1'],
            vec!['1', '0', '1', '0', '1', '0', '1'],
            vec!['1', '0', '1', '0', '0', '0', '1'],
            vec!['1', '0', '1', '1', '1', '1', '1'],
        ];
        assert_eq!(Solution::num_islands(input), 1);

        let input = vec![
            vec!['1', '0', '1', '0', '1'],
            vec!['0', '1', '0', '1', '0'],
            vec!['1', '0', '1', '0', '1'],
            vec!['0', '1', '0', '1', '0'],
        ];
        assert_eq!(Solution::num_islands(input), 10);

        let input = vec![
            vec!['1', '1', '1', '0', '1', '1'],
            vec!['1', '1', '0', '0', '1', '1'],
            vec!['0', '0', '1', '1', '0', '1'],
            vec!['1', '0', '1', '1', '0', '0'],
            vec!['1', '1', '0', '0', '1', '1'],
            vec!['1', '1', '0', '1', '1', '1'],
        ];
        assert_eq!(Solution::num_islands(input), 5);

        let input = vec![vec!['1']];
        assert_eq!(Solution::num_islands(input), 1);

        let input = vec![vec!['0']];
        assert_eq!(Solution::num_islands(input), 0);
    }
}
