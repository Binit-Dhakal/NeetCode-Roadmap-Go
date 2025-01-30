package tries

type TrieNode struct {
	children [26]*TrieNode
	word     bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

type WordDictionary struct {
	root *TrieNode
}

func ConstructWD() WordDictionary {
	return WordDictionary{root: NewTrieNode()}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for _, ch := range word {
		idx := ch - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = NewTrieNode()
		}
		cur = cur.children[idx]
	}
	cur.word = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.dfs(word, 0, this.root)
}

func (this *WordDictionary) dfs(word string, jdx int, root *TrieNode) bool {
	cur := root
	for idx := jdx; idx < len(word); idx++ {

		if word[idx] == '.' {
			// we have to do dfs
			for j := 0; j < 26; j++ {
				if cur.children[j] != nil && this.dfs(word, idx+1, cur.children[j]) {
					return true
				}
			}

			return false
		}

		if cur.children[word[idx]-'a'] == nil {
			return false
		}
		cur = cur.children[word[idx]-'a']

	}
	return cur.word
}
