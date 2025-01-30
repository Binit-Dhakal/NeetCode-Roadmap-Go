package stack

import (
	"strconv"
	"strings"
)

type stack struct {
	datas []string
}

func (s *stack) Push(data string) {
	s.datas = append(s.datas, data)
}

func (s *stack) Pop() string {
	data := s.datas[len(s.datas)-1]
	s.datas = s.datas[:len(s.datas)-1]
	return data
}

func EvalRPNStack(tokens []string) int {
	stack := stack{}
	operators := "+-*/"

	var operand1, operand2 int
	for _, token := range tokens {
		if !strings.Contains(operators, token) {
			stack.Push(token)
			continue
		}
		operand1, _ = strconv.Atoi(stack.Pop())
		operand2, _ = strconv.Atoi(stack.Pop())

		var res int
		switch token {
		case "+":
			res = operand1 + operand2
		case "-":
			res = operand2 - operand1
		case "*":
			res = operand2 * operand1
		case "/":
			res = operand2 / operand1
		}

		stack.Push(strconv.Itoa(res))
	}
	val, _ := strconv.Atoi(stack.Pop())
	return val
}

func EvalRPNRecursion(tokens []string) int {
	index := len(tokens) - 1

	var dfs func() int

	dfs = func() int {
		token := tokens[index]
		index--

		if !strings.Contains("+-*/", token) {
			v, _ := strconv.Atoi(token)
			return v
		}
		right := dfs()
		left := dfs()

		switch token {
		case "+":
			return left + right
		case "-":
			return left - right
		case "*":
			return left * right
		default:
			return left / right
		}
	}

	return dfs()
}
