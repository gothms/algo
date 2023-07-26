//给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
//
// 下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第
//一个元素）。具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1
//, col + 1) 。
//
//
//
// 示例 1：
//
//
//
//
//输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
//输出：13
//解释：如图所示，为和最小的两条下降路径
//
//
// 示例 2：
//
//
//
//
//输入：matrix = [[-19,57],[-40,-5]]
//输出：-59
//解释：如图所示，为和最小的下降路径
//
//
//
//
// 提示：
//
//
// n == matrix.length == matrix[i].length
// 1 <= n <= 100
// -100 <= matrix[i][j] <= 100
//
//
// Related Topics 数组 动态规划 矩阵 👍 240 👎 0

package main

func main() {

}

/*
思路：dp
	1.对于某一行 matrix[i] 的某个位置 matrix[i][j] 的最小下降路径和，取决于它的
		左上、正上、右上 三者的最小下降路径和中最小的那个，再加上 matrix[i][j]
	2.状态转移方程
		dp[i][j] = matrix[i][j] + min(dp[i-1][j-1],dp[i-1][j],dp[i-1][j+1])
	3.优化
		3.1.在 j=0 的左侧和 j=n-1 的右侧，定义哨兵
		3.2.使用滚动数组代替 dp[i][j] 矩阵
*/
// leetcode submit region begin(Prohibit modification and deletion)
func minFallingPathSum(matrix [][]int) int {
	// dp
	minVal := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	minPathSum, n := 10000, len(matrix)
	dpu := make([]int, n+2) // 上一行
	dpu[0], dpu[n+1] = 10000, 10000
	copy(dpu[1:], matrix[0]) // 初始化第 0 行
	for i := 1; i < n; i++ {
		dpi := make([]int, n+2) // 当前行
		dpi[0], dpi[n+1] = 10000, 10000
		for j := 1; j <= n; j++ { // 当前行的最小下降路径和
			dpi[j] = matrix[i][j-1] + minVal(dpu[j-1], minVal(dpu[j], dpu[j+1]))
		}
		dpu = dpi // 滚动数组
	}
	for i := 1; i <= n; i++ { // 找出最后一行的最小下降路径和
		minPathSum = minVal(minPathSum, dpu[i])
	}
	return minPathSum
}

//leetcode submit region end(Prohibit modification and deletion)
