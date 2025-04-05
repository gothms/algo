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
func isCompleteTree(root *TreeNode) bool {
	// lc 思路
	cnt, maxI := 0, 0
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, i int) {
		if root == nil {
			return
		}
		cnt++                   // 节点总数
		maxI = max(maxI, i)     // 最大“标号”
		dfs(root.Left, i<<1)    // 左节点的标号 = i<<1
		dfs(root.Right, i<<1+1) // 右节点的标号 = i<<1+1
	}
	dfs(root, 1)
	return cnt == maxI

	// 个人：dfs
	//mx, re := 0, false // 深度 & 是否“降”过一次深度
	//for cur := root.Left; cur != nil; cur = cur.Left {
	//	mx++
	//}
	//var dfs func(root *TreeNode, level int) bool
	//dfs = func(root *TreeNode, level int) bool {
	//	if root.Left == nil {
	//		if root.Right != nil || // 左 nil，右非 nil
	//			level != mx && (re || !re && level+1 != mx) { // 深度不同 && (降过 || 即将降深度 && 不是降一层)
	//			return false
	//		} else if level+1 == mx { // 降深度
	//			mx--
	//			re = true
	//		}
	//		return true
	//	} else if root.Right == nil { // 左非 nil，右 nil
	//		if re || level+1 != mx { // 已降过深度 || 即将降深度
	//			return false
	//		}
	//		mx-- // 降深度
	//		re = true
	//		return true
	//	}
	//	level++
	//	return dfs(root.Left, level) && dfs(root.Right, level)
	//}
	//return dfs(root, 0)
}

//leetcode submit region end(Prohibit modification and deletion)
