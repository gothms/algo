package main

import "fmt"

func main() {
	this := newArrayStack(1)
	this.Push(1)
	this.Push(2)
	this.Push(3)
	this.Push(4)
	pop := this.Pop()
	fmt.Println(pop)
	this.Push(5)
	pop = this.Pop()
	fmt.Println(pop)
	push := this.Push(1)
	fmt.Println(push)
}

type ArrayStack struct {
	st  []int
	top int
}

func newArrayStack(stackSize int) *ArrayStack {
	return &ArrayStack{make([]int, stackSize), 0}
}

func (this *ArrayStack) Push(value int) bool {
	if this.top == len(this.st) {
		return false
	}
	this.st[this.top] = value
	this.top++
	return true
}

func (this *ArrayStack) Pop() int {
	if this.IsEmpty() {
		return -1
	}
	this.top--
	return this.st[this.top]
}

func (this *ArrayStack) Peek() int {
	if this.IsEmpty() {
		return -1
	}
	return this.st[this.top-1]
}

func (this *ArrayStack) IsEmpty() bool {
	return this.top == 0
}

//type ArrayStack struct {
//	st []int
//}
//
//func newArrayStack(stackSize int) *ArrayStack {
//	return &ArrayStack{make([]int, 0, stackSize)}
//}
//
//func (this *ArrayStack) Push(value int) bool {
//	if len(this.st) == cap(this.st) {
//		return false
//	}
//	this.st = append(this.st, value)
//	return true
//}
//
//func (this *ArrayStack) Pop() int {
//	if this.IsEmpty() {
//		return -1
//	}
//	v := this.st[len(this.st)-1]
//	this.st = this.st[:len(this.st)-1]
//	return v
//}
//
//func (this *ArrayStack) Peek() int {
//	if this.IsEmpty() {
//		return -1
//	}
//	return this.st[len(this.st)-1]
//}
//
//func (this *ArrayStack) IsEmpty() bool {
//	return len(this.st) == 0
//}
