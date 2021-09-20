package main

import (
	"container/heap"
	"fmt"
)

type PriorityQueue interface {
	Push(v interface{})
	Pop() interface{}
}

type Item struct {
	priority int
	value    interface{}
}

type PriorityQueueImpl struct {
	data []*Item
}

func NewPriorityQueue(cap int) *PriorityQueueImpl {
	return &PriorityQueueImpl{
		data: []*Item{},
	}
}

func (pq *PriorityQueueImpl) Len() int           { return len(pq.data) }
func (pq *PriorityQueueImpl) Less(i, j int) bool { return pq.data[i].priority < pq.data[j].priority }
func (pq *PriorityQueueImpl) Swap(i, j int)      { pq.data[i], pq.data[j] = pq.data[j], pq.data[i] }

func (pq *PriorityQueueImpl) Push(v interface{}) {
	item := v.(Item)
	i := &Item{
		priority: item.priority,
		value:    item.value,
	}
	pq.data = append(pq.data, i)
}

func (pq *PriorityQueueImpl) Pop() interface{} {
	old := pq.data
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	pq.data = old[:n-1]
	return item
}

func main() {
	pq := NewPriorityQueue(100)
	pq.data = []*Item{
		{priority: 2, value: 100},
		{priority: 1, value: 500},
		{priority: 5, value: 20},
	}
	heap.Init(pq)
	item := Item{
		priority: 0,
		value:    1000,
	}
	heap.Push(pq, item)
	fmt.Printf("minimum: %d\n", (*pq).data[0])
	for pq.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(pq))
	}
}
