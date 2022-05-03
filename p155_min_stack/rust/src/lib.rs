pub struct MinStack {
    min: i32,
    vals: Vec<i32>,
}

impl MinStack {
    pub fn new() -> Self {
        MinStack {
            min: i32::MAX,
            vals: vec![],
        }
    }

    pub fn push(&mut self, val: i32) {
        self.vals.push(val);
        if val < self.min {
            self.min = val;
        }
    }

    pub fn pop(&mut self) {
        if self.vals.pop().unwrap() == self.min {
            self.min = i32::MAX;
            if self.vals.len() != 0 {
                self.vals.iter().for_each(|x| {
                    if *x < self.min {
                        self.min = *x
                    }
                });
            }
        }
    }

    pub fn top(&self) -> i32 {
        *self.vals.last().unwrap()
    }

    pub fn get_min(&self) -> i32 {
        self.min
    }
}

#[cfg(test)]
mod tests {
    use super::MinStack;

    #[test]
    fn provided() {
        let mut stack = MinStack::new();
        stack.push(-2);
        stack.push(0);
        stack.push(-3);
        assert_eq!(stack.get_min(), -3);
        stack.pop();
        assert_eq!(stack.top(), 0);
        assert_eq!(stack.get_min(), -2);
    }

    #[test]
    fn basic() {
        let mut stack = MinStack::new();
        stack.push(1);
        stack.push(2);
        assert_eq!(stack.get_min(), 1);
        assert_eq!(stack.top(), 2);
        stack.pop();
        assert_eq!(stack.get_min(), 1);
        assert_eq!(stack.top(), 1);
        stack.pop();
        stack.push(3);
        assert_eq!(stack.get_min(), 3);
        assert_eq!(stack.top(), 3);
    }
}
