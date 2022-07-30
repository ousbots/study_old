use std::collections::HashMap;

pub struct Solution;

impl Solution {
    pub fn find_anagram(s: String, p: String) -> Vec<i32> {
        let input_ref: Vec<char> = p.chars().collect();
        let mut count_ref: HashMap<&char, i32> = HashMap::new();

        for elem in &input_ref {
            *count_ref.entry(elem).or_insert(0) += 1;
        }

        let mut pos_l = 0;
        let mut pos_r = 0;
        let mut count = HashMap::new();

        let input: Vec<char> = s.chars().collect();
        let mut starting: Vec<i32> = vec![];

        while pos_r < input.len() {
            *count.entry(&input[pos_r]).or_insert(0) += 1;

            if pos_r - pos_l == input_ref.len() - 1 {
                if maps_equal(&count, &count_ref) {
                    starting.push(pos_l as i32);
                }

                *count.entry(&input[pos_l]).or_insert(1) -= 1;
                if count[&input[pos_l]] == 0 {
                    count.remove(&input[pos_l]);
                }
                pos_l += 1;
            }

            pos_r += 1;
        }

        starting
    }
}

fn maps_equal<'a>(first: &HashMap<&'a char, i32>, second: &HashMap<&'a char, i32>) -> bool {
    if first.len() != second.len() {
        return false;
    }

    for key in first.keys() {
        if !second.contains_key(key) {
            return false;
        }

        if first[key] != second[key] {
            return false;
        }
    }

    true
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn equals_test() {
        assert_eq!(
            maps_equal(&HashMap::from([(&'a', 1)]), &HashMap::from([(&'a', 1)])),
            true,
        );
        assert_eq!(maps_equal(&HashMap::new(), &HashMap::new()), true);
        assert_eq!(
            maps_equal(
                &HashMap::from([(&'a', 1), (&'b', 5)]),
                &HashMap::from([(&'a', 1), (&'b', 5)]),
            ),
            true,
        );
        assert_eq!(
            maps_equal(
                &HashMap::from([(&'a', 1), (&'b', 5)]),
                &HashMap::from([(&'a', 1), (&'b', 4)]),
            ),
            false,
        );
        assert_eq!(
            maps_equal(
                &HashMap::from([(&'a', 1), (&'b', 5)]),
                &HashMap::from([(&'a', 1), (&'b', 5), (&'c', 0)]),
            ),
            false,
        );
        assert_eq!(
            maps_equal(&HashMap::new(), &HashMap::from([(&'a', 1)])),
            false
        );
        assert_eq!(
            maps_equal(&HashMap::from([(&'a', 1)]), &HashMap::new()),
            false
        );
    }

    #[test]
    fn basic_test() {
        assert_eq!(
            Solution::find_anagram("abcdcab".to_string(), "abc".to_string()),
            vec![0, 4]
        );
        assert_eq!(
            Solution::find_anagram("abcdcab".to_string(), "a".to_string()),
            vec![0, 5]
        );
        assert_eq!(
            Solution::find_anagram("abab".to_string(), "ab".to_string()),
            vec![0, 1, 2]
        );
        assert_eq!(
            Solution::find_anagram("".to_string(), "".to_string()),
            vec![]
        );
    }
}
