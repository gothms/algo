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
func kthLargestPerfectSubtree(root *TreeNode, k int) int {
	// lc
	high := [11]int{}
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int { // 返回树高
		if root == nil {
			return 0 // 树高为 0
		}
		l, r := dfs(root.Left), dfs(root.Right)
		if l == -1 || l != r {
			return -1 // 非完全二叉树
		}
		l++
		high[l]++
		return l // 返回值 >=0，则是完全二叉树
	}
	dfs(root)
	i, c := len(high)-1, 0
	for ; i > 0 && c < k; i-- {
		c += high[i]
	}
	if c < k {
		return -1
	}
	return 1<<(i+1) - 1

	// 个人
	//pTree := make([]int, 0)
	//var dfs func(*TreeNode) (int, bool)
	//dfs = func(root *TreeNode) (c int, is bool) {
	//	if root == nil {
	//		return 0, true
	//	}
	//	lc, lis := dfs(root.Left)
	//	rc, ris := dfs(root.Right)
	//	c = lc + rc + 1
	//	if lis && ris && lc == rc { // 是完全二叉树
	//		pTree = append(pTree, c)
	//		is = true
	//	}
	//	return
	//}
	//dfs(root)
	//if len(pTree) < k {
	//	return -1
	//}
	//sort.Ints(pTree)
	//return pTree[len(pTree)-k]
}

//leetcode submit region end(Prohibit modification and deletion)
