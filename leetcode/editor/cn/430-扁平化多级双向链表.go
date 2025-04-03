package main

import (
	"fmt"
)

func main() {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n1.Next, n1.Child = n2, n3
	n2.Prev = n1
	//n1.Child, n2.Child = n2, n3
	node := flatten(n1)
	for cur := node; cur != nil; cur = cur.Next {
		fmt.Println(cur.Val)
	}
	fmt.Println(n1.Prev)

	//lc := lc(n1)
	//fmt.Println(node == lc)
	//
	//equal := reflect.DeepEqual(node, lc)
	//fmt.Println(equal)
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	var dfs func(*Node) *Node
	dfs = func(r *Node) (tail *Node) { // 返回尾结点
		var next *Node
		for r != nil {
			if next = r.Next; r.Child != nil {
				tail = dfs(r.Child)                             // dfs Child
				r.Next, r.Child.Prev, r.Child = r.Child, r, nil // Child 为 Next
				if next != nil {                                // r.Next 有可能为 nil
					tail.Next, next.Prev = next, tail
				}
			} else {
				tail = r // 记录上一个节点
			}
			r = next // 提前记录了 r.Next
		}
		return
	}
	dfs(root)
	return root
}

// leetcode submit region end(Prohibit modification and deletion)
