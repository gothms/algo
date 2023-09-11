//检查子树。你有两棵非常大的二叉树：T1，有几万个节点；T2，有几万个节点。设计一个算法，判断 T2 是否为 T1 的子树。
//
// 如果 T1 有这么一个节点 n，其子树与 T2 一模一样，则 T2 为 T1 的子树，也就是说，从节点 n 处把树砍断，得到的树与 T2 完全相同。
//
// 注意：此题相对书上原题略有改动。
//
// 示例1:
//
//
// 输入：t1 = [1, 2, 3], t2 = [2]
// 输出：true
//
//
// 示例2:
//
//
// 输入：t1 = [1, null, 2, 4], t2 = [3, 2]
// 输出：false
//
//
// 提示：
//
//
// 树的节点数目范围为[0, 20000]。
//
//
// Related Topics 树 深度优先搜索 二叉树 字符串匹配 哈希函数 👍 80 👎 0

package main

func main() {

}

/*
思路：dfs
	1.匹配两个二叉树的子树，只需要在 dfs 时比较对应节点是否都为 nil 或 Val 是否相等
		关键在于匹配的起始、结束、细节的考虑
	2.两个dfs
		第一个 dfs：匹配 t1 的某个节点是否匹配 t2 的根
			如果 val 相等，则开始尝试匹配两棵完整的二叉树
		第二个 dfs：匹配两棵二叉树是否”相等“，关键在于细节的判断
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
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	// dfs
	if t2 == nil {
		return true
	}
	var ok func(*TreeNode, *TreeNode) bool
	ok = func(r *TreeNode, sub *TreeNode) bool { // 匹配两棵树是否相等
		if r == nil && sub == nil { // 同为nil，匹配
			return true
		}
		if r == nil || sub == nil || r.Val != sub.Val { // 不匹配
			return false
		}
		return ok(r.Left, sub.Left) && ok(r.Right, sub.Right) // 匹配子节点
	}
	var dfs func(*TreeNode) bool
	dfs = func(r *TreeNode) bool { // 匹配t1的节点与t2的根节点
		if r == nil {
			return false
		}
		// 尝试匹配子树
		if r.Val == t2.Val && ok(r.Left, t2.Left) && ok(r.Right, t2.Right) {
			return true
		}
		return dfs(r.Left) || dfs(r.Right)
	}
	return dfs(t1)
}

//leetcode submit region end(Prohibit modification and deletion)

// 子树得完全匹配
//var dfs func(*TreeNode, *strings.Builder)
//dfs = func(r *TreeNode, sb *strings.Builder) {
//	if r == nil {
//		return
//	}
//	sb.WriteRune(rune(r.Val))
//	dfs(r.Left, sb)
//	dfs(r.Right, sb)
//}
//var s, sub strings.Builder
//dfs(t1, &s)
//dfs(t2, &sub)
