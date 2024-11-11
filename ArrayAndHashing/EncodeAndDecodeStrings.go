package arrayandhashing

import (
	"fmt"
	"strings"
)

func Encode(strs []string) string {
	var sb strings.Builder

	for _, s := range strs {
		w := fmt.Sprintf("%d#%s", len(s), s)
		sb.WriteString(w)
	}

	return sb.String()
}

func Decode(strs string) []string {
	var results []string

	var length int
	for len(strs) != 0 {
		if _, err := fmt.Sscanf(strs, "%3d#", &length); err != nil {
			return results
		}

		prefix := fmt.Sprintf("%d#", length)
		res := strs[len(prefix) : len(prefix)+length]
		results = append(results, res)
		prefix = fmt.Sprintf("%d#%s", length, res)

		strs = strings.TrimPrefix(strs, prefix)

	}
	return results
}
