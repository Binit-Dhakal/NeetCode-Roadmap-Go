package arrayandhashing

func ContainsDuplicate(nums []int) bool {
	// struct{} with no fields occupy zero byte of memory
	hashset := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := hashset[n]; ok {
			return true
		}
		hashset[n] = struct{}{}
	}
	return false
}
