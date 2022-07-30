pub struct Solution;

#[derive(PartialEq, Clone)]
enum State {
    None,
    Seen,
    Start,
}

impl Solution {
    pub fn can_finish(num_courses: i32, prerequisites: Vec<Vec<i32>>) -> bool {
        let mut graph: Vec<Vec<i32>> = vec![vec![]; num_courses as usize];
        for pre in prerequisites {
            graph[pre[1] as usize].push(pre[0]);
        }

        let mut states = vec![State::None; num_courses as usize];

        for course in 0..num_courses {
            if !Solution::search(&graph, &mut states, course as usize) {
                return false;
            }
        }

        true
    }

    fn search(graph: &Vec<Vec<i32>>, states: &mut Vec<State>, node: usize) -> bool {
        if states[node] == State::Start {
            return false;
        }

        if states[node] == State::Seen {
            return true;
        }

        states[node] = State::Start;
        for neighbor in &graph[node] {
            if !Solution::search(graph, states, *neighbor as usize) {
                return false;
            }
        }
        states[node] = State::Seen;

        true
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(Solution::can_finish(2, vec![vec![1, 0]]), true);
        assert_eq!(Solution::can_finish(2, vec![vec![1, 0], vec![0, 1]]), false);
    }

    #[test]
    fn basic() {
        assert_eq!(Solution::can_finish(3, vec![vec![0, 1], vec![1, 2]]), true);
        assert_eq!(
            Solution::can_finish(3, vec![vec![2, 1], vec![1, 0], vec![0, 2]]),
            false
        );

        assert_eq!(
            Solution::can_finish(
                20,
                vec![
                    vec![0, 10],
                    vec![3, 18],
                    vec![5, 5],
                    vec![6, 11],
                    vec![11, 1],
                    vec![13, 1],
                    vec![15, 1],
                    vec![17, 4]
                ]
            ),
            false
        );

        assert_eq!(
            Solution::can_finish(5, vec![vec![1, 4], vec![2, 4], vec![3, 1], vec![3, 2]]),
            true
        );
    }
}
