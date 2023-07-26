//给定一个方阵，其中每个单元(像素)非黑即白。设计一个算法，找出 4 条边皆为黑色像素的最大子方阵。
//
// 返回一个数组 [r, c, size] ，其中 r, c 分别代表子方阵左上角的行号和列号，size 是子方阵的边长。若有多个满足条件的子方阵，返回 r
//最小的，若 r 相同，返回 c 最小的子方阵。若无满足条件的子方阵，返回空数组。
//
// 示例 1:
//
// 输入:
//[
//   [1,0,1],
//   [0,0,1],
//   [0,0,1]
//]
//输出: [1,0,2]
//解释: 输入中 0 代表黑色，1 代表白色，标粗的元素即为满足条件的最大子方阵
//
//
// 示例 2:
//
// 输入:
//[
//   [0,1,1],
//   [1,0,1],
//   [1,1,0]
//]
//输出: [0,0,1]
//
//
// 提示：
//
//
// matrix.length == matrix[0].length <= 200
//
//
// Related Topics 数组 动态规划 矩阵 👍 53 👎 0

package main

func main() {
	//[
	//[1, 1, 1, 1, 0, 1, 0, 1, 1, 1],
	//[1, 1, 0, 0, 0, 0, 0, 0, 0, 0],
	//[1, 1, 1, 1, 0, 1, 0, 1, 0, 1],
	//[1, 1, 1, 1, 0, 0, 0, 0, 0, 0],
	//[1, 0, 1, 0, 1, 1, 1, 1, 1, 1],
	//[1, 1, 0, 0, 1, 0, 1, 0, 0, 1],
	//[0, 0, 0, 1, 1, 1, 0, 1, 0, 1],
	//[0, 0, 0, 1, 0, 1, 0, 1, 0, 1],
	//[1, 0, 1, 0, 1, 1, 0, 1, 1, 1],
	//[1, 1, 1, 0, 1, 0, 0, 1, 1, 1]]
}

/*
思路：dp
	1.注意题意是，最大子方阵的4条边全部为 0，而“中心”地带可以为 1
	2.对于任意 i j，假设行 i 的最长连续 0 长度为 x，列 j 的最长连续 0 长度为 y
		那么最大子方阵的最大边长为 edge <= minVal(x, y)
		且满足
			行 i+1-edge 的最长连续 0 长度 >= edge
			列 j+1-edge 的最长连续 0 长度 >= edge
	3.两个二维数组 row col 分别记录 matrix 在 i j 位置的最大连续 0 长度
		然后遍历 [0,edge]，检查 row[i+1-edge][j] 和 col[i][j+1-edge] 是否 >= edge
		同时满足时，edge 即为 i j 当前最大子方阵的边长
*/
// leetcode submit region begin(Prohibit modification and deletion)
func findSquare(matrix [][]int) []int {
	// dp
	minVal := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	r, c, size, n := 0, 0, 0, len(matrix)
	row, col := make([][]int, n+1), make([][]int, n+1) // 行和列最大连续 0 的长度
	row[0], col[0] = make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		row[i], col[i] = make([]int, n+1), make([]int, n+1)
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == 0 {
				row[i][j], col[i][j] = row[i][j-1]+1, col[i-1][j]+1
				edge := minVal(row[i][j], col[i][j]) // 可能的最大子方阵边长
				for row[i+1-edge][j] < edge || col[i][j+1-edge] < edge {
					edge-- // 求出当前最大子方阵边长
				}
				if edge > size { // 记录最大的最大子方阵
					r, c, size = i, j, edge
				}
			}
		}
	}
	if size > 0 { // 存在最大子方阵
		return []int{r - size, c - size, size}
	}
	return nil

	// dp：这个解法的中心不能为 1，题意是方针的中心可以为 1
	//minVal := func(a, b int) int {
	//	if a < b {
	//		return a
	//	}
	//	return b
	//}
	//r, c, size, n := 0, 0, 0, len(matrix)
	//dp := make([][]int, n+1)
	//dp[0] = make([]int, n+1)
	//for i := 0; i < n; i++ {
	//	dp[i+1] = make([]int, n+1)
	//	for j := 0; j < n; j++ {
	//		if matrix[i][j] == 0 {
	//			v := minVal(dp[i][j], minVal(dp[i][j+1], dp[i+1][j])) + 1
	//			if v > size {
	//				r, c, size = i, j, v
	//			}
	//			dp[i+1][j+1] = v
	//		}
	//	}
	//}
	//if size > 0 {
	//	return []int{r - size + 1, c - size + 1, size}
	//}
	//return nil
}

//leetcode submit region end(Prohibit modification and deletion)
