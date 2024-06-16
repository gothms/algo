package main

type LinkedStack struct {
	h *ListNode
}

func (this *LinkedStack) Push(value int) {
	head := &ListNode{value, this.h}
	this.h = head
}

func (this *LinkedStack) Pop() int {
	if this.IsEmpty() {
		return -1
	}
	v := this.h.val
	this.h = this.h.next
	return v
}

func (this *LinkedStack) Peek() int {
	if this.IsEmpty() {
		return -1
	}
	return this.h.val
}

func (this *LinkedStack) IsEmpty() bool {
	return this.h == nil
}
