package main

import "math"

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
func findMode(root *TreeNode) []int {
	// lc：O(1) 空间
	pre, cur, mx := math.MinInt64, 0, 0
	ans := make([]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		v := root.Val
		if v == pre {
			cur++
		} else {
			pre, cur = v, 1
		}
		if cur == mx {
			ans = append(ans, v)
		} else if cur > mx {
			mx = cur
			ans = []int{v}
		}
		dfs(root.Right)
	}
	dfs(root)
	return ans

	//memo := make(map[int]int)
	//mx := 0
	//var dfs func(*TreeNode)
	//dfs = func(root *TreeNode) {
	//	if root == nil {
	//		return
	//	}
	//	memo[root.Val]++
	//	mx = max(mx, memo[root.Val])
	//	dfs(root.Left)
	//	dfs(root.Right)
	//}
	//dfs(root)
	//ans := make([]int, 0, 1)
	//for k, v := range memo {
	//	if v == mx {
	//		ans = append(ans, k)
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
