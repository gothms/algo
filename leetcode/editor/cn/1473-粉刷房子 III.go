//在一个小城市里，有 m 个房子排成一排，你需要给每个房子涂上 n 种颜色之一（颜色编号为 1 到 n ）。有的房子去年夏天已经涂过颜色了，所以这些房子不可以
//被重新涂色。
//
// 我们将连续相同颜色尽可能多的房子称为一个街区。（比方说 houses = [1,2,2,3,3,2,1,1] ，它包含 5 个街区 [{1}, {2,2}
//, {3,3}, {2}, {1,1}] 。）
//
// 给你一个数组 houses ，一个 m * n 的矩阵 cost 和一个整数 target ，其中：
//
//
// houses[i]：是第 i 个房子的颜色，0 表示这个房子还没有被涂色。
// cost[i][j]：是将第 i 个房子涂成颜色 j+1 的花费。
//
//
// 请你返回房子涂色方案的最小总花费，使得每个房子都被涂色后，恰好组成 target 个街区。如果没有可用的涂色方案，请返回 -1 。
//
//
//
// 示例 1：
//
//
//输入：houses = [0,0,0,0,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n
//= 2, target = 3
//输出：9
//解释：房子涂色方案为 [1,2,2,1,1]
//此方案包含 target = 3 个街区，分别是 [{1}, {2,2}, {1,1}]。
//涂色的总花费为 (1 + 1 + 1 + 1 + 5) = 9。
//
//
// 示例 2：
//
//
//输入：houses = [0,2,1,2,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n
//= 2, target = 3
//输出：11
//解释：有的房子已经被涂色了，在此基础上涂色方案为 [2,2,1,2,2]
//此方案包含 target = 3 个街区，分别是 [{2,2}, {1}, {2,2}]。
//给第一个和最后一个房子涂色的花费为 (10 + 1) = 11。
//
//
// 示例 3：
//
//
//输入：houses = [0,0,0,0,0], cost = [[1,10],[10,1],[1,10],[10,1],[1,10]], m = 5,
//n = 2, target = 5
//输出：5
//
//
// 示例 4：
//
//
//输入：houses = [3,1,2,3], cost = [[1,1,1],[1,1,1],[1,1,1],[1,1,1]], m = 4, n = 3,
// target = 3
//输出：-1
//解释：房子已经被涂色并组成了 4 个街区，分别是 [{3},{1},{2},{3}] ，无法形成 target = 3 个街区。
//
//
//
//
// 提示：
//
//
// m == houses.length == cost.length
// n == cost[i].length
// 1 <= m <= 100
// 1 <= n <= 20
// 1 <= target <= m
// 0 <= houses[i] <= n
// 1 <= cost[i][j] <= 10^4
//
//
// Related Topics 数组 动态规划 👍 198 👎 0

package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	// dp：三维 m、n、target
	// best：二维 m、target，数据结构 firstMax、firstMaxIdx、secondMax
	// 状态转移方程
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	const inf = math.MaxInt32 >> 1
	type fsMS struct {
		f, s int // 第一小，第二小
		idx  int // 第一小涂的颜色
	}
	dp, memo := make([][][]int, m), make([][]fsMS, m)
	for i := 0; i < m; i++ {
		houses[i]-- // house 从 0 开始，方便计算
		dp[i] = make([][]int, n)
		memo[i] = make([]fsMS, target)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, target)
			for k := 0; k < target; k++ {
				dp[i][j][k] = inf // 初始化最小 cost
				memo[i][k] = fsMS{inf, inf, -1}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if houses[i] != -1 && houses[i] != j { // 已经涂色的房子，不能涂其他色
				continue
			}
			for k := 0; k < target; k++ {
				v := &dp[i][j][k]
				if i == 0 { // 状态转移方程：dp[0][j][0] = 0
					if k == 0 {
						*v = 0
					}
				} else {
					*v = dp[i-1][j][k] // 和上一个房子涂一样的色
					if k > 0 {         // 涂不一样的色
						if fs := &memo[i-1][k-1]; fs.idx == j {
							*v = minVal(*v, fs.s) // 选第二小
						} else {
							*v = minVal(*v, fs.f) // 选第一小
						}
					}
				}
				if *v != inf && houses[i] == -1 { // 未涂色
					*v += cost[i][j]
				}
				if fs := &memo[i][k]; *v < fs.f { // 修正第一第二小
					fs.f, fs.s, fs.idx = *v, fs.f, j
				} else if *v < fs.s {
					fs.s = *v
				}
			}
		}
	}
	mc, last := inf, target-1
	for _, d := range dp[m-1] { // 遍历 [0,n)
		mc = minVal(mc, d[last])
	}
	if mc == inf {
		return -1
	}
	return mc
}

// leetcode submit region end(Prohibit modification and deletion)
//type entry struct {
//	first, firstIdx, second int
//}
//
//func minCost(houses []int, cost [][]int, m, n, target int) int {
//	const inf = math.MaxInt64 / 2 // 防止整数相加溢出
//
//	// 将颜色调整为从 0 开始编号，没有被涂色标记为 -1
//	for i := range houses {
//		houses[i]--
//	}
//
//	// dp 所有元素初始化为极大值
//	dp := make([][][]int, m)
//	for i := range dp {
//		dp[i] = make([][]int, n)
//		for j := range dp[i] {
//			dp[i][j] = make([]int, target)
//			for k := range dp[i][j] {
//				dp[i][j][k] = inf
//			}
//		}
//	}
//	best := make([][]entry, m)
//	for i := range best {
//		best[i] = make([]entry, target)
//		for j := range best[i] {
//			best[i][j] = entry{inf, -1, inf}
//		}
//	}
//
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if houses[i] != -1 && houses[i] != j {
//				continue
//			}
//
//			for k := 0; k < target; k++ {
//				if i == 0 {
//					if k == 0 {
//						dp[i][j][k] = 0
//					}
//				} else {
//					dp[i][j][k] = dp[i-1][j][k]
//					if k > 0 {
//						// 使用 best(i-1,k-1) 直接得到 dp(i,j,k) 的值
//						if b := best[i-1][k-1]; j == b.firstIdx {
//							dp[i][j][k] = min(dp[i][j][k], b.second)
//						} else {
//							dp[i][j][k] = min(dp[i][j][k], b.first)
//						}
//					}
//				}
//
//				if dp[i][j][k] != inf && houses[i] == -1 {
//					dp[i][j][k] += cost[i][j]
//				}
//
//				// 用 dp(i,j,k) 更新 best(i,k)
//				if b := &best[i][k]; dp[i][j][k] < b.first {
//					b.second = b.first
//					b.first = dp[i][j][k]
//					b.firstIdx = j
//				} else if dp[i][j][k] < b.second {
//					b.second = dp[i][j][k]
//				}
//			}
//		}
//	}
//
//	ans := inf
//	for _, res := range dp[m-1] {
//		ans = min(ans, res[target-1])
//	}
//	if ans == inf {
//		return -1
//	}
//	return ans
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
