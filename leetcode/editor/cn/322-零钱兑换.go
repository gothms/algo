//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
// 你可以认为每种硬币的数量是无限的。
//
//
//
// 示例 1：
//
//
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//
// 示例 2：
//
//
//输入：coins = [2], amount = 3
//输出：-1
//
// 示例 3：
//
//
//输入：coins = [1], amount = 0
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 2³¹ - 1
// 0 <= amount <= 10⁴
//
//
// Related Topics 广度优先搜索 数组 动态规划 👍 2749 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	c := []int{1, 2, 5}
	amount := 11
	c = []int{2}
	amount = 3
	change := coinChange(c, amount)
	fmt.Println(change)
}

// leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, v := range coins {
			if i >= v {
				dp[i] = min(dp[i], dp[i-v]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]

	// 参考 2585，还未AC
	//n := amount + 1
	//dp := make([]int, n)
	//for i := 1; i < n; i++ {
	//	dp[i] = n
	//}
	//for _, c := range coins {
	//	for i := amount; i > 0; i-- {
	//		v := n
	//		for j, m := 1, i/c; j <= m; j++ {
	//			v = min(v, dp[i-j*c]+j)
	//		}
	//		if v < n {
	//			dp[i] = v
	//		}
	//	}
	//}
	//if dp[amount] < n {
	//	return dp[amount]
	//}
	//return -1

	// dp
	//n := amount + 1
	//dp := make([]int, n)
	//for i := 1; i < n; i++ {
	//	dp[i] = n
	//	for _, c := range coins {
	//		if i >= c {
	//			dp[i] = min(dp[i], dp[i-c]+1)
	//		}
	//	}
	//}
	//if dp[amount] == n {
	//	return -1
	//}
	//return dp[amount]
}

//leetcode submit region end(Prohibit modification and deletion)
