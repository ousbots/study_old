use std::collections::hash_map::Entry::{Occupied, Vacant};
use std::collections::HashMap;

pub struct Trie {
    end: bool,
    next: HashMap<char, Trie>,
}

impl Trie {
    pub fn new() -> Self {
        Trie {
            end: false,
            next: HashMap::new(),
        }
    }

    pub fn insert(&mut self, word: String) {
        let mut node = self;

        for c in word.chars() {
            let entry = match node.next.entry(c) {
                Occupied(n) => n.into_mut(),
                Vacant(n) => n.insert(Trie::new()),
            };

            node = entry;
        }

        node.end = true;
    }

    pub fn search(&self, word: String) -> bool {
        let mut node = self;

        for c in word.chars() {
            if let Some(next) = node.next.get(&c) {
                node = next;
            } else {
                return false;
            }
        }

        node.end
    }

    pub fn starts_with(&self, word: String) -> bool {
        let mut node = self;

        for c in word.chars() {
            if let Some(next) = node.next.get(&c) {
                node = next;
            } else {
                return false;
            }
        }

        true
    }
}

#[cfg(test)]
mod tests {
    use super::Trie;

    #[test]
    fn provided() {
        let mut trie = Trie::new();
        trie.insert("apple".to_string());
        assert_eq!(trie.search("apple".to_string()), true);
        assert_eq!(trie.search("app".to_string()), false);
        assert_eq!(trie.starts_with("app".to_string()), true);
        trie.insert("app".to_string());
        assert_eq!(trie.search("app".to_string()), true);
    }

    #[test]
    fn basic() {
        let mut trie = Trie::new();
        trie.insert("a".to_string());
        assert_eq!(trie.search("a".to_string()), true);
        assert_eq!(trie.search("ab".to_string()), false);
        assert_eq!(trie.search("b".to_string()), false);
        assert_eq!(trie.starts_with("a".to_string()), true);

        trie.insert("falafe".to_string());
        trie.insert("falafel".to_string());
        assert_eq!(trie.search("falaf".to_string()), false);
        assert_eq!(trie.search("falafe".to_string()), true);
        assert_eq!(trie.search("falafel".to_string()), true);
        assert_eq!(trie.starts_with("falaf".to_string()), true);
        assert_eq!(trie.starts_with("falafe".to_string()), true);
        assert_eq!(trie.starts_with("falafel".to_string()), true);
    }
}
