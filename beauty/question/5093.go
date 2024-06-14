package main

import "fmt"

func main() {
	t := NewTripleQueueInOne(1)
	t.Enqueue(0, 1)
	t.Enqueue(1, 2)
	t.Enqueue(2, 3)
	t.Enqueue(2, 4)
	t.Enqueue(1, 5)
	t.Enqueue(0, 6)
	v := t.Dequeue(0)
	fmt.Println(v)
	v = t.Dequeue(1)
	fmt.Println(v)
	v = t.Dequeue(2)
	fmt.Println(v)
}

type TripleQueueInOne struct {
	q   []int     // cap * 3
	ids [3][2]int // [2]int：h & t
	k   int       // cap
}

func NewTripleQueueInOne(queueSize int) *TripleQueueInOne {
	queueSize++
	return &TripleQueueInOne{q: make([]int, queueSize*3), k: queueSize}
}

func (t *TripleQueueInOne) Enqueue(queueID, value int) bool {
	if t.IsFull(queueID) {
		return false
	}
	is := t.ids[queueID]
	t.q[is[1]*3+queueID] = value
	t.ids[queueID][1] = (is[1] + 1) % t.k
	return true
}

func (t *TripleQueueInOne) Dequeue(queueID int) int {
	if t.IsEmpty(queueID) {
		return -1
	}
	is := t.ids[queueID]
	v := t.q[is[0]*3+queueID]
	t.ids[queueID][0] = (is[0] + 1) % t.k
	return v
}
func (t *TripleQueueInOne) IsFull(queueID int) bool {
	is := t.ids[queueID]
	return (is[1]+1)%t.k == is[0]
}
func (t *TripleQueueInOne) IsEmpty(queueID int) bool {
	is := t.ids[queueID]
	return is[0] == is[1]
}
