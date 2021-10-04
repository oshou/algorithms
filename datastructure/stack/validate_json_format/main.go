package main

import (
	"errors"
	"fmt"
)

// Input {'key1': 'value1', 'key2': [1, 2, 3], 'key3': (1, 2, 3)}  Output True
// Input [{'key1': 'value1', 'key2': [1, 2, 3], 'key3': (1, 2, 3)} Output False

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.data) != 0 {
		v := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return v, nil
	}
	return nil, errors.New("not found")
}

func (s *Stack) Peak() interface{} {
	if len(s.data) != 0 {
		return s.data[len(s.data)-1]
	}
	return 0
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func validateJSONFormat(stack Stack, chars string) bool {
	for _, c := range chars {
		switch c {
		case '{':
			stack.Push('}')
		case '[':
			stack.Push(']')
		case '(':
			stack.Push(')')
		default:
			if c == stack.Peak() {
				if _, err := stack.Pop(); err != nil {
					return false
				}
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
	stack := Stack{
		data: []interface{}{},
	}

	fmt.Println(validateJSONFormat(stack, chars1))
	fmt.Println(validateJSONFormat(stack, chars2))
}
