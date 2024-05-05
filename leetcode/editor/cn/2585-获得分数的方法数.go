//考试中有 n 种类型的题目。给你一个整数 target 和一个下标从 0 开始的二维整数数组 types ，其中 types[i] = [counti,
//marksi] 表示第 i 种类型的题目有 counti 道，每道题目对应 marksi 分。
//
// 返回你在考试中恰好得到 target 分的方法数。由于答案可能很大，结果需要对 10⁹ +7 取余。
//
// 注意，同类型题目无法区分。
//
//
// 比如说，如果有 3 道同类型题目，那么解答第 1 和第 2 道题目与解答第 1 和第 3 道题目或者第 2 和第 3 道题目是相同的。
//
//
//
//
// 示例 1：
//
// 输入：target = 6, types = [[6,1],[3,2],[2,3]]
//输出：7
//解释：要获得 6 分，你可以选择以下七种方法之一：
//- 解决 6 道第 0 种类型的题目：1 + 1 + 1 + 1 + 1 + 1 = 6
//- 解决 4 道第 0 种类型的题目和 1 道第 1 种类型的题目：1 + 1 + 1 + 1 + 2 = 6
//- 解决 2 道第 0 种类型的题目和 2 道第 1 种类型的题目：1 + 1 + 2 + 2 = 6
//- 解决 3 道第 0 种类型的题目和 1 道第 2 种类型的题目：1 + 1 + 1 + 3 = 6
//- 解决 1 道第 0 种类型的题目、1 道第 1 种类型的题目和 1 道第 2 种类型的题目：1 + 2 + 3 = 6
//- 解决 3 道第 1 种类型的题目：2 + 2 + 2 = 6
//- 解决 2 道第 2 种类型的题目：3 + 3 = 6
//
//
// 示例 2：
//
// 输入：target = 5, types = [[50,1],[50,2],[50,5]]
//输出：4
//解释：要获得 5 分，你可以选择以下四种方法之一：
//- 解决 5 道第 0 种类型的题目：1 + 1 + 1 + 1 + 1 = 5
//- 解决 3 道第 0 种类型的题目和 1 道第 1 种类型的题目：1 + 1 + 1 + 2 = 5
//- 解决 1 道第 0 种类型的题目和 2 道第 1 种类型的题目：1 + 2 + 2 = 5
//- 解决 1 道第 2 种类型的题目：5
//
//
// 示例 3：
//
// 输入：target = 18, types = [[6,1],[3,2],[2,3]]
//输出：1
//解释：只有回答所有题目才能获得 18 分。
//
//
//
//
// 提示：
//
//
// 1 <= target <= 1000
// n == types.length
// 1 <= n <= 50
// types[i].length == 2
// 1 <= counti, marksi <= 50
//
//
// Related Topics 数组 动态规划 👍 40 👎 0

package main

import (
	"fmt"
)

func main() {
	types := [][]int{{6, 1}, {2, 3}, {3, 2}}
	target := 18 // 1
	reachTarget := waysToReachTarget(target, types)
	fmt.Println(reachTarget)
}

/*
【模板】多重背包求方案数
	https://leetcode.cn/problems/number-of-ways-to-earn-points/solutions/2148313/fen-zu-bei-bao-pythonjavacgo-by-endlessc-ludl/
*/

// leetcode submit region begin(Prohibit modification and deletion)
func waysToReachTarget(target int, types [][]int) int {
	const mod = 1_000_000_007
	// dp
	// 爬楼梯：排列，方法数：组合
	// 时间复杂度：O(TS)，其中 T 为 target，S 为所有 count-i 之和
	dp := make([]int, target+1)
	dp[0] = 1
	for _, t := range types {
		for i := target; i > 0; i-- {
			//for j := 1; j <= min(t[0], min(t[0], target/t[1]); j++ { // min(t[0], target/t[1])：不能超过 target
			//	k := i - j*t[1]
			//	if k < 0 {
			//		break
			//	}
			//dp[i] += dp[k] % mod
			for j := 1; j <= min(t[0], i/t[1]); j++ { // 优化
				k := i - j*t[1]
				dp[i] += dp[k] % mod
			}
		}
	}
	return dp[target] % mod
}

//leetcode submit region end(Prohibit modification and deletion)
