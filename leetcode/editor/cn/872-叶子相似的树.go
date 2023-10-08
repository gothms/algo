//请考虑一棵二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个 叶值序列 。
//
//
//
// 举个例子，如上图所示，给定一棵叶值序列为 (6, 7, 4, 9, 8) 的树。
//
// 如果有两棵二叉树的叶值序列是相同，那么我们就认为它们是 叶相似 的。
//
// 如果给定的两个根结点分别为 root1 和 root2 的树是叶相似的，则返回 true；否则返回 false 。
//
//
//
// 示例 1：
//
//
//
//
//输入：root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,
//null,null,null,null,9,8]
//输出：true
//
//
// 示例 2：
//
//
//
//
//输入：root1 = [1,2,3], root2 = [1,3,2]
//输出：false
//
//
//
//
// 提示：
//
//
// 给定的两棵树结点数在 [1, 200] 范围内
// 给定的两棵树上的值在 [0, 200] 范围内
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 210 👎 0

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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// dfs
	temp := make([]int, 0)
	var dfs func(*TreeNode) *TreeNode
	dfs = func(r *TreeNode) *TreeNode {
		if r != nil && dfs(r.Left) == dfs(r.Right) {
			temp = append(temp, r.Val)
		}
		return r
	}
	dfs(root1) // 存储 root1 的叶值序列
	i := 0
	var check func(*TreeNode) bool
	check = func(r *TreeNode) bool { // 校验两者的叶值序列
		if r == nil {
			return true
		}
		if r.Left == nil && r.Right == nil {
			if i == len(temp) || r.Val != temp[i] { // 越界 / 不等
				return false
			}
			i++
		}
		if !check(r.Left) || !check(r.Right) {
			return false
		}
		return true
	}
	return check(root2) && i == len(temp) // i == len(temp)

	// dfs
	//vs := make([]int, 0)
	//var dfs func(*TreeNode)
	//dfs = func(r *TreeNode) {
	//	if r == nil {
	//		return
	//	}
	//	if r.Left == nil && r.Right == nil {
	//		vs = append(vs, r.Val)
	//	}
	//	dfs(r.Left)
	//	dfs(r.Right)
	//}
	//dfs(root1)
	//temp := append([]int{}, vs...)
	//vs = make([]int, 0)
	//dfs(root2)
	//return reflect.DeepEqual(temp, vs)
}

//leetcode submit region end(Prohibit modification and deletion)
