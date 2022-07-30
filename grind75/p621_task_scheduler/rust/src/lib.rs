pub struct Solution;

impl Solution {
    // Find the highest frequency task, subtract one from its count. The maximum number of idle
    // cycles is the highest frequency task count times n. They can be filled in by the other
    // tasks (max of highest frequency count minus one). Then the minimum number of cycles is the
    // number of tasks plus any idle cycles left.
    pub fn least_interval(tasks: Vec<char>, n: i32) -> i32 {
        let mut counts = vec![0; 26];
        let total = tasks.len();

        for task in tasks {
            counts[task as usize - 'A' as usize] += 1;
        }
        counts.sort();

        let max_freq = counts[25] - 1;
        let mut idle_cycles = max_freq * n;

        let mut i = 24;
        loop {
            idle_cycles -= counts[i].min(max_freq);

            if i == 0 {
                break;
            } else {
                i -= 1;
            }
        }

        idle_cycles.max(0) + total as i32
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[derive(Debug)]
    struct TestCase {
        tasks: Vec<char>,
        n: i32,
        output: i32,
    }

    #[test]
    fn provided() {
        let cases: Vec<TestCase> = vec![
            TestCase {
                tasks: vec!['A', 'A', 'A', 'B', 'B', 'B'],
                n: 2,
                output: 8,
            },
            TestCase {
                tasks: vec!['A', 'A', 'A', 'B', 'B', 'B'],
                n: 0,
                output: 6,
            },
            TestCase {
                tasks: vec!['A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'],
                n: 2,
                output: 16,
            },
        ];

        for case in cases {
            assert_eq!(Solution::least_interval(case.tasks, case.n), case.output);
        }
    }

    #[test]
    fn basic() {
        let cases: Vec<TestCase> = vec![
            TestCase {
                tasks: vec!['A'],
                n: 0,
                output: 1,
            },
            TestCase {
                tasks: vec!['A'],
                n: 1,
                output: 1,
            },
            TestCase {
                tasks: vec!['A', 'A'],
                n: 0,
                output: 2,
            },
            TestCase {
                tasks: vec!['A', 'A'],
                n: 1,
                output: 3,
            },
            TestCase {
                tasks: vec!['A', 'B'],
                n: 1,
                output: 2,
            },
            TestCase {
                tasks: vec!['A', 'B'],
                n: 2,
                output: 2,
            },
            TestCase {
                tasks: vec!['A', 'A', 'B'],
                n: 1,
                output: 3,
            },
            TestCase {
                tasks: vec!['A', 'A', 'B'],
                n: 2,
                output: 4,
            },
            TestCase {
                tasks: vec!['A', 'A', 'A', 'B', 'C'],
                n: 1,
                output: 5,
            },
            TestCase {
                tasks: vec!['A', 'A', 'A', 'B', 'C'],
                n: 2,
                output: 7,
            },
        ];

        for case in cases {
            assert_eq!(Solution::least_interval(case.tasks, case.n), case.output);
        }
    }

    #[test]
    fn more() {
        let cases: Vec<TestCase> = vec![TestCase {
            tasks: vec![
                'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O',
                'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
            ],
            n: 29,
            output: 31,
        }];

        for case in cases {
            assert_eq!(Solution::least_interval(case.tasks, case.n), case.output);
        }
    }
}
