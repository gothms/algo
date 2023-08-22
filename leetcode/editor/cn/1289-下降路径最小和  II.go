//给你一个 n x n 整数矩阵 grid ，请你返回 非零偏移下降路径 数字和的最小值。
//
// 非零偏移下降路径 定义为：从 grid 数组中的每一行选择一个数字，且按顺序选出来的数字中，相邻数字不在原数组的同一列。
//
//
//
// 示例 1：
//
//
//
//
//输入：grid = [[1,2,3],[4,5,6],[7,8,9]]
//输出：13
//解释：
//所有非零偏移下降路径包括：
//[1,5,9], [1,5,7], [1,6,7], [1,6,8],
//[2,4,8], [2,4,9], [2,6,7], [2,6,8],
//[3,4,8], [3,4,9], [3,5,7], [3,5,9]
//下降路径中数字和最小的是 [1,5,7] ，所以答案是 13 。
//
//
// 示例 2：
//
//
//输入：grid = [[7]]
//输出：7
//
//
//
//
// 提示：
//
//
// n == grid.length == grid[i].length
// 1 <= n <= 200
// -99 <= grid[i][j] <= 99
//
//
// Related Topics 数组 动态规划 矩阵 👍 95 👎 0

package main

import (
	"fmt"
)

func main() {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// -192
	grid = [][]int{
		{-73, 61, 43, -48, -36},
		{3, 30, 27, 57, 10},
		{96, -76, 84, 59, -15},
		{5, -49, 76, 31, -7},
		{97, 91, 61, -46, 67}}
	//grid = [][]int{{9}}
	sum := minFallingPathSum(grid)
	fmt.Println(sum)
}

/*
思路：dp
	1.如何选择 非零偏移下降路径 的数字？
		我们记录每一行的累加最小值 & 累加次最小值为 m0、m1
		对于当前行 grid[i]，grid[i][j] 与上一行的元素 grid[i-1][j] 是同一列时
		那么，累加和 = m1 + grid[i][j]
		否则，累加和 = m0 + grid[i][j]
	2.dp
		我们用一个数组记录每一行的累加最小值
		从这些最小值中挑选出 m0、m1，用于下一行的累加最小值计算
		若与上一行元素同行：dp[j] = m1 + grid[i][j]
		若与上一行元素不同行：dp[j] = m0 + grid[i][j]
		详见：dp 非最优
	3.优化
		dp 数组记录的元素中，只有 最小值 和 次最小值 是有效值，其余的元素都是可以淘汰的
		那么我们只需要记录 当前行的累加最小值、次累加最小值、累加最小值的索引
		详见：dp 终版
*/
// leetcode submit region begin(Prohibit modification and deletion)
func minFallingPathSum(grid [][]int) int {
	// dp：终版
	min0, min1, minIdx, n := 0, 0, -1, len(grid)
	for i := 0; i < n; i++ {
		// 当前行的累加最小值、次累加最小值、累加最小值的索引
		nMin0, nMin1, nMinIdx := 20000, 20000, -1
		for j := 0; j < n; j++ {
			sum := grid[i][j]
			if j == minIdx { // 相邻数字是原数组的同一列
				sum += min1
			} else {
				sum += min0
			}
			if sum < nMin0 { // 筛选累加最小值
				nMin0, nMin1, nMinIdx = sum, nMin0, j
			} else if sum < nMin1 { // 筛选次累加最小值
				nMin1 = sum
			}
		}
		// 保存当前行数据
		min0, min1, minIdx = nMin0, nMin1, nMinIdx
	}
	return min0

	// dp：非最优
	//minV, n := math.MaxInt32, len(grid)
	//dp := make([]int, n)
	//copy(dp, grid[0]) // 第一行
	//for i := 1; i < n; i++ {
	//	m0, m1 := math.MaxInt32, math.MaxInt32
	//	for j := 0; j < n; j++ { // 查询上一行累加最小值 & 累加次最小值
	//		if dp[j] < m1 {
	//			m1 = dp[j]
	//			if m0 > m1 {
	//				m0, m1 = m1, m0
	//			}
	//		}
	//	}
	//	for j := 0; j < n; j++ { // 计算当前行累加最小值
	//		if dp[j] == m0 {
	//			dp[j] = m1 + grid[i][j]
	//		} else {
	//			dp[j] = m0 + grid[i][j]
	//		}
	//	}
	//}
	//for _, v := range dp { // 查询最小值
	//	if v < minV {
	//		minV = v
	//	}
	//}
	//return minV

	// 理解错题意
	//minV, n := math.MaxInt32, len(grid)
	//if n == 1 {
	//	return grid[0][0]
	//}
	//minVal := func(a, b int) int {
	//	if a < b {
	//		return a
	//	}
	//	return b
	//}
	//dp := make([]int, n+2)
	//dp[0], dp[n+1] = math.MaxInt32, math.MaxInt32
	//for i := 0; i < n; i++ {
	//	dpI := make([]int, n+2)
	//	dpI[0], dpI[n+1] = math.MaxInt32, math.MaxInt32
	//	for j := 0; j < n; j++ {
	//		dpI[j+1] = minVal(dp[j], dp[j+2]) + grid[i][j]
	//	}
	//	dp = dpI
	//}
	//for _, v := range dp {
	//	minV = minVal(minV, v)
	//}
	//return minV
}

//leetcode submit region end(Prohibit modification and deletion)
