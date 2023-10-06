package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{6, nil}}}}
	//head = &ListNode{2, &ListNode{4, nil}}
	contruct := reContruct(head)
	for ; contruct != nil; contruct = contruct.Next {
		fmt.Println(contruct.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reContruct(head *ListNode) *ListNode {
	pre := &ListNode{Next: head}
	for cur := pre; cur.Next != nil; {
		if cur.Next.Val&1 == 0 {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return pre.Next
}
