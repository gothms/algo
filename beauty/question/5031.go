package main

type ListNode struct {
	val  int
	next *ListNode
}

func insertAtHead(head *ListNode, value int) *ListNode {
	return &ListNode{value, head}
}

func insertAtTail(head *ListNode, value int) *ListNode {
	if head == nil {
		return &ListNode{val: value}
	}
	pre := head
	for pre.next != nil {
		pre = pre.next
	}
	pre.next = &ListNode{value, pre.next}
	return head
}

func size(head *ListNode) int {
	ans := 0
	for cur := head; cur != nil; cur = cur.next {
		ans++
	}
	return ans
}

func prev(head *ListNode, enode *ListNode) *ListNode {
	var pre *ListNode
	for cur := head; cur != nil && cur != enode; cur = cur.next {
		pre = cur
	}
	return pre
}

func delete(head *ListNode, enode *ListNode) *ListNode {
	if enode == head { // fast path
		newHead := enode.next
		enode.next = nil
		return newHead
	}
	pre := prev(head, enode)
	if pre != nil {
		pre.next, enode.next = enode.next, nil
	}
	return head
}
