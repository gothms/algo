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

}

// leetcode submit region end(Prohibit modification and deletion)

func flatten_(root *Node) *Node {
	var dfs func(*Node) *Node
	dfs = func(root *Node) *Node { // 返回尾结点
		var next, tail *Node
		for root != nil {
			next = root.Next
			if root.Child != nil {
				tail = dfs(root.Child)                                         // dfs Child
				root.Next, root.Child.Prev, root.Child = root.Child, root, nil // Child 为 Next
				if next != nil {                                               // root.Next 有可能为 nil
					tail.Next, next.Prev = next, tail
				}
			} else {
				tail = root // 记录上一个节点
			}
			root = next // 提前记录了 root.Next
		}
		return tail
	}
	dfs(root)
	return root
}
