package main

import "fmt"

type Stack struct {
	data []int32
}

func (s *Stack) Push(v int32) {
	s.data = append(s.data, v)
}

func (s *Stack) Peek() (int32, bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	length := len(s.data)
	return s.data[length-1], true
}

func (s *Stack) Pop() (int32, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	length := len(s.data)
	value := s.data[length-1]
	s.data = s.data[:length-1]
	return value, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func NewStack() *Stack {
	return &Stack{data: make([]int32, 0)}
}

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*/
func isValid(s string) bool {
	valMap := map[int32]int32{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := NewStack()
	for _, v := range s {
		val, exists := stack.Peek()
		if !exists {
			stack.Push(v)
		} else {
			if valMap[val] == v {
				stack.Pop()
			} else {
				stack.Push(v)
			}
		}
	}

	return stack.IsEmpty()
}

func main() {
	fmt.Println(isValid("([])"))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("([]){"))
}
