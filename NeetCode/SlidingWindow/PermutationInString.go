package slidingwindow

func PermutationInString(s1 string, s2 string) bool {
	windowSize := len(s1)

	s1Array := [26]int{}
	s2Array := [26]int{}

	for idx := range s1 {
		s1Array[s1[idx]-'a']++
	}

	for idx := 0; idx < len(s2); idx++ {
		s2Array[s2[idx]-'a']++

		if idx >= windowSize {
			index := idx - windowSize
			s2Array[s2[index]-'a']--
		}

		if s1Array == s2Array {
			return true
		}

	}
	return false
}
