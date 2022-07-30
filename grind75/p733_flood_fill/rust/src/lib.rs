pub struct Solution;

impl Solution {
    pub fn flood_fill(image: Vec<Vec<i32>>, sr: i32, sc: i32, new_color: i32) -> Vec<Vec<i32>> {
        let color = image[sr as usize][sc as usize];

        if color == new_color {
            return image;
        }

        let mut image = image;
        let mut queue = vec![(sr as usize, sc as usize)];

        while queue.len() > 0 {
            let (r, c) = queue.pop().unwrap();

            image[r][c] = new_color;

            if c > 0 && image[r][c - 1] == color {
                queue.push((r, c - 1));
            }

            if c < image[r].len() - 1 && image[r][c + 1] == color {
                queue.push((r, c + 1));
            }

            if r > 0 && image[r - 1][c] == color {
                queue.push((r - 1, c));
            }

            if r < image.len() - 1 && image[r + 1][c] == color {
                queue.push((r + 1, c));
            }
        }

        image
    }

    // Recursion is much faster than iteration.
    pub fn flood_fill_recurse(
        mut image: Vec<Vec<i32>>,
        sr: i32,
        sc: i32,
        new_color: i32,
    ) -> Vec<Vec<i32>> {
        let color = image[sr as usize][sc as usize];

        if color != new_color {
            Self::change(&mut image, sr as usize, sc as usize, color, new_color);
        }

        image
    }

    pub fn change(image: &mut Vec<Vec<i32>>, sr: usize, sc: usize, color: i32, new_color: i32) {
        if image[sr][sc] == color {
            image[sr][sc] = new_color;

            if sc > 0 {
                Self::change(image, sr, sc - 1, color, new_color);
            }

            if sc < image[sr].len() - 1 {
                Self::change(image, sr, sc + 1, color, new_color);
            }

            if sr > 0 {
                Self::change(image, sr - 1, sc, color, new_color);
            }

            if sr < image.len() - 1 {
                Self::change(image, sr + 1, sc, color, new_color);
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    fn check(a: Vec<Vec<i32>>, b: Vec<Vec<i32>>) -> bool {
        if a.len() != b.len() {
            return false;
        }

        for outer in 0..a.len() {
            if a[outer].len() != b[outer].len() {
                return false;
            }

            for inner in 0..a[outer].len() {
                if a[outer][inner] != b[outer][inner] {
                    return false;
                }
            }
        }

        true
    }

    #[test]
    fn provided() {
        let input = vec![vec![1, 1, 1], vec![1, 1, 0], vec![1, 0, 1]];
        let expected = vec![vec![2, 2, 2], vec![2, 2, 0], vec![2, 0, 1]];
        assert_eq!(check(Solution::flood_fill(input, 1, 1, 2), expected), true);

        let input = vec![vec![0, 0, 0], vec![0, 0, 0]];
        let expected = vec![vec![2, 2, 2], vec![2, 2, 2]];
        assert_eq!(check(Solution::flood_fill(input, 0, 0, 2), expected), true);
    }

    #[test]
    fn basic() {
        let input = vec![vec![1, 1, 1], vec![1, 1, 0], vec![1, 0, 1]];
        let expected = vec![vec![1, 1, 1], vec![1, 1, 0], vec![1, 0, 1]];
        assert_eq!(check(Solution::flood_fill(input, 1, 1, 1), expected), true);

        let input = vec![vec![1, 1, 0], vec![0, 1, 0], vec![1, 1, 1]];
        let expected = vec![vec![2, 2, 0], vec![0, 2, 0], vec![2, 2, 2]];
        assert_eq!(check(Solution::flood_fill(input, 0, 0, 2), expected), true);
    }
}
