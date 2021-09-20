package main

import (
	"errors"
	"fmt"
)

type Stack interface {
	Push(v interface{})
	Pop() (interface{}, error)
	IsEmpty() bool
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
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return v, nil
}

func (s *StackImpl) IsEmpty() bool {
	return len(s.data) == 0
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
