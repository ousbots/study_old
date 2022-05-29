use std::collections::HashMap;

pub struct Solution;
impl Solution {
    pub fn accounts_merge(accounts: Vec<Vec<String>>) -> Vec<Vec<String>> {
        let mut names = HashMap::<String, String>::new();
        let mut emails = HashMap::<String, HashMap<String, bool>>::new();

        for line in accounts {
            let name = line[0].clone();

            for i in 1..line.len() {
                let email = line[i].clone();
                names.insert(email.clone(), name.clone());

                let entry = emails.entry(email.clone()).or_insert(HashMap::new());

                if i == 1 {
                    continue;
                }

                let other = line[i - 1].clone();
                entry.insert(other.clone(), true);

                let entry = emails.entry(other).or_insert(HashMap::new());
                entry.insert(email, true);
            }
        }

        let mut seen = HashMap::<String, bool>::new();
        let mut output = vec![];

        for (email, _) in &emails {
            let mut que = vec![email];
            let mut list = vec![];

            while let Some(key) = que.pop() {
                if seen.contains_key(key) {
                    continue;
                }
                seen.insert(key.clone(), true);

                list.push(key.clone());
                for (other, _) in emails.get(key).unwrap() {
                    que.push(other);
                }
            }

            if !list.is_empty() {
                list.sort_unstable();
                list.insert(0, names.get(email).unwrap().clone());
                output.push(list);
            }
        }

        output
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn provided() {
        let input = vec![
            vec!["John", "johnsmith@mail.com", "john_newyork@mail.com"],
            vec!["John", "johnsmith@mail.com", "john00@mail.com"],
            vec!["Mary", "mary@mail.com"],
            vec!["John", "johnnybravo@mail.com"],
        ]
        .into_iter()
        .map(|i| i.into_iter().map(|j| j.to_string()).collect())
        .collect();
        let mut valid = vec![
            vec![
                "John",
                "john00@mail.com",
                "john_newyork@mail.com",
                "johnsmith@mail.com",
            ],
            vec!["Mary", "mary@mail.com"],
            vec!["John", "johnnybravo@mail.com"],
        ];
        valid.sort_unstable();
        let mut ans = Solution::accounts_merge(input);
        ans.sort_unstable();
        assert_eq!(ans, valid);

        let input = vec![
            vec!["Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"],
            vec!["Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"],
            vec!["Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"],
            vec!["Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"],
            vec!["Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"],
        ]
        .into_iter()
        .map(|i| i.into_iter().map(|j| j.to_string()).collect())
        .collect();
        let mut valid = vec![
            vec!["Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"],
            vec!["Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"],
            vec!["Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"],
            vec!["Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"],
            vec!["Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"],
        ];
        valid.sort_unstable();
        let mut ans = Solution::accounts_merge(input);
        ans.sort_unstable();
        assert_eq!(ans, valid);
    }

    #[test]
    fn basic() {
        let input = vec![
            vec!["Mary", "marypoopins@mail.com"],
            vec!["John", "joon@mail.com", "jern@mail.com"],
            vec!["John", "johnsmith@mail.com", "john_newyork@mail.com"],
            vec!["John", "johnsmith@mail.com", "john00@mail.com"],
            vec!["John", "joon@mail.com", "john00@mail.com"],
            vec!["Mary", "mary@mail.com"],
            vec![
                "Mary",
                "mary@mail.com",
                "marypoppins@mail.com",
                "marylittlelamb@mail.com",
            ],
            vec!["Mary", "marypoppins@mail.com", "mary@mail.com"],
            vec!["John", "johnnybravo@mail.com"],
        ]
        .into_iter()
        .map(|i| i.into_iter().map(|j| j.to_string()).collect())
        .collect();
        let mut valid = vec![
            vec![
                "John",
                "jern@mail.com",
                "john00@mail.com",
                "john_newyork@mail.com",
                "johnsmith@mail.com",
                "joon@mail.com",
            ],
            vec![
                "Mary",
                "mary@mail.com",
                "marylittlelamb@mail.com",
                "marypoppins@mail.com",
            ],
            vec!["Mary", "marypoopins@mail.com"],
            vec!["John", "johnnybravo@mail.com"],
        ];
        valid.sort_unstable();
        let mut ans = Solution::accounts_merge(input);
        ans.sort_unstable();
        assert_eq!(ans, valid);
    }
}
