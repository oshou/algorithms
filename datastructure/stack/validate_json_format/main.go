package main

import "fmt"

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() interface{} {
	if len(s.data) > 0 {
		v := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return v
	}
	return nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func validateFormat(chars string) bool {
	stack := Stack{}
	for _, char := range chars {
		switch char {
		case '[':
			stack.Push(']')
		case '(':
			stack.Push(')')
		case '{':
			stack.Push('}')
		case ']', ')', '}':
			if stack.IsEmpty() {
				return false
			}
			if stack.Pop() != char {
				return false
			}
		}
	}

	if stack.IsEmpty() {
		return true
	} else {
		return false
	}
}

func main() {
	chars1 := "{'key1': 'value1', 'key2': [1, 2, 3], 'key3': (1, 2, 3)}"
	chars2 := "[{'key1': 'value1', 'key2': [1, 2, 3], 'key3': (1, 2, 3)}"
	fmt.Println(validateFormat(chars1))
	fmt.Println(validateFormat(chars2))
}
