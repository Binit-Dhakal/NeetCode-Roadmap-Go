package stack

type Stack struct {
	datas []int
}

func (s *Stack) Pop() int {
	val := s.datas[len(s.datas)-1]
	s.datas = s.datas[:len(s.datas)-1]
	return val
}

func (s *Stack) Push(elem int) {
	s.datas = append(s.datas, elem)
}

func (s *Stack) Peek() int {
	val := s.datas[len(s.datas)-1]
	return val
}

func (s *Stack) Length() int {
	return len(s.datas)
}

func DailyTemperatures(temperatures []int) []int {
	s := new(Stack)
	result := make([]int, len(temperatures))

	for idx, tempt := range temperatures {
		for s.Length() > 0 {
			if stack_idx := s.Peek(); temperatures[stack_idx] < tempt {
				stack_idx := s.Pop()
				result[stack_idx] = idx - stack_idx
			} else {
				break
			}
		}
		s.Push(idx)
	}
	return result
}
