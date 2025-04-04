package main

func main() {

}

/*
思路：dfs
	1.递
		递归参数为：当前节点 r *TreeNode、节点的深度 level int
	2.归
		左子树的最深子节点深度为 ll，右子树的最深子节点深度为 lr，递归返回值分三种情况：
		ll > rl：返回 ll，及那个左子树中唯一的最深子节点
		ll < rl：返回 lr，及那个右子树中唯一的最深子节点
		ll == rl：说明左右子树的最深子节点在同一深度，则返回 ll，及左右子树的最近公共祖先
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var (
		ans *TreeNode
		d   int
	)
	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, level int) int {
		if root == nil {
			return level
		}
		level++
		l, r := dfs(root.Left, level), dfs(root.Right, level)
		if l == r && l >= d {
			ans, d = root, l
		}
		return max(l, r)
	}
	dfs(root, 0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func lcaDeepestLeaves_(root *TreeNode) *TreeNode {
	// dfs
	var dfs func(*TreeNode, int) (*TreeNode, int)
	dfs = func(r *TreeNode, level int) (*TreeNode, int) {
		if r == nil {
			return nil, level
		}
		lt, ll := dfs(r.Left, level+1) // 继续递
		rt, rl := dfs(r.Right, level+1)
		switch { // 归的三种情况
		case ll > rl:
			return lt, ll
		case ll < rl:
			return rt, rl
		default:
			return r, ll
		}
	}
	p, _ := dfs(root, 0)
	return p
}
