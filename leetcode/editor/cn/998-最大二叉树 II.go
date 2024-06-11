package main

func main() {

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
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {

}

//leetcode submit region end(Prohibit modification and deletion)

func insertIntoMaxTree_(root *TreeNode, val int) *TreeNode {
	if val > root.Val { // fast path
		return &TreeNode{val, root, nil}
	}
	for p, cur := root, root.Right; ; p, cur = cur, cur.Right { // 只考虑右子树
		if cur == nil || val > cur.Val { // 添加 / 插入
			p.Right = &TreeNode{val, cur, nil}
			return root
		}
	}
}
