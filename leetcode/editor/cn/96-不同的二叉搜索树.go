//给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：5
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 19
//
//
// Related Topics 树 二叉搜索树 数学 动态规划 二叉树 👍 2340 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	// 数学：卡塔兰数
	// https://leetcode.cn/problems/unique-binary-search-trees/solutions/329807/bu-tong-de-er-cha-sou-suo-shu-by-leetcode-solution/
	C := 1
	for i := 0; i < n; i++ {
		C = (i<<1 + 1) * C << 1 / (i + 2)
	}
	return C

	// dp
	//dp := make([]int, n+1)
	//dp[0], dp[1] = 1, 1
	//for i := 2; i <= n; i++ {
	//	for j := 0; j < i; j++ {
	//		dp[i] += dp[j] * dp[i-1-j]
	//	}
	//}
	//return dp[n]

	// 记忆化搜索
	//memo := make([]int, n+1)
	//var dfs func(int) int
	//dfs = func(i int) int {
	//	if i <= 1 {
	//		return 1
	//	}
	//	var v *int
	//	if v = &memo[i]; *v > 0 {
	//		return *v
	//	}
	//	cnt := 0
	//	for j := 0; j < i; j++ {
	//		cnt += dfs(j) * dfs(i-1-j) // 左 * 右
	//	}
	//	*v = cnt
	//	return cnt
	//}
	//return dfs(n)

	// 旧写法
	//memo := make([]int, n+1)
	//memo[0], memo[1] = 1, 1
	//return numTreesRecursion(n, memo)
}

func numTreesRecursion(n int, memo []int) int {
	if memo[n] == 0 {
		for i := 0; i < n; i++ {
			memo[n] += numTreesRecursion(n-1-i, memo) * numTreesRecursion(i, memo)
		}
	}
	return memo[n]
}

//leetcode submit region end(Prohibit modification and deletion)
