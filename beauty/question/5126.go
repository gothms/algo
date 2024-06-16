package main

import "fmt"

func reversePrint(head *ListNode) {
	var dfs func(*ListNode)
	dfs = func(cur *ListNode) {
		if cur == nil {
			return
		}
		dfs(cur.next)
		cur.print()
	}
	dfs(head)
}
func (node *ListNode) print() {
	fmt.Print(node.val, " ")
}
