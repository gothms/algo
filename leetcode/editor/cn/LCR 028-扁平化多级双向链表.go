package main

import "fmt"

func main() {
	root := &Node{Val: 1}
	root.Child = &Node{Val: 2, Child: &Node{Val: 3}}
	n := flatten(root)
	fmt.Println(n)
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
	dfs = func(root *Node) *Node {
		pre := root
		for cur := root; cur != nil; {
			if cur.Child != nil {
				end := dfs(cur.Child)
				// 插入节点，且 pre cur cur.Child 赋予新值
				cur.Next, cur.Child.Prev, end.Next, pre, cur, cur.Child =
					cur.Child, cur, cur.Next, end, cur.Next, nil // 此时 cur 可能已为 nil
				if cur != nil {
					cur.Prev = end
				}
			} else {
				pre, cur = cur, cur.Next
			}
		}
		return pre // 返回该层级的最后一个结点
	}
	dfs(root)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
