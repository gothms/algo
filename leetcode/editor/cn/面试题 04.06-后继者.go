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
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	// 二叉搜索树性质
	var ans *TreeNode
	if p.Right != nil {
		ans = p.Right
		for ans.Left != nil {
			ans = ans.Left
		}
		return ans
	}
	for cur := root; cur != nil; {
		if cur.Val > p.Val {
			ans = cur
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return ans

	//var last *TreeNode
	//var dfs func(root *TreeNode) *TreeNode
	//dfs = func(root *TreeNode) *TreeNode {
	//	if root == nil {
	//		return nil
	//	}
	//	if l := dfs(root.Left); l != nil {
	//		return l
	//	}
	//	if last == p {
	//		return root
	//	}
	//	last = root
	//	return dfs(root.Right)
	//}
	//return dfs(root)
}

//leetcode submit region end(Prohibit modification and deletion)
