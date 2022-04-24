pub struct MyQueue {
    input: Vec<i32>,
    output: Vec<i32>,
}

// Original implementation.
//impl MyQueue {
//    pub fn new() -> Self {
//        return Self {
//            input: vec![],
//            output: vec![],
//        };
//    }

//    pub fn push(&mut self, x: i32) {
//        while self.output.len() != 0 {
//            self.input.push(self.output.pop().unwrap());
//        }

//        self.input.push(x);
//    }

//    pub fn pop(&mut self) -> i32 {
//        while self.input.len() != 0 {
//            self.output.push(self.input.pop().unwrap());
//        }

//        self.output.pop().unwrap()
//    }

//    pub fn peek(&mut self) -> i32 {
//        while self.input.len() != 0 {
//            self.output.push(self.input.pop().unwrap());
//        }

//        *self.output.last().unwrap()
//    }

//    pub fn empty(&self) -> bool {
//        return self.input.len() + self.output.len() == 0;
//    }
//}

impl MyQueue {
    pub fn new() -> Self {
        return Self {
            input: vec![],
            output: vec![],
        };
    }

    pub fn push(&mut self, x: i32) {
        self.input.push(x);
    }

    pub fn pop(&mut self) -> i32 {
        if self.output.is_empty() {
            while self.input.len() != 0 {
                self.output.push(self.input.pop().unwrap());
            }
        }

        self.output.pop().unwrap()
    }

    pub fn peek(&mut self) -> i32 {
        if self.output.is_empty() {
            while self.input.len() != 0 {
                self.output.push(self.input.pop().unwrap());
            }
        }

        *self.output.last().unwrap()
    }

    pub fn empty(&self) -> bool {
        return self.input.len() + self.output.len() == 0;
    }
}

#[cfg(test)]
mod tests {
    use super::MyQueue;

    #[test]
    fn provided() {
        let mut queue = MyQueue::new();
        queue.push(1);
        queue.push(2);

        assert_eq!(queue.peek(), 1);
        assert_eq!(queue.pop(), 1);
        assert_eq!(queue.empty(), false);
    }

    #[test]
    fn basic() {
        let mut queue = MyQueue::new();
        queue.push(1);
        queue.push(2);

        assert_eq!(queue.peek(), 1);
        assert_eq!(queue.pop(), 1);
        assert_eq!(queue.empty(), false);

        queue.push(3);
        assert_eq!(queue.peek(), 2);
        assert_eq!(queue.pop(), 2);
        assert_eq!(queue.empty(), false);

        assert_eq!(queue.peek(), 3);
        assert_eq!(queue.pop(), 3);
        assert_eq!(queue.empty(), true);
    }
}
