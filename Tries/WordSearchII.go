package tries

func WordSearchIIBackTracking(board [][]byte, words []string) []string {
	res := []string{}

	var findWords func(word string, idx int, r, c int) bool
	findWords = func(word string, idx int, r, c int) bool {
		if idx >= len(word) {
			return true
		}

		if r >= len(board) || r < 0 || c < 0 || c >= len(board[0]) || board[r][c] != word[idx] {
			return false
		}

		board[r][c] = '*'

		defer func() { board[r][c] = word[idx] }()

		return findWords(word, idx+1, r, c-1) || findWords(word, idx+1, r, c+1) || findWords(word, idx+1, r-1, c) || findWords(word, idx+1, r+1, c)
	}
	for _, word := range words {

		var valid bool
		for r := range board {
			if valid == true {
				break
			}
			for c := range board[r] {
				if board[r][c] == word[0] && findWords(word, 0, r, c) {
					valid = true
					break
				}
			}
		}

		if valid {
			res = append(res, word)
		}
	}
	return res
}

type WTrieNode struct {
	children [26]*WTrieNode
	idx      int
	refs     int // to delete node already used
}

func NewWTrieNode() *WTrieNode {
	return &WTrieNode{idx: -1}
}

func (w *WTrieNode) Insert(word string, idx int) {
	cur := w
	cur.refs++
	for _, ch := range word {
		idx := ch - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = NewWTrieNode()
		}
		cur = cur.children[idx]
		cur.refs++
	}
	cur.idx = idx
}

func (w *WTrieNode) Delete(word string) {
	cur := w
	cur.refs--
	for _, ch := range word {
		idx := ch - 'a'
		next := cur.children[idx]
		if next == nil {
			return
		}
		next.refs--
		if next.refs == 0 {
			next = nil
			break
		}
		cur = next
	}
}

func WordSearchIITrie(board [][]byte, words []string) []string {
	trie := NewWTrieNode()
	for idx, word := range words {
		trie.Insert(word, idx)
	}

	var result []string

	var findWord func(r int, c int, root *WTrieNode)
	findWord = func(r, c int, root *WTrieNode) {
		if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) || board[r][c] == '*' || root.children[board[r][c]-'a'] == nil {
			return
		}

		node := root.children[board[r][c]-'a']
		if node.idx != -1 {
			word := words[node.idx]
			result = append(result, word)
			node.idx = -1
			trie.Delete(word)

		}

		tmp := board[r][c]
		board[r][c] = '*'

		findWord(r, c+1, node)
		findWord(r, c-1, node)
		findWord(r-1, c, node)
		findWord(r+1, c, node)
		board[r][c] = tmp
	}

	for r := range board {
		for c := range board[r] {
			findWord(r, c, trie)
		}
	}
	return result
}
