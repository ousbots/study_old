use std::collections::VecDeque;
pub struct Solution;

struct Point {
    x: usize,
    y: usize,
}

impl Solution {
    pub fn update_matrix(mat: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut dist: Vec<Vec<i32>> = vec![];
        let mut queue: VecDeque<Point> = VecDeque::new();

        for y in 0..mat.len() {
            dist.push(vec![]);
            for x in 0..mat[y].len() {
                if mat[y][x] == 0 {
                    dist[y].push(0);
                    queue.push_back(Point { x, y });
                } else {
                    dist[y].push(-1);
                }
            }
        }

        while queue.len() > 0 {
            let tmp = queue.pop_front().unwrap();
            let value = dist[tmp.y][tmp.x] + 1;

            if tmp.y > 0 {
                if dist[tmp.y - 1][tmp.x] == -1 {
                    dist[tmp.y - 1][tmp.x] = value;
                    queue.push_back(Point {
                        x: tmp.x,
                        y: tmp.y - 1,
                    });
                } else if value < dist[tmp.y - 1][tmp.x] {
                    dist[tmp.y - 1][tmp.x] = value;
                }
            }

            if tmp.y < mat.len() - 1 {
                if dist[tmp.y + 1][tmp.x] == -1 {
                    dist[tmp.y + 1][tmp.x] = value;
                    queue.push_back(Point {
                        x: tmp.x,
                        y: tmp.y + 1,
                    });
                } else if value < dist[tmp.y + 1][tmp.x] {
                    dist[tmp.y + 1][tmp.x] = value;
                }
            }

            if tmp.x > 0 {
                if dist[tmp.y][tmp.x - 1] == -1 {
                    dist[tmp.y][tmp.x - 1] = value;
                    queue.push_back(Point {
                        x: tmp.x - 1,
                        y: tmp.y,
                    });
                } else if value < dist[tmp.y][tmp.x - 1] {
                    dist[tmp.y][tmp.x - 1] = value;
                }
            }

            if tmp.x < mat[tmp.y].len() - 1 {
                if dist[tmp.y][tmp.x + 1] == -1 {
                    dist[tmp.y][tmp.x + 1] = value;
                    queue.push_back(Point {
                        x: tmp.x + 1,
                        y: tmp.y,
                    });
                } else if value < dist[tmp.y][tmp.x + 1] {
                    dist[tmp.y][tmp.x + 1] = value;
                }
            }
        }

        dist
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        let input = vec![vec![0, 0, 0], vec![0, 1, 0], vec![0, 0, 0]];
        let valid = vec![vec![0, 0, 0], vec![0, 1, 0], vec![0, 0, 0]];
        assert_eq!(Solution::update_matrix(input), valid);

        let input = vec![vec![0, 0, 0], vec![0, 1, 0], vec![1, 1, 1]];
        let valid = vec![vec![0, 0, 0], vec![0, 1, 0], vec![1, 2, 1]];
        assert_eq!(Solution::update_matrix(input), valid);
    }

    #[test]
    fn basic() {
        let input = vec![vec![1, 1, 0], vec![1, 1, 1], vec![0, 1, 1]];
        let valid = vec![vec![2, 1, 0], vec![1, 2, 1], vec![0, 1, 2]];
        assert_eq!(Solution::update_matrix(input), valid);

        let input = vec![vec![1, 1, 0], vec![1, 1, 1], vec![1, 1, 0]];
        let valid = vec![vec![2, 1, 0], vec![3, 2, 1], vec![2, 1, 0]];
        assert_eq!(Solution::update_matrix(input), valid);

        let input = vec![vec![0, 1, 1], vec![1, 1, 1], vec![1, 1, 1]];
        let valid = vec![vec![0, 1, 2], vec![1, 2, 3], vec![2, 3, 4]];
        assert_eq!(Solution::update_matrix(input), valid);
    }
}
