//给定一个有 N 个结点的二叉树的根结点 root，树中的每个结点上都对应有 node.val 枚硬币，并且总共有 N 枚硬币。
//
// 在一次移动中，我们可以选择两个相邻的结点，然后将一枚硬币从其中一个结点移动到另一个结点。(移动可以是从父结点到子结点，或者从子结点移动到父结点。)。
//
// 返回使每个结点上只有一枚硬币所需的移动次数。
//
//
//
// 示例 1：
//
//
//
// 输入：[3,0,0]
//输出：2
//解释：从树的根结点开始，我们将一枚硬币移到它的左子结点上，一枚硬币移到它的右子结点上。
//
//
// 示例 2：
//
//
//
// 输入：[0,3,0]
//输出：3
//解释：从根结点的左子结点开始，我们将两枚硬币移到根结点上 [移动两次]。然后，我们把一枚硬币从根结点移到右子结点上。
//
//
// 示例 3：
//
//
//
// 输入：[1,0,2]
//输出：2
//
//
// 示例 4：
//
//
//
// 输入：[1,0,0,null,3]
//输出：4
//
//
//
//
// 提示：
//
//
// 1<= N <= 100
// 0 <= node.val <= N
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 364 👎 0

package main

func main() {

}

/*
思路：dfs
	1.关键点
		总结为一句话就是，对于任意子树curr，它多余/缺少硬币，都应该先“内部解决”，再寻求“外部帮助”
	2.举个例子
		假设对于任意节点curr，如果它和它的左右子树的总硬币有多余，但是curr和左子树的总硬币却不够
		那么最优的硬币移动方案是，优先从它的右子树移动到curr和左子树，而不是优先向curr的父节点索要
	3.dfs分析
		每个子树curr，都应该有两个值需要处理（两者的算法不一样）
		1)统计curr子树硬币的移动次数
		2)左和右子树分别需要往父节点移出多少个硬币，或者向父节点索要多少个硬币
			移出/索要硬币，具有传递性，也就是移动的次数会“累加”
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
func distributeCoins(root *TreeNode) int {
	// dfs
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}
	dist := 0
	var dfs func(*TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		// lv/rv为正：往父节点分别移出多少个硬币，为负：向父节点分别索要多少个硬币
		lv, rv := dfs(r.Left), dfs(r.Right)
		dist += abs(lv) + abs(rv)  // 统计移动硬币的次数
		return lv + rv + r.Val - 1 // 当前节点向父节点移出/索要硬币
	}
	dfs(root)
	return dist
}

//leetcode submit region end(Prohibit modification and deletion)
