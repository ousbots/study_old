// Calculate how many days until the next highest temperature.
pub fn daily_temperatures(temperatures: Vec<i32>) -> Vec<i32> {
    let mut days: Vec<i32> = vec![];
    let mut stack: Vec<(i32, usize)> = vec![];

    for (index, temp) in temperatures.iter().rev().enumerate() {
        while stack.len() > 0 && *temp >= stack.last().unwrap().0 {
            stack.pop();
        }

        if stack.len() > 0 {
            days.push((index - stack.last().unwrap().1) as i32);
        } else {
            days.push(0);
        }

        stack.push((*temp, index));
    }

    days.reverse();

    days
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn provided_tests() {
        assert_eq!(
            daily_temperatures(vec![73, 74, 75, 71, 69, 72, 76, 73]),
            [1, 1, 4, 2, 1, 1, 0, 0]
        );

        assert_eq!(daily_temperatures(vec![30, 40, 50, 60]), [1, 1, 1, 0]);

        assert_eq!(daily_temperatures(vec![30, 60, 90]), [1, 1, 0]);
    }
}
