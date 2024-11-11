package arrayandhashing

import "sort"

func GroupAnagramsArray(strs []string) [][]string {
	strsMap := make(map[[26]byte][]string)

	for _, str := range strs {
		count := [26]byte{}
		for _, s := range str {
			count[s-'a']++
		}
		strsMap[count] = append(strsMap[count], str)
	}

	var result [][]string
	for _, str := range strsMap {
		result = append(result, str)
	}
	return result
}

func GroupAnagramsSort(strs []string) [][]string {
	type Value struct {
		originalStr string
		sortedStr   string
	}
	var values []Value

	for _, str := range strs {
		s := []rune(str)
		sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })

		value := Value{
			originalStr: str,
			sortedStr:   string(s),
		}

		values = append(values, value)
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i].sortedStr < values[j].sortedStr
	})

	result := [][]string{}
	currentStr := values[0].sortedStr
	res := []string{}
	for _, val := range values {

		if val.sortedStr != currentStr {
			result = append(result, res)
			currentStr = val.sortedStr
			res = []string{}
		}

		res = append(res, val.originalStr)
	}
	if len(res) != 0 {
		result = append(result, res)
	}
	return result
}
