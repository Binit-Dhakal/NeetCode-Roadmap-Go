package main

import (
	"fmt"
	"sort"
)

func sortString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}

func groupAnagrams(strs []string) [][]string {
	sorted_strs := make(map[string][]string)

	for _, str := range strs {
		sorted_str := sortString(str)
		sorted_strs[sorted_str] = append(sorted_strs[sorted_str], str)
	}

	var output [][]string

	for _, v := range sorted_strs {
		output = append(output, v)
	}

	return output
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
