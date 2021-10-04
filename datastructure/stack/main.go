package main

import (
	"errors"
	"fmt"
)

type Stack interface {
	Push(v interface{})
	Pop() (interface{}, error)
}

type StackImpl struct {
	data []interface{}
}

func NewStack(cap int) *StackImpl {
	d := make([]interface{}, 0, cap)
	return &StackImpl{
		data: d,
	}
}

func (s *StackImpl) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *StackImpl) Pop() (interface{}, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return v, nil
}

func (s *StackImpl) isEmpty() bool {
	return len(s.data) == 0
}

func main() {
	s := NewStack(3)
	s.Push(1)
	fmt.Println(s.data)
	s.Push(2)
	fmt.Println(s.data)
	s.Push(3)
	fmt.Println(s.data)

	s.Pop()
	fmt.Println(s.data)
	s.Pop()
	fmt.Println(s.data)
	s.Pop()
	fmt.Println(s.data)
	s.Pop()
	fmt.Println(s.data)
}
