//给定一个正整数、负整数和 0 组成的 N × M 矩阵，编写代码找出元素总和最大的子矩阵。
//
// 返回一个数组 [r1, c1, r2, c2]，其中 r1, c1 分别代表子矩阵左上角的行号和列号，r2, c2 分别代表右下角的行号和列号。若有多个满
//足条件的子矩阵，返回任意一个均可。
//
// 注意：本题相对书上原题稍作改动
//
// 示例：
//
// 输入：
//[
//   [-1,0],
//   [0,-1]
//]
//输出：[0,1,0,1]
//解释：输入中标粗的元素即为输出所表示的矩阵
//
//
//
// 说明：
//
//
// 1 <= matrix.length, matrix[0].length <= 200
//
//
// Related Topics 数组 动态规划 矩阵 前缀和 👍 193 👎 0

package main

import "fmt"

func main() {
	matrix := [][]int{
		{-1, -2, -9},
		{8, -9, 10},
		{7, -7, 9}}
	maxMatrix := getMaxMatrix(matrix)
	fmt.Println(maxMatrix)
}

// leetcode submit region begin(Prohibit modification and deletion)
func getMaxMatrix(matrix [][]int) []int {
	// dp + 前缀和：O(n^4)，dp 记录矩阵和
	// dp[0][0] = m[0][0]
	// dp[i][0] = dp[i-1][0] + m[i][0]
	// dp[0][j] = dp[0][j-1] + m[0][j]
	// dp[i][j] = m[i][j] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]

	// dp：O(n^3)
	n, m := len(matrix), len(matrix[0])
	max, u, l, d, r := matrix[0][0], 0, 0, 0, 0
	for s := 0; s < n; s++ {
		sumCol := make([]int, m) // 从每一行开始，sumCol[j] 记录每一列的当前和
		for i := s; i < n; i++ {
			dp, left := -1, 0 // dp：矩阵和
			for j := 0; j < m; j++ {
				sumCol[j] += matrix[i][j]
				if dp < 0 { // 重要逻辑：取当前列
					dp, left = sumCol[j], j // 并记录左上的列索引
				} else { // 累加列
					dp += sumCol[j]
				}
				if dp > max { // 更新最大矩阵
					max, u, l, d, r = dp, s, left, i, j
				}
			}
		}
	}
	return []int{u, l, d, r}
}

//leetcode submit region end(Prohibit modification and deletion)
