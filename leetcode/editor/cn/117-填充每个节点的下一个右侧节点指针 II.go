package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	var pre, head *Node
	next := func(r *Node) {
		if r == nil {
			return
		}
		if pre == nil {
			pre = r // 层的 head
		}
		if head != nil {
			head.Next = r // 层的链表
		}
		head = r
	}
	for lev := root; lev != nil; {
		pre, head = nil, nil
		for cur := lev; cur != nil; cur = cur.Next { // 层级遍历
			next(cur.Left)
			next(cur.Right)
		}
		lev = pre
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
