//给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 k 。
//
// 返回到目标结点 target 距离为 k 的所有结点的值的列表。 答案可以以 任何顺序 返回。
//
//
//
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
//输出：[7,4,1]
//解释：所求结点为与目标结点（值为 5）距离为 2 的结点，值分别为 7，4，以及 1
//
//
// 示例 2:
//
//
//输入: root = [1], target = 1, k = 3
//输出: []
//
//
//
//
// 提示:
//
//
// 节点数在 [1, 500] 范围内
// 0 <= Node.val <= 500
// Node.val 中所有值 不同
// 目标结点 target 是树上的结点。
// 0 <= k <= 1000
//
//
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 656 👎 0

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
func distanceK(root *TreeNode, target *TreeNode, k int) []int {

}

//leetcode submit region end(Prohibit modification and deletion)

func distanceK_(root *TreeNode, target *TreeNode, k int) []int {
	// 比下面写法更好
	ans := make([]int, 0)
	var down func(*TreeNode, int)
	down = func(r *TreeNode, t int) {
		if r == nil {
			return
		}
		if t == 0 {
			ans = append(ans, r.Val)
			return
		}
		down(r.Left, t-1)
		down(r.Right, t-1)
	}
	var dfs func(*TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return -1
		}
		if r == target {
			down(r, k)
			return k - 1
		}
		if left := dfs(r.Left); left >= 0 {
			if left == 0 {
				ans = append(ans, r.Val)
				return -1
			}
			down(r.Right, left-1)
			return left - 1
		} else if right := dfs(r.Right); right >= 0 {
			if right == 0 {
				ans = append(ans, r.Val)
				return -1
			}
			down(r.Left, right-1)
			return right - 1
		}
		return -1
	}
	dfs(root)
	return ans

	// dfs
	//if k == 0 { // fast path
	//	return []int{target.Val}
	//}
	//dk := make([]int, 0)
	//var down func(*TreeNode, int, int)
	//down = func(r *TreeNode, i, t int) { // 搜索 target 的子节点
	//	if r == nil {
	//		return
	//	}
	//	if i == t {
	//		dk = append(dk, r.Val) // 终止 dfs，更深的节点，距离已超过 k
	//		return
	//	}
	//	i++
	//	down(r.Left, i, t)
	//	down(r.Right, i, t)
	//}
	//var dfs func(*TreeNode) int
	//dfs = func(r *TreeNode) int { // 递：查找 target，归：返回与 target 的距离（-1 表示不合法/超出 k）
	//	if r == nil {
	//		return -1 // 路径中没找到 target
	//	}
	//	if r == target { // 找到 target
	//		down(r.Left, 1, k)
	//		down(r.Right, 1, k)
	//		return 0 // 返回距离为 0
	//	}
	//	// 写法一
	//	//dist := dfs(r.Left) + 1
	//	//if dist > 0 { // target 在左子树
	//	//	if dist == k {
	//	//		dk = append(dk, r.Val)
	//	//		return -1 // 已超出 k
	//	//	} else if dist < k {
	//	//		down(r.Right, 1, k-dist) // 找右子树
	//	//	}
	//	//	return dist
	//	//}
	//	//dist = dfs(r.Right) + 1
	//	//if dist > 0 { // target 在右子树
	//	//	if dist == k {
	//	//		dk = append(dk, r.Val)
	//	//		return -1
	//	//	} else if dist < k {
	//	//		down(r.Left, 1, k-dist) // 找左子树
	//	//	}
	//	//	return dist
	//	//}
	//	//return -1
	//
	//	// 写法二
	//	dist, left := 0, false
	//	if dist = dfs(r.Left); dist == -1 {
	//		left = true
	//		if dist = dfs(r.Right); dist == -1 {
	//			return -1
	//		}
	//	}
	//	if dist++; dist == k { // dist + 1
	//		dk = append(dk, r.Val)
	//		return -1
	//	}
	//	if left {
	//		down(r.Left, 1, k-dist) // 找左子树
	//	} else {
	//		down(r.Right, 1, k-dist) // 找右子树
	//	}
	//	return dist
	//}
	//dfs(root)
	//return dk
}
