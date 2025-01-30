package stack

func ValidParenthesesStack(s string) bool {
	bracketsMap := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	st := make([]rune, 0, len(s))

	for _, t := range s {
		if v, ok := bracketsMap[t]; ok {
			// then it is a closing bracket
			if len(st) > 0 && v == st[len(st)-1] {
				// pop opening bracket from stack
				st = st[:len(st)-1]
			} else {
				return false
			}
		} else {
			st = append(st, t)
		}
	}

	if len(st) != 0 {
		return false
	}
	return true
}
