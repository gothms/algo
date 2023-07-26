//给你一个 2 行 n 列的二进制数组：
//
//
// 矩阵是一个二进制矩阵，这意味着矩阵中的每个元素不是 0 就是 1。
// 第 0 行的元素之和为 upper。
// 第 1 行的元素之和为 lower。
// 第 i 列（从 0 开始编号）的元素之和为 colsum[i]，colsum 是一个长度为 n 的整数数组。
//
//
// 你需要利用 upper，lower 和 colsum 来重构这个矩阵，并以二维整数数组的形式返回它。
//
// 如果有多个不同的答案，那么任意一个都可以通过本题。
//
// 如果不存在符合要求的答案，就请返回一个空的二维数组。
//
//
//
// 示例 1：
//
// 输入：upper = 2, lower = 1, colsum = [1,1,1]
//输出：[[1,1,0],[0,0,1]]
//解释：[[1,0,1],[0,1,0]] 和 [[0,1,1],[1,0,0]] 也是正确答案。
//
//
// 示例 2：
//
// 输入：upper = 2, lower = 3, colsum = [2,2,1,1]
//输出：[]
//
//
// 示例 3：
//
// 输入：upper = 5, lower = 5, colsum = [2,1,2,0,1,0,1,2,0,1]
//输出：[[1,1,1,0,1,0,0,1,0,0],[1,0,1,0,0,0,1,1,0,1]]
//
//
//
//
// 提示：
//
//
// 1 <= colsum.length <= 10^5
// 0 <= upper, lower <= colsum.length
// 0 <= colsum[i] <= 2
//
//
// Related Topics 贪心 数组 矩阵 👍 30 👎 0

package main

import "fmt"

func main() {
	upper := 5
	lower := 5
	colsum := []int{2, 1, 2, 0, 1, 0, 1, 2, 0, 1}
	matrix := reconstructMatrix(upper, lower, colsum)
	fmt.Println(matrix)
}

/*
思路：贪心
	1.对于任意 colsum[i]，3种情况
		0：不处理
		1：matrix[0]/matrix[1] = 1
		2：matrix[0], matrix[1] = 1, 1
	2.失败条件为：colsum[i] 的和 != upper+lower，或者 2 的总数 > upper/lower
	3.贪心
		遍历 colsum，遇到 2 就先处理，并 upper-- 和 lower--，用于计算失败条件
		遇到 1 则先统计 1 的数量，且先满足 matrix[0]，后满足 matrix[1]
*/
// leetcode submit region begin(Prohibit modification and deletion)
func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	// 贪心
	sum, n := 0, len(colsum)
	matrix := make([][]int, 2)
	matrix[0], matrix[1] = make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		switch colsum[i] {
		case 2: // 处理 2
			upper--
			lower--
			matrix[0][i], matrix[1][i] = 1, 1
		case 1:
			sum++
			matrix[0][i] = 1 // 贪心，第0行先填充 1
		}
	}
	if upper < 0 || lower < 0 || sum != upper+lower {
		return nil // 2 超标 / 1 总数不符
	}
	for i := n - 1; lower > 0; i-- { // 填充第1行的 1
		if colsum[i] == 1 {
			matrix[0][i], matrix[1][i] = 0, 1
			lower--
		}
	}
	return matrix
}

//leetcode submit region end(Prohibit modification and deletion)
