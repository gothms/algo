//这里有 n 个一样的骰子，每个骰子上都有 k 个面，分别标号为 1 到 k 。
//
// 给定三个整数 n , k 和 target ，返回可能的方式(从总共 kⁿ 种方式中)滚动骰子的数量，使正面朝上的数字之和等于 target 。
//
// 答案可能很大，你需要对 10⁹ + 7 取模 。
//
//
//
// 示例 1：
//
//
//输入：n = 1, k = 6, target = 3
//输出：1
//解释：你扔一个有 6 个面的骰子。
//得到 3 的和只有一种方法。
//
//
// 示例 2：
//
//
//输入：n = 2, k = 6, target = 7
//输出：6
//解释：你扔两个骰子，每个骰子有 6 个面。
//得到 7 的和有 6 种方法：1+6 2+5 3+4 4+3 5+2 6+1。
//
//
// 示例 3：
//
//
//输入：n = 30, k = 30, target = 500
//输出：222616187
//解释：返回的结果必须是对 10⁹ + 7 取模。
//
//
//
// 提示：
//
//
// 1 <= n, k <= 30
// 1 <= target <= 1000
//
//
// Related Topics 动态规划 👍 178 👎 0

package main

import "fmt"

func main() {
	n, k, target := 1, 6, 3
	n, k, target = 2, 6, 7
	n, k, target = 30, 30, 500 // 222616187
	toTarget := numRollsToTarget(n, k, target)
	fmt.Println(toTarget)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numRollsToTarget(n int, k int, target int) int {
	// 优化
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	const mod = 1_000_000_007
	dp := make([]int, target+1)
	dp[0] = 1
	var max int // 当前最大累加值
	for i := 0; i < n; i++ {
		max = minVal(max+k, target)
		for j := 1; j <= max; j++ { // 前缀和
			dp[j] += dp[j-1]
		}
		for j := max; j >= k; j-- {
			dp[j] = (dp[j] - dp[j-k]) % mod
		}
	}
	return dp[target-n]

	// dp
	//minVal := func(a, b int) int {
	//	if b < a {
	//		return b
	//	}
	//	return a
	//}
	//const mod = 1_000_000_007
	//dp := make([]int, target+1)
	//max := minVal(k, target) // 当前最大累加值
	//for i := 1; i <= max; i++ {
	//	dp[i] = 1 // 初始化
	//}
	//for i := 1; i < n; i++ {
	//	temp := make([]int, target+1)
	//	for j := 1; j <= max; j++ {
	//		for v := 1; v <= k; v++ {
	//			if j+v > target {
	//				break
	//			}
	//			temp[j+v] = (temp[j+v] + dp[j]) % mod
	//		}
	//	}
	//	max = minVal(max+k, target)
	//	dp = temp
	//}
	//return dp[target]
}

//leetcode submit region end(Prohibit modification and deletion)
