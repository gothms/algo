package main

type LinkedQueue struct {
	queue []int
}

func newLinkedQueue() *LinkedQueue {
	return &LinkedQueue{make([]int, 0)}
}

func (q *LinkedQueue) enqueue(value int) {
	q.queue = append(q.queue, value)
}

func (q *LinkedQueue) dequeue() int {
	if q.isEmpty() {
		return -1
	}
	v := q.queue[0]
	q.queue = q.queue[1:]
	return v
}

func (q *LinkedQueue) isEmpty() bool {
	return len(q.queue) == 0
}
