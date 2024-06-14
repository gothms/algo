package main

type ArrayStack struct {
	st []int
}

func newArrayStack(stackSize int) *ArrayStack {
	return &ArrayStack{make([]int, 0, stackSize)}
}

func (this *ArrayStack) Push(value int) bool {
	if len(this.st) == cap(this.st) {
		return false
	}
	this.st = append(this.st, value)
	return true
}

func (this *ArrayStack) Pop() int {
	if this.IsEmpty() {
		return -1
	}
	v := this.st[len(this.st)-1]
	this.st = this.st[:len(this.st)-1]
	return v
}

func (this *ArrayStack) Peek() int {
	if this.IsEmpty() {
		return -1
	}
	return this.st[len(this.st)-1]
}

func (this *ArrayStack) IsEmpty() bool {
	return len(this.st) == 0
}
