package main

type MultiSet struct {
	items map[int]int
}

func NewMultiSet() *MultiSet {
	return &MultiSet{
		items: make(map[int]int),
	}
}

func (s *MultiSet) AddNums(nums []int) {
	for _, num := range nums {
		s.AddNum(num)
	}
}

func (s *MultiSet) AddNum(num int) {
	s.items[num]++
}

func (s *MultiSet) OccuranceOf(num int) int {
	occ, ok := s.items[num]
	if !ok {
		return 0
	}
	return occ
}

func (s *MultiSet) HasDuplicate(num int) bool {
	occ := s.OccuranceOf(num)
	return occ > 1
}
