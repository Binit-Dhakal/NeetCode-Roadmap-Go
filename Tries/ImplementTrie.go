package tries

type Trie struct {
	children [26]*Trie
	word     bool
}

func Constructor() Trie {
	return Trie{
		children: [26]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = &Trie{}
		}

		cur = cur.children[idx]
	}
	cur.word = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if cur.children[idx] == nil {
			return false
		}

		cur = cur.children[idx]
	}
	return cur.word
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		idx := prefix[i] - 'a'
		if cur.children[idx] == nil {
			return false
		}
		cur = cur.children[idx]
	}
	return true
}
