package main

import "fmt"

func main() {
	st := NewTripleStackInOne(1)
	st.Push(0, 1)
	st.Push(1, 2)
	st.Push(2, 3)
	pop := st.Pop(1)
	fmt.Println(pop)
}

type TripleStackInOne struct {
	st  []int
	ids [3]int
	cap int
}

func NewTripleStackInOne(stackSize int) *TripleStackInOne {
	return &TripleStackInOne{make([]int, stackSize*3), [3]int{}, stackSize}
}

func (ts *TripleStackInOne) Push(stackID, value int) bool {
	if ts.ids[stackID] == ts.cap {
		return false
	}
	ts.st[ts.ids[stackID]*3+stackID] = value
	ts.ids[stackID]++
	return true
}

func (ts *TripleStackInOne) Pop(stackID int) int {
	if ts.IsEmpty(stackID) {
		return -1
	}
	v := ts.st[(ts.ids[stackID]-1)*3+stackID]
	ts.ids[stackID]--
	return v
}

func (ts *TripleStackInOne) Peek(stackID int) int {
	if ts.IsEmpty(stackID) {
		return -1
	}
	return ts.st[(ts.ids[stackID]-1)*3+stackID]
}

func (ts *TripleStackInOne) IsEmpty(stackID int) bool {
	return ts.ids[stackID] == 0
}
