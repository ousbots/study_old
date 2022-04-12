pub struct Solution;

impl Solution {
    pub fn game_of_life(board: &mut Vec<Vec<i32>>) {
        let orig = board.clone();

        for y in 0..board.len() {
            for x in 0..board[y].len() {
                let neighbors = neighbor_count(&orig, x, y);

                if orig[y][x] == 0 {
                    if neighbors == 3 {
                        board[y][x] = 1;
                    }
                } else {
                    if neighbors < 2 || neighbors > 3 {
                        board[y][x] = 0;
                    }
                }
            }
        }
    }
}

fn neighbor_count(board: &Vec<Vec<i32>>, x: usize, y: usize) -> i32 {
    let mut count = 0;

    if x > 0 {
        if board[y][x - 1] == 1 {
            count += 1;
        }

        if y > 0 {
            if board[y - 1][x - 1] == 1 {
                count += 1;
            }
        }

        if y < board.len() - 1 {
            if board[y + 1][x - 1] == 1 {
                count += 1;
            }
        }
    }

    if x < board[y].len() - 1 {
        if board[y][x + 1] == 1 {
            count += 1;
        }

        if y > 0 {
            if board[y - 1][x + 1] == 1 {
                count += 1;
            }
        }

        if y < board.len() - 1 {
            if board[y + 1][x + 1] == 1 {
                count += 1;
            }
        }
    }

    if y > 0 {
        if board[y - 1][x] == 1 {
            count += 1;
        }
    }

    if y < board.len() - 1 {
        if board[y + 1][x] == 1 {
            count += 1;
        }
    }

    count
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        let mut board = vec![vec![1, 1], vec![1, 0]];
        let ans = vec![vec![1, 1], vec![1, 1]];
        Solution::game_of_life(&mut board);
        assert_eq!(board, ans);

        let mut board = vec![vec![0, 1, 0], vec![0, 0, 1], vec![1, 1, 1]];
        let ans = vec![vec![0, 0, 0], vec![1, 0, 1], vec![0, 1, 1]];
        Solution::game_of_life(&mut board);
        assert_eq!(board, ans);
    }
}
