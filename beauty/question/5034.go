package main

import "fmt"

func main() {
	list := Constructor()
	list.add2Tail(1)
	list.add2Tail(2)
	list.add2Tail(3)
	list.add2Tail(4)
	list.add(2, 5)
	list.remove(4)
	get := list.get(0)
	fmt.Println(get)
}

type ListNode struct {
	val  int
	next *ListNode
}
type LinkedList struct {
	h, t *ListNode
	k    int
}

func Constructor() LinkedList {
	var h ListNode
	return LinkedList{&h, &h, 0}
}

func (list *LinkedList) add2Tail(value int) {
	list.t.next = &ListNode{val: value}
	list.t = list.t.next
	list.k++
}

func (list *LinkedList) add(index int, value int) bool {
	if index > list.k { // fast fail
		return false
	}
	if index == list.k { // fast path
		list.add2Tail(value)
	} else {
		pre := list.h
		for i := 0; i < index; i++ {
			pre = pre.next
		}
		pre.next = &ListNode{value, pre.next}
		list.k++
	}
	return true
}

func (list *LinkedList) remove(index int) int {
	if index >= list.k {
		return -1
	}
	pre := list.h
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	ans := pre.next.val
	pre.next, pre.next.next = pre.next.next, nil
	if index == list.k-1 { // 修正 t
		list.t = pre
	}
	list.k--
	return ans
}

func (list *LinkedList) get(index int) int {
	if index >= list.k {
		return -1
	}
	if index == list.k-1 { // fast path
		return list.t.val
	}
	cur := list.h.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.val
}
