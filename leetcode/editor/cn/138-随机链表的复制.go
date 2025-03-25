package main

import "fmt"

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	// 非 hash
	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
	}
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}
	pre := &Node{Next: head}
	for cur := pre; cur.Next != nil; cur = cur.Next {
		cur.Next, cur.Next.Next = cur.Next.Next, cur.Next.Next.Next // 新的链表 & 恢复原链表
	}
	return pre.Next

	// hash
	//pre := &Node{}
	//nodes := make([]*Node, 0)
	//cur, newCur := head, pre
	//for i := 0; cur != nil; i++ {
	//	newCur.Next = &Node{Val: cur.Val}
	//	cur, newCur, cur.Val = cur.Next, newCur.Next, i
	//	nodes = append(nodes, newCur)
	//}
	//for _, node := range nodes {
	//	if head.Random != nil {
	//		node.Random = nodes[head.Random.Val]
	//	}
	//	head = head.Next
	//}
	//return pre.Next
}

//leetcode submit region end(Prohibit modification and deletion)
