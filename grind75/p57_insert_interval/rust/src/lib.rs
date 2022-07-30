pub struct Solution;

impl Solution {
    pub fn insert_slow(intervals: Vec<Vec<i32>>, new_interval: Vec<i32>) -> Vec<Vec<i32>> {
        if intervals.len() == 0 {
            return vec![new_interval];
        }

        let mut intervals = intervals;

        let mut front = bsearch(&intervals, new_interval[0]);
        if front != intervals.len() - 1 && intervals[front][1] < new_interval[0] {
            front += 1;
        }

        if front == intervals.len() - 1 && intervals[front][1] < new_interval[0] {
            intervals.push(new_interval);
            return intervals;
        }

        let mut back = bsearch(&intervals, new_interval[1]);
        if back != 0 && intervals[back][0] > new_interval[1] {
            back -= 1;
        }

        if back == 0 && intervals[back][0] > new_interval[1] {
            let mut temp = vec![new_interval];
            temp.append(&mut intervals);

            return temp;
        }

        if front > back {
            intervals.splice(front..front, vec![new_interval.clone()].into_iter());
            return intervals;
        }

        if intervals[front][0] > new_interval[0] {
            intervals[front][0] = new_interval[0];
        }

        if intervals[back][1] < new_interval[1] {
            intervals[back][1] = new_interval[1];
        }

        if front != back {
            intervals[front][1] = intervals[back][1];
            intervals.splice(front + 1..back + 1, vec![].into_iter());
        }

        intervals
    }

    pub fn insert(intervals: Vec<Vec<i32>>, new_interval: Vec<i32>) -> Vec<Vec<i32>> {
        let mut intervals = intervals;
        intervals.push(new_interval);
        intervals.sort();

        let mut result: Vec<Vec<i32>> = vec![];
        for element in intervals {
            if let Some(tail) = result.last_mut() {
                if tail[1] < element[0] {
                    result.push(element);
                } else {
                    tail[1] = tail[1].max(element[1]);
                }
            } else {
                result.push(element);
            }
        }

        result
    }
}

pub fn bsearch(intervals: &Vec<Vec<i32>>, val: i32) -> usize {
    let mut low: usize = 0;
    let mut high: usize = intervals.len() - 1;

    while low <= high {
        let middle = (low + high) / 2;

        if intervals[middle][0] > val {
            if middle == 0 {
                break;
            }
            high = middle - 1;
        } else if intervals[middle][1] < val {
            low = middle + 1;
        } else {
            return middle;
        }
    }

    (low + high) / 2
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        assert_eq!(
            Solution::insert(vec![vec![1, 3], vec![6, 9]], vec![2, 5]),
            vec![vec![1, 5], vec![6, 9]]
        );

        assert_eq!(
            Solution::insert(
                vec![
                    vec![1, 2],
                    vec![3, 5],
                    vec![6, 7],
                    vec![8, 10],
                    vec![12, 16]
                ],
                vec![4, 8]
            ),
            vec![vec![1, 2], vec![3, 10], vec![12, 16]]
        );
    }

    #[test]
    fn basic() {
        assert_eq!(
            Solution::insert(vec![vec![0, 2],], vec![4, 6]),
            vec![vec![0, 2], vec![4, 6],]
        );

        assert_eq!(
            Solution::insert(vec![vec![4, 6],], vec![0, 2]),
            vec![vec![0, 2], vec![4, 6],]
        );

        assert_eq!(
            Solution::insert(
                vec![
                    vec![0, 2],
                    vec![4, 6],
                    vec![8, 10],
                    vec![12, 14],
                    vec![16, 18]
                ],
                vec![7, 7]
            ),
            vec![
                vec![0, 2],
                vec![4, 6],
                vec![7, 7],
                vec![8, 10],
                vec![12, 14],
                vec![16, 18]
            ]
        );

        assert_eq!(
            Solution::insert(
                vec![
                    vec![0, 2],
                    vec![4, 6],
                    vec![8, 10],
                    vec![12, 14],
                    vec![16, 18]
                ],
                vec![11, 11]
            ),
            vec![
                vec![0, 2],
                vec![4, 6],
                vec![8, 10],
                vec![11, 11],
                vec![12, 14],
                vec![16, 18]
            ]
        );

        assert_eq!(
            Solution::insert(
                vec![
                    vec![0, 2],
                    vec![4, 6],
                    vec![8, 10],
                    vec![12, 14],
                    vec![16, 18]
                ],
                vec![5, 7]
            ),
            vec![
                vec![0, 2],
                vec![4, 7],
                vec![8, 10],
                vec![12, 14],
                vec![16, 18]
            ]
        );

        assert_eq!(
            Solution::insert(
                vec![
                    vec![0, 2],
                    vec![4, 6],
                    vec![8, 10],
                    vec![12, 14],
                    vec![16, 18]
                ],
                vec![5, 9]
            ),
            vec![vec![0, 2], vec![4, 10], vec![12, 14], vec![16, 18]]
        );

        assert_eq!(
            Solution::insert(
                vec![
                    vec![0, 2],
                    vec![4, 6],
                    vec![8, 10],
                    vec![12, 14],
                    vec![16, 18]
                ],
                vec![11, 13]
            ),
            vec![
                vec![0, 2],
                vec![4, 6],
                vec![8, 10],
                vec![11, 14],
                vec![16, 18]
            ]
        );

        assert_eq!(
            Solution::insert(
                vec![
                    vec![0, 2],
                    vec![4, 6],
                    vec![8, 10],
                    vec![12, 14],
                    vec![16, 18]
                ],
                vec![9, 13]
            ),
            vec![vec![0, 2], vec![4, 6], vec![8, 14], vec![16, 18]]
        );

        assert_eq!(
            Solution::insert(
                vec![
                    vec![0, 2],
                    vec![4, 6],
                    vec![8, 10],
                    vec![12, 14],
                    vec![16, 18]
                ],
                vec![1, 17]
            ),
            vec![vec![0, 18],]
        );

        assert_eq!(
            Solution::insert(
                vec![
                    vec![0, 2],
                    vec![4, 6],
                    vec![8, 10],
                    vec![12, 14],
                    vec![16, 18]
                ],
                vec![-1, 19]
            ),
            vec![vec![-1, 19],]
        );
    }
}
