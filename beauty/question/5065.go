package main

import "fmt"

func main() {
	sts := NewSetOfStacks(2)
	sts.Push(1)
	sts.Push(2)
	sts.Push(3)
	pop := sts.Pop()
	fmt.Println(pop)
	pop = sts.Pop()
	fmt.Println(pop)
	pop = sts.Pop()
	fmt.Println(pop)
}

type SetOfStacks struct {
	sts [][]int
	k   int
}

func NewSetOfStacks(sizePerStack int) *SetOfStacks {
	return &SetOfStacks{make([][]int, 0), sizePerStack}
}

func (s *SetOfStacks) Push(value int) {
	var lastSt []int
	if s.IsEmpty() { // 总栈为空
		lastSt = make([]int, 0, s.k)
		s.sts = append(s.sts, lastSt)
	} else if lastSt = s.sts[len(s.sts)-1]; len(lastSt) == s.k { // 栈已满
		lastSt = make([]int, 0, s.k)
		s.sts = append(s.sts, lastSt)
	}
	s.sts[len(s.sts)-1] = append(lastSt, value)
}

func (s *SetOfStacks) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	lastSt := &s.sts[len(s.sts)-1]       // 最后一个栈
	v := (*lastSt)[len(*lastSt)-1]       // 最后一个元素
	*lastSt = (*lastSt)[:len(*lastSt)-1] // 删除元素
	if len(*lastSt) == 0 {               // 栈已空
		s.sts = s.sts[:len(s.sts)-1]
	}
	return v
}

func (s *SetOfStacks) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	lastSt := s.sts[len(s.sts)-1]
	return lastSt[len(lastSt)-1]
}

func (s *SetOfStacks) IsEmpty() bool {
	return len(s.sts) == 0
}
