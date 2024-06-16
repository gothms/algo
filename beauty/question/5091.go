package main

import "fmt"

func main() {
	q := newLinkedQueue()
	q.enqueue(3)
	q.enqueue(2)
	q.enqueue(1)
	dequeue := q.dequeue()
	dequeue = q.dequeue()
	dequeue = q.dequeue()
	dequeue = q.dequeue()
	q.enqueue(4)
	fmt.Println(dequeue)
}

type ListNode struct {
	val  int
	next *ListNode
}

type LinkedQueue struct {
	h, t *ListNode
}

func newLinkedQueue() *LinkedQueue {
	l := &ListNode{}
	l.next = l
	return &LinkedQueue{l, l}
}

func (q *LinkedQueue) enqueue(value int) {
	q.t.next = &ListNode{value, q.h}
	q.t = q.t.next
}

func (q *LinkedQueue) dequeue() int {
	if q.isEmpty() {
		return -1
	}
	cur := q.h.next
	if cur == q.t {
		q.t = q.h
	}
	q.h.next, cur.next = cur.next, nil
	return cur.val
}

func (q *LinkedQueue) isEmpty() bool {
	return q.h == q.t
}

//type LinkedQueue struct {
//	queue []int
//}
//
//func newLinkedQueue() *LinkedQueue {
//	return &LinkedQueue{make([]int, 0)}
//}
//
//func (q *LinkedQueue) enqueue(value int) {
//	q.queue = append(q.queue, value)
//}
//
//func (q *LinkedQueue) dequeue() int {
//	if q.isEmpty() {
//		return -1
//	}
//	v := q.queue[0]
//	q.queue = q.queue[1:]
//	return v
//}
//
//func (q *LinkedQueue) isEmpty() bool {
//	return len(q.queue) == 0
//}
