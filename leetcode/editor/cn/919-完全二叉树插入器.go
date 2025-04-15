package main

import (
	"fmt"
	"math/bits"
	"slices"
)

func main() {
	v := 13
	buf := []rune(fmt.Sprintf("%b\n", v))
	slices.Reverse(buf)
	fmt.Println(string(buf))
	var k int32 = 1<<bits.Len(uint(v)) - 1
	i := int32(bits.Reverse(uint(v))) & k
	fmt.Printf("%b\n", int(bits.Reverse(uint(v))))
	fmt.Println(k, i)
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
个人写法
lc 写法
	队列：bfs 将可新增子节点的节点加入一个队列，从左取，从右放
	二进制表示
*/

type CBTInserter struct {
	root *TreeNode
	cnt  int
}

func Constructor(root *TreeNode) CBTInserter {
	cnt := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		cnt++
		cur := q[0]
		q = q[1:]
		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}
	return CBTInserter{root, cnt}
}

func (this *CBTInserter) Insert(val int) int {
	this.cnt++
	n := this.cnt
	child := &TreeNode{Val: val}
	cur := this.root
	for i := bits.Len(uint(n)) - 2; i > 0; i-- {
		if n>>i&1 == 0 {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if n&1 == 0 {
		cur.Left = child
	} else {
		cur.Right = child
	}
	return cur.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
//leetcode submit region end(Prohibit modification and deletion)

// ==========个人写法==========
//type CBTInserter struct {
//	root         *TreeNode
//	leaf1, leaf0 []*TreeNode
//}
//
//func Constructor(root *TreeNode) CBTInserter {
//	leaf1, leaf0 := make([]*TreeNode, 0), make([]*TreeNode, 0)
//	mx := -1
//	var dfs func(*TreeNode, int)
//	dfs = func(root *TreeNode, level int) {
//		if root.Left == nil { // root.Right 必 nil
//			mx = max(mx, level)
//			if level == mx {
//				leaf0 = append(leaf0, root)
//			} else {
//				leaf1 = append(leaf1, root)
//			}
//			return
//		}
//		if root.Right == nil {
//			mx = max(mx, level+1) // leaf0 只有一个节点的情况
//			leaf0 = append(leaf0, root.Left)
//			leaf1 = append(leaf1, root)
//			return
//		}
//		level++
//		dfs(root.Left, level)
//		dfs(root.Right, level)
//	}
//	dfs(root, 0)
//	return CBTInserter{root: root, leaf0: leaf0, leaf1: leaf1}
//}
//
//func (this *CBTInserter) Insert(val int) int {
//	in := &TreeNode{Val: val}
//	var cur *TreeNode
//	if len(this.leaf1) == 0 {
//		this.leaf1, this.leaf0 = this.leaf0, this.leaf1
//	}
//	cur = this.leaf1[0]
//	if cur.Left == nil { // 插入左节点
//		cur.Left = in
//	} else { // 插入右节点
//		cur.Right = in
//		this.leaf1 = this.leaf1[1:] // 移除子节点
//	}
//	this.leaf0 = append(this.leaf0, in) // 新增子节点
//	return cur.Val
//}
//
//func (this *CBTInserter) Get_root() *TreeNode {
//	return this.root
//}
