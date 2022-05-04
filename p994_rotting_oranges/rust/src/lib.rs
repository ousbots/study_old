pub struct Solution;
impl Solution {
    pub fn oranges_rotting(grid: Vec<Vec<i32>>) -> i32 {
        let mut change = vec![];
        let mut grid = grid;
        let mut time = 0;

        loop {
            for y in 0..grid.len() {
                for x in 0..grid[y].len() {
                    if grid[y][x] == 2 {
                        Solution::rot(&mut grid, &mut change, y, x);
                    }
                }
            }

            if !change.is_empty() {
                for pair in &change {
                    grid[pair.0][pair.1] = 2;
                }
                change.clear();
                time += 1;
            } else {
                break;
            }
        }

        for y in 0..grid.len() {
            for x in 0..grid[y].len() {
                if grid[y][x] == 1 {
                    return -1;
                }
            }
        }

        time
    }

    fn rot(grid: &mut Vec<Vec<i32>>, change: &mut Vec<(usize, usize)>, y: usize, x: usize) {
        let mut visit = vec![(y, x)];

        while let Some(check) = visit.pop() {
            let val = grid[check.0][check.1];
            if val == 2 {
                grid[check.0][check.1] = 3;

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
                    visit.push((check.0, check.1 + 1));
                }
            }

            if val == 1 {
                change.push(check);
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        let input = vec![vec![2, 1, 1], vec![1, 1, 0], vec![0, 1, 1]];
        assert_eq!(Solution::oranges_rotting(input), 4);

        let input = vec![vec![2, 1, 1], vec![0, 1, 1], vec![1, 0, 1]];
        assert_eq!(Solution::oranges_rotting(input), -1);

        assert_eq!(Solution::oranges_rotting(vec![vec![0, 2]]), 0);
    }

    #[test]
    fn basic() {
        assert_eq!(Solution::oranges_rotting(vec![vec![0]]), 0);
        assert_eq!(Solution::oranges_rotting(vec![vec![2]]), 0);
        assert_eq!(Solution::oranges_rotting(vec![vec![1]]), -1);

        let input = vec![
            vec![2, 1, 1, 1],
            vec![1, 1, 1, 1],
            vec![1, 1, 1, 1],
            vec![1, 1, 1, 1],
        ];
        assert_eq!(Solution::oranges_rotting(input), 6);

        let input = vec![
            vec![2, 1, 1, 1],
            vec![1, 1, 1, 1],
            vec![1, 1, 1, 1],
            vec![1, 1, 1, 2],
        ];
        assert_eq!(Solution::oranges_rotting(input), 3);

        let input = vec![
            vec![2, 1, 2, 1],
            vec![1, 2, 1, 2],
            vec![2, 1, 2, 1],
            vec![1, 2, 1, 2],
        ];
        assert_eq!(Solution::oranges_rotting(input), 1);

        let input = vec![
            vec![2, 0, 2, 0],
            vec![0, 1, 0, 1],
            vec![2, 0, 2, 0],
            vec![0, 1, 0, 1],
        ];
        assert_eq!(Solution::oranges_rotting(input), -1);
    }
}
