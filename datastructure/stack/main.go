package main

import (
	"errors"
	"fmt"
)

type Stack interface {
	Push(v int)
	Pop() (int, error)
	IsEmpty() bool
}

type stack struct {
	data []int
	size int
}

func NewStack(cap int) *stack {
	d := make([]int, 0, cap)
	s := 0
	return &stack{data: d, size: s}
}

func (s *stack) Push(v int) {
	s.data = append(s.data, v)
	s.size += 1
}

func (s *stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	v := s.data[s.size-1]
	s.data = s.data[:s.size-1]
	s.size -= 1

	return v, nil
}

func (s *stack) IsEmpty() bool {
	return s.size == 0
}

func main() {
	s := NewStack(3)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for i := 1; i <= 4; i++ {
		v, err := s.Pop()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(v)
		}
	}
}
