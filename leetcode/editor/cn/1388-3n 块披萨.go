//给你一个披萨，它由 3n 块不同大小的部分组成，现在你和你的朋友们需要按照如下规则来分披萨：
//
//
// 你挑选 任意 一块披萨。
// Alice 将会挑选你所选择的披萨逆时针方向的下一块披萨。
// Bob 将会挑选你所选择的披萨顺时针方向的下一块披萨。
// 重复上述过程直到没有披萨剩下。
//
//
// 每一块披萨的大小按顺时针方向由循环数组 slices 表示。
//
// 请你返回你可以获得的披萨大小总和的最大值。
//
//
//
// 示例 1：
//
//
//
//
//输入：slices = [1,2,3,4,5,6]
//输出：10
//解释：选择大小为 4 的披萨，Alice 和 Bob 分别挑选大小为 3 和 5 的披萨。然后你选择大小为 6 的披萨，Alice 和 Bob 分别挑选大小
//为 2 和 1 的披萨。你获得的披萨总大小为 4 + 6 = 10 。
//
//
// 示例 2：
//
//
//
//
//输入：slices = [8,9,8,6,1,1]
//输出：16
//解释：两轮都选大小为 8 的披萨。如果你选择大小为 9 的披萨，你的朋友们就会选择大小为 8 的披萨，这种情况下你的总和不是最大的。
//
//
//
//
// 提示：
//
//
// 1 <= slices.length <= 500
// slices.length % 3 == 0
// 1 <= slices[i] <= 1000
//
//
// Related Topics 贪心 数组 动态规划 堆（优先队列） 👍 124 👎 0

package main

import (
	"fmt"
)

func main() {
	slices := []int{1, 2, 3, 4, 5, 6}                  // 10
	slices = []int{8, 9, 8, 6, 1, 1}                   // 16
	slices = []int{9, 5, 1, 7, 8, 4, 4, 5, 5, 8, 7, 7} // 30
	//slices = []int{3, 1, 2}                            // 30
	sizeSlices := maxSizeSlices(slices)
	fmt.Println(sizeSlices)
}

/*
思路：记忆化搜索
	1.不可以挑选的披萨？假如 slices = [1,2,3,4,5,6]，如果挑选了 1，那么下一块不可挑选的是 2、6
		即当我们挑选当前最小的索引 0 时，只要不是 索引1 或 最后一个索引，其他的都可以挑选
		比如我们可以挑选出 1,3,5，对应的索引为 0,2,4
		结论：只要没有相邻的元素的组合，都可以成为最终挑选的结果
	2.模型定义
		我们用 i 表示当前考虑要不要挑选的索引，j 表示当前已经挑选了几个
		memo[i][j] 表示 i j 条件下，最大的披萨和
	3.记忆化搜索
		由 1. 分析我们可知，挑选了 索引0，就不能挑选最后一个索引
		所以我们要对 slices 进行两次搜索，分别是 slices[0:-1] 和 slices[1:]
	4.核心逻辑为：为了方便判断，我们倒着搜索 slices
		不选当前数 s[i]，则下次搜索为 dfs(i-1, j)
			即什么都不做，memo[i][j] 仍然为上一次选择 j 个披萨的最大值
		选当前数 s[i]，则下次搜索为 dfs(i-2, j-1)，但要加上 s[i]，i-2 < 0 表示没有可搜索的了
			即选择这块披萨，memo[i][j] 为上次选择 j-1 个披萨的最大值，加上 s[i]
思路：dp
	思路和记忆化搜索相同，状态转移方程为：
	dp[i][j] = maxVal(dp[i-1][j], dp[i-2][j-1]+s[i])
	dp[i-1][j]：不选择当前数 s[i]
	dp[i-2][j-1]+s[i]：选择当前数 s[i]，所以上一次搜索是 i-2，且个数是 j-1
*/
// leetcode submit region begin(Prohibit modification and deletion)
func maxSizeSlices(slices []int) int {
	// dp
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n, k := len(slices)-1, len(slices)/3
	dpMSS := func(s []int) int {
		dp := make([][]int, n)
		for i := 0; i < n; i++ {
			dp[i] = make([]int, k+1)
		}
		// 初始化前两个元素的选择
		dp[0][1], dp[1][1] = s[0], maxVal(s[0], s[1])
		for i := 2; i < n; i++ {
			for j := 1; j <= k; j++ {
				dp[i][j] = maxVal(dp[i-1][j], dp[i-2][j-1]+s[i])
			} // 不选 s[i] 或选 s[i]
		}
		return dp[n-1][k]
	}
	return maxVal(dpMSS(slices[:n]), dpMSS(slices[1:])) // 两次搜索

	// 记忆化搜索
	//maxVal := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//n, k := len(slices)-1, len(slices)/3
	//m0, m1 := make([][]int, n), make([][]int, n)
	//for i := 0; i < n; i++ {
	//	m0[i], m1[i] = make([]int, k+1), make([]int, k+1)
	//}
	//var dfs func(int, int, [][]int, []int) int
	//dfs = func(i, j int, m [][]int, s []int) int {
	//	if i < 0 || j == 0 {
	//		return 0
	//	}
	//	var v *int
	//	if v = &m[i][j]; *v > 0 { // 已搜索
	//		return *v
	//	}
	//	// 不选 s[i] 或先把 s[i] 选了（防止 i-3<0），而且先把 j=1 处理了
	//	cnt := maxVal(dfs(i-1, j, m, s), dfs(i-2, j-1, m, s)+s[i])
	//	for ni := i - 3; ni >= 0; ni-- {
	//		cnt = maxVal(cnt, dfs(ni, j-1, m, s)+s[i])
	//	}
	//	*v = cnt // 记忆化
	//	return cnt
	//}
	//return maxVal(dfs(n-1, k, m0, slices), dfs(n-1, k, m1, slices[1:]))
}

//leetcode submit region end(Prohibit modification and deletion)
