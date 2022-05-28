use std::collections::HashMap;

pub struct TimeMap {
    time: HashMap<String, Vec<i32>>,
    vals: HashMap<String, Vec<String>>,
}

impl TimeMap {
    pub fn new() -> Self {
        TimeMap {
            time: HashMap::new(),
            vals: HashMap::new(),
        }
    }

    pub fn set(&mut self, key: String, value: String, timestamp: i32) {
        let time = self.time.entry(key.clone()).or_insert(vec![]);
        let vals = self.vals.entry(key).or_insert(vec![]);

        time.push(timestamp);
        vals.push(value);
    }

    pub fn get(&self, key: String, timestamp: i32) -> String {
        if let Some(time) = self.time.get(&key) {
            if let Some(vals) = self.vals.get(&key) {
                let i = time.partition_point(|e| e <= &timestamp);

                if i == 0 {
                    return String::from("");
                }

                return vals[i - 1].clone();
            } else {
                panic!("inconsistent data for {}", key)
            }
        }

        String::from("")
    }
}

#[cfg(test)]
mod tests {
    use super::TimeMap;

    #[test]
    fn provided() {
        let mut tm = TimeMap::new();
        tm.set(String::from("foo"), String::from("bar"), 1);
        assert_eq!(tm.get(String::from("foo"), 1), "bar");
        assert_eq!(tm.get(String::from("foo"), 3), "bar");
        tm.set(String::from("foo"), String::from("bar2"), 4);
        assert_eq!(tm.get(String::from("foo"), 4), "bar2");
        assert_eq!(tm.get(String::from("foo"), 5), "bar2");
    }

    #[test]
    fn basic() {
        let mut tm = TimeMap::new();
        tm.set(String::from("exists"), String::from("a"), 1);
        tm.set(String::from("exists"), String::from("b"), 4);
        assert_eq!(tm.get(String::from("exists"), 0), "");
        assert_eq!(tm.get(String::from("exists"), 1), "a");
        assert_eq!(tm.get(String::from("exists"), 3), "a");
        assert_eq!(tm.get(String::from("exists"), 4), "b");
        assert_eq!(tm.get(String::from("exists"), 99), "b");
        assert_eq!(tm.get(String::from("doesn't exist"), 99), "");
    }
}
