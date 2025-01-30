package dp1

func WordBreakNaive(s string, wordDict []string) bool {
	var dfs func(t string) bool

	dfs = func(t string) bool {
		if t == "" {
			return true
		}
		ptr := 0
		for _, word := range wordDict {
			if len(word) > len(t) {
				continue
			}
			flag := true
			for i := 0; i < len(word); i++ {
				if word[i] != t[ptr] {
					flag = false
					break
				} else {
					ptr++
				}
			}

			if flag == true && dfs(t[ptr:]) {
				return true
			}
			ptr = 0
		}

		return false
	}

	return dfs(s)
}

func WordBreakTopDown(s string, wordDict []string) bool {
	// time limit exceeded
	dp := make(map[int]bool)
	dp[len(s)] = true

	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if val, ok := dp[idx]; ok {
			return val
		}

		for _, word := range wordDict {
			n := len(word)
			if idx+n <= len(s) && s[idx:idx+n] == word && dfs(idx+n) {
				dp[idx] = true
				return true
			}
		}
		dp[idx] = false
		return false
	}
	return dfs(0)
}

func WordBreakTopDownHashSet(s string, wordDict []string) bool {
	dp := make(map[int]bool)
	dp[len(s)] = true

	wordSet := make(map[string]struct{})
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}

	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if val, ok := dp[idx]; ok {
			return val
		}
		for i := idx; i < len(s); i++ {
			if _, ok := wordSet[s[idx:i+1]]; ok && dfs(i+1) {
				dp[idx] = true
				return true
			}
		}
		dp[idx] = false
		return false
	}
	return dfs(0)
}

func WordBreakBottomUp(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true

	for i := len(s) - 1; i >= 0; i-- {
		for _, w := range wordDict {
			if i+len(w) <= len(s) && s[i:i+len(w)] == w {
				dp[i] = dp[i+len(w)]
			}
			if dp[i] {
				break
			}
		}
	}

	return dp[0]
}

type Trie struct {
	children [26]*Trie
	word     bool
}

func ConstructTrie() *Trie {
	return &Trie{
		children: [26]*Trie{},
		word:     false,
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = ConstructTrie()
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

func WordBreakTrie(s string, wordDict []string) bool {
	trie := ConstructTrie()
	maxLength := 0
	for _, word := range wordDict {
		maxLength = max(maxLength, len(word))
		trie.Insert(word)
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[n] = true

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s) && j < i+maxLength; j++ {
			if trie.Search(s[i : j+1]) {
				dp[i] = dp[j+1]
				if dp[i] {
					break
				}

			}
		}
	}
	return dp[0]
}
