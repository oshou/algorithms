package main

import (
	"errors"
	"fmt"
)

type Queue interface {
	Enqueue(v interface{})
	Dequeue() (interface{}, error)
	IsEmpty() bool
}

type QueueImpl struct {
	data []interface{}
}

func NewQueue(cap int) *QueueImpl {
	d := make([]interface{}, 0, cap)
	return &QueueImpl{
		data: d,
	}
}

func (q *QueueImpl) Enqueue(v interface{}) {
	q.data = append(q.data, v)
}

func (q *QueueImpl) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	v := q.data[0]
	q.data = q.data[1:]

	return v, nil
}

func (q *QueueImpl) IsEmpty() bool {
	return len(q.data) == 0
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
			fmt.Println("dequeue:", v)
		}
	}
}
