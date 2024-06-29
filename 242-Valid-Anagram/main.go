package main

import "fmt"

func usingHashmap(s string, t string) bool {
	// Space - O(n) and time - O(n)
	// It evens works for unicode

	if len(s) != len(t) {
		return false
	}

	set := make(map[rune]int)

	for _, char := range s {
		_, ok := set[char]

		if !ok {
			set[char] = 1
		} else {
			set[char]++
		}
	}

	for _, char := range t {
		_, ok := set[char]
		if !ok {
			return false
		} else {
			set[char]--
		}
	}

	// now check if all the keys in the map has 0 value
	for k := range set {
		if set[k] != 0 {
			return false
		}
	}
	return true
}

func usingArray(s string, t string) bool {
	// Space: O(n), Time: O(n)
	// It uses the fact that both string have only lower case letter
	arr := [26]int{}

	for _, char := range s {
		arr[int(char-'a')]++
	}

	for _, char := range t {
		arr[int(char-'a')]--
	}

	for _, v := range arr {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(usingHashmap("anagram", "nagrama"))

}
