package main

func insert(head *ListNode, value int) *ListNode {
	if head == nil || head.val >= value {
		return &ListNode{value, head}
	}
	pre := head
	for pre.next != nil && pre.next.val < value {
		pre = pre.next
	}
	pre.next = &ListNode{value, pre.next}
	return head
}
