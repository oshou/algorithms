package main

import (
	"errors"
	"fmt"
)

type Queue interface {
	Enqueue(v int)
	Dequeue() (int, error)
	IsEmpty() bool
}

type queue struct {
	data []int
	size int
}

func NewQueue(cap int) *queue {
	d := make([]int, 0, cap)
	s := 0
	return &queue{data: d, size: s}
}

func (q *queue) Enqueue(v int) {
	q.data = append(q.data, v)
	q.size += 1
}

func (q *queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	v := q.data[0]
	q.data = q.data[1:]
	q.size -= 1

	return v, nil
}

func (q *queue) IsEmpty() bool {
	return q.size == 0
}

func main() {
	q := NewQueue(3)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	for i := 1; i <= 4; i++ {
		v, err := q.Dequeue()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(v)
		}
	}
}
