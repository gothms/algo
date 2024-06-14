package main

func removeElements(head *ListNode, val int) *ListNode {
	sentinel := &ListNode{0, head}
	for pre, cur := sentinel, head; cur != nil; cur = cur.next {
		if cur.val == val {
			pre.next = cur.next
		} else {
			pre = pre.next
		}
	}
	return sentinel.next
}
