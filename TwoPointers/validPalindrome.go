package twopointers

import "strings"

func IsPalindromeNaive(s string) bool {
	isAlphaNumeric := func(s byte) bool {
		if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9') {
			return true
		}
		return false
	}

	changeCase := func(s byte) byte {
		if s >= 'A' && s <= 'Z' {
			return 'a' - 'A' + s
		}
		return s
	}

	leftPtr, rightPtr := 0, len(s)-1

	for leftPtr <= rightPtr {
		for leftPtr <= rightPtr && !isAlphaNumeric(s[leftPtr]) {
			leftPtr++
		}

		for leftPtr <= rightPtr && !isAlphaNumeric(s[rightPtr]) {
			rightPtr--
		}

		if leftPtr <= rightPtr && changeCase(s[leftPtr]) != changeCase(s[rightPtr]) {
			return false
		}

		leftPtr++
		rightPtr--
	}

	return true
}

func IsPalindromeRemoveNonAlphanumeric(s string) bool {
	isAlphaNumeric := func(s byte) bool {
		if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9') {
			return true
		}
		return false
	}

	changeCase := func(s byte) byte {
		if s >= 'A' && s <= 'Z' {
			return 'a' - 'A' + s
		}
		return s
	}
	var b strings.Builder
	for idx := range s {
		if isAlphaNumeric(s[idx]) {
			b.WriteByte(changeCase(s[idx]))
		}
	}

	s = b.String()

	leftPtr, rightPtr := 0, len(s)-1
	for leftPtr <= rightPtr {
		if s[leftPtr] != s[rightPtr] {
			return false
		}
		leftPtr++
		rightPtr--
	}

	return true
}
