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

// leetcode submit region begin(Prohibit modification and deletion)
func waysToReachTarget(target int, types [][]int) int {
	const mod = 1_000_000_007
	// lc：dp
	// 爬楼梯：排列，方法数：组合
	// 时间复杂度：O(TS)，其中 T 为 target，S 为所有 count-i 之和
	dp := make([]int, target+1)
	dp[0] = 1
	for _, t := range types {
		c, m := t[0], t[1]
		for i := target; i > 0; i-- { // 滚动数组，则从 target 开始遍历
			for j, maxC := 1, min(c, i/m); j <= maxC; j++ { // 最多有 c 或 i/m 题
				dp[i] += dp[i-j*m]
			}
			dp[i] %= mod
		}
	}
	return dp[target]

	// dp
	//sort.Slice(types, func(i, j int) bool {
	//	return types[i][1] < types[j][1]
	//})
	//n := len(types)
	//dp, sum := make([][]int, target+1), make([]int, n)
	//for i := 0; i <= target; i++ {
	//	dp[i] = make([]int, n)
	//}
	//for j := 0; j < n; j++ {
	//	dp[0][j] = 1
	//	sum[j] = types[j][0] * types[j][1]
	//	if j > 0 {
	//		sum[j] += sum[j-1]
	//	}
	//}
	//fmt.Println(sum)
	//defer func() {
	//	fmt.Println(dp)
	//}()
	//for i := 1; i <= target; i++ {
	//	for j := 0; j < n; j++ {
	//		t := i - types[j][1]
	//		switch {
	//		case t < 0:
	//			break
	//		case t == 0:
	//			dp[i][j] = 1
	//		case t > 0:
	//			for k := 0; k < j; k++ {
	//				dp[i][j] += dp[t][k]
	//			}
	//			if sum[j] >= i { // TODO 超过总数
	//				dp[i][j] += dp[t][j]
	//			}
	//		}
	//	}
	//}
	//ret := 0
	//for _, v := range dp[target] {
	//	ret = (ret + v) % mod
	//}
	//return ret
}

//leetcode submit region end(Prohibit modification and deletion)
