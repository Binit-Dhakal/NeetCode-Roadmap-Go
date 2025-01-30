package graph

func WordLadder(beginWord string, endWord string, wordList []string) int {
	patternMap := make(map[string][]string)

	addBegin := true
	endWordExists := false
	for _, word := range wordList {
		if word == endWord {
			endWordExists = true
		}

		if beginWord == word {
			addBegin = false
		}
	}

	if endWordExists == false {
		return 0
	}

	if addBegin == true {
		wordList = append(wordList, beginWord)
	}

	for _, word := range wordList {
		for i := range word {
			w := word
			w = w[:i] + "*" + w[i+1:]
			patternMap[w] = append(patternMap[w], word)
		}
	}

	// shortest path BFS
	queue := make([]string, 0) // toVisitWord, parentWord
	queue = append(queue, beginWord)
	visitedMap := make(map[string]struct{})

	res := 1
	for len(queue) > 0 {
		n := len(queue)
		for n != 0 {
			node := queue[0]
			queue = queue[1:]
			visitedMap[node] = struct{}{}
			if node == endWord {
				return res
			}

			for i := range node {
				w := node[:i] + "*" + node[i+1:]
				for _, neiWord := range patternMap[w] {
					if _, ok := visitedMap[neiWord]; !ok {
						visitedMap[neiWord] = struct{}{}
						queue = append(queue, neiWord)
					}
				}

			}
			n--
		}
		res++
	}
	return 0
}

// func WordLadderTwoBFS(beginWord string, endWord string, wordList []string) int {
// 	patternMap := make(map[string][]string)
//
// 	addBegin := true
// 	endWordExists := false
// 	for _, word := range wordList {
// 		if word == endWord {
// 			endWordExists = true
// 		}
//
// 		if beginWord == word {
// 			addBegin = false
// 		}
//
// 	}
//
// 	if endWordExists == false {
// 		return 0
// 	}
//
// 	if addBegin == true {
// 		wordList = append(wordList, beginWord)
// 	}
//
// 	for _, word := range wordList {
// 		for i := range word {
// 			w := word
// 			w = w[:i] + "*" + w[i+1:]
// 			patternMap[w] = append(patternMap[w], word)
// 		}
// 	}
//
// 	queue1 := []string{beginWord}
// 	queue2 := []string{endWord}
// 	fromBegin := map[string]int{beginWord: 1}
// 	fromEnd := map[string]int{endWord: 1}
//
// 	if len(queue1) > 0 && len(queue2) > 0 {
// 		if len(queue1) > len(queue2) {
// 			queue1, queue2 = queue2, queue1
// 		}
//
// 		queue1
// 	}
// }
