package implement_trie

type Node struct {
	Char rune
	End  bool
	Next map[rune]*Node
}

type Trie struct {
	Root Node
}

func Constructor() Trie {
	return Trie{Root: Node{Next: make(map[rune]*Node)}}
}

func (this *Trie) Insert(word string) {
	node := &this.Root

	for _, char := range word {
		if child := node.Next[char]; child != nil {
			node = child
		} else {
			temp := Node{Char: char, Next: make(map[rune]*Node)}
			node.Next[char] = &temp
			node = &temp
		}
	}
	node.End = true
}

func (this *Trie) Search(word string) bool {
	node := &this.Root

	for _, char := range word {
		if child := node.Next[char]; child != nil {
			node = child
		} else {
			return false
		}
	}

	return node.End
}

func (this *Trie) StartsWith(word string) bool {
	node := &this.Root

	for _, char := range word {
		if child := node.Next[char]; child != nil {
			node = child
		} else {
			return false
		}
	}

	return true
}
