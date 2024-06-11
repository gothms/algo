package main

import "fmt"

func main() {
	d := [][]int{{85, 82, 1},
		{74, 85, 1},
		{39, 70, 0},
		{82, 38, 1},
		{74, 39, 0},
		{39, 13, 1}}
	tree := createBinaryTree(d)
	vs, q := make([]int, 0, len(d)+1), []*TreeNode{tree}
	for ; len(q) > 0; q = q[1:] {
		cur := q[0]
		vs = append(vs, cur.Val)
		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}
	fmt.Println(vs) // [74,85,39,82,null,13,70,38]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
func createBinaryTree(descriptions [][]int) *TreeNode {
	memo := make(map[int]*TreeNode, len(descriptions)+1)
	for _, d := range descriptions {
		p, c := memo[d[0]], memo[d[1]]
		if p == nil {
			p = &TreeNode{Val: d[0]}
			memo[d[0]] = p
		}
		if c == nil {
			c = &TreeNode{Val: d[1]}
			memo[d[1]] = c
		}
		switch d[2] {
		case 0:
			p.Right = c
		case 1:
			p.Left = c
		}
	}
	for _, d := range descriptions { // 或者考虑节点的入度
		delete(memo, d[1])
	}
	for _, v := range memo {
		return v
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
