pub struct Solution;
impl Solution {
    pub fn spiral_order(matrix: Vec<Vec<i32>>) -> Vec<i32> {
        let mut xmax: i32 = (matrix[0].len() - 1) as i32;
        let mut ymax: i32 = (matrix.len() - 1) as i32;
        let mut xmin: i32 = 0;
        let mut ymin: i32 = 0;

        let mut output = Vec::with_capacity(matrix.len() * matrix[0].len());
        let mut direction = 0;

        let mut x: i32 = 0;
        let mut y: i32 = 0;

        loop {
            output.push(matrix[y as usize][x as usize]);
            match direction {
                0 => {
                    if x == xmax {
                        direction = 1;
                        ymin += 1;
                        y += 1;
                    } else {
                        x += 1;
                    }
                }
                1 => {
                    if y == ymax {
                        direction = 2;
                        xmax -= 1;
                        x -= 1;
                    } else {
                        y += 1;
                    }
                }
                2 => {
                    if x == xmin {
                        direction = 3;
                        ymax -= 1;
                        y -= 1;
                    } else {
                        x -= 1;
                    }
                }
                3 => {
                    if y == ymin {
                        direction = 0;
                        xmin += 1;
                        x += 1;
                    } else {
                        y -= 1;
                    }
                }
                _ => panic!("bad direction"),
            }

            if x < xmin || x > xmax || y < ymin || y > ymax {
                break;
            }
        }

        output
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn basic() {
        assert_eq!(Solution::spiral_order(vec![vec![1]]), vec![1]);

        assert_eq!(Solution::spiral_order(vec![vec![1, 2, 3]]), vec![1, 2, 3]);

        assert_eq!(
            Solution::spiral_order(vec![vec![1], vec![2], vec![3]]),
            vec![1, 2, 3]
        );

        assert_eq!(
            Solution::spiral_order(vec![vec![1, 2, 3], vec![8, 9, 4], vec![7, 6, 5]]),
            vec![1, 2, 3, 4, 5, 6, 7, 8, 9]
        );
    }
}
