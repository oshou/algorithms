package main

import "fmt"

type Queue struct {
	data []interface{}
}

func (q *Queue) Enqueue(v interface{}) {
	q.data = append(q.data, v)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.data) != 0 {
		v := q.data[0]
		q.data = q.data[1:]
		return v
	}
	return nil
}

func main() {
	q := Queue{}

	q.Enqueue(1)
	fmt.Println(q.data)
	q.Enqueue(2)
	fmt.Println(q.data)
	q.Enqueue(3)
	fmt.Println(q.data)

	fmt.Println(q.Dequeue())
	fmt.Println(q.data)
	fmt.Println(q.Dequeue())
	fmt.Println(q.data)
	fmt.Println(q.Dequeue())
	fmt.Println(q.data)
	fmt.Println(q.Dequeue())
	fmt.Println(q.data)
}
