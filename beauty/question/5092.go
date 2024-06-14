package main

import "fmt"

func main() {
	cq := NewCircularQueue(1)
	b := cq.Enqueue(2)
	fmt.Println(b)
	b = cq.Enqueue(3)
	fmt.Println(b)
	d := cq.Dequeue()
	fmt.Println(d)
	d = cq.Dequeue()
	fmt.Println(d)
}

type CircularQueue struct {
	q    []int
	h, t int
	k    int
}

func NewCircularQueue(queueSize int) *CircularQueue {
	queueSize++
	return &CircularQueue{q: make([]int, queueSize), k: queueSize}
}

func (cq *CircularQueue) Enqueue(item int) bool {
	if cq.IsFull() {
		return false
	}
	cq.q[cq.t] = item
	cq.t = (cq.t + 1) % cq.k
	return true
}

func (cq *CircularQueue) Dequeue() int {
	if cq.IsEmpty() {
		return -1
	}
	v := cq.q[cq.h]
	cq.h = (cq.h + 1) % cq.k
	return v
}
func (cq *CircularQueue) IsFull() bool {
	return (cq.t+1)%cq.k == cq.h
}
func (cq *CircularQueue) IsEmpty() bool {
	return cq.h == cq.t
}
