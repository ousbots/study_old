pub struct Solution;

impl Solution {
    pub fn k_closest(points: Vec<Vec<i32>>, k: i32) -> Vec<Vec<i32>> {
        let mut points = points;
        points.sort_by(|a, b| (a[0] * a[0] + a[1] * a[1]).cmp(&(b[0] * b[0] + b[1] * b[1])));

        points.drain(0..k as usize).collect()
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }
}
