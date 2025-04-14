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
func findFrequentTreeSum(root *TreeNode) []int {
	mx := 0
	memo := make(map[int]int)
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		v := root.Val + dfs(root.Left) + dfs(root.Right)
		if memo[v] == mx {
			mx++
		}
		memo[v]++
		return v
	}
	dfs(root)
	ans := make([]int, 0)
	for k, v := range memo {
		if v == mx {
			ans = append(ans, k)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
