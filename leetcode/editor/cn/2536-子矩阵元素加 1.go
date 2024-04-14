//给你一个正整数 n ，表示最初有一个 n x n 、下标从 0 开始的整数矩阵 mat ，矩阵中填满了 0 。
//
// 另给你一个二维整数数组 query 。针对每个查询 query[i] = [row1i, col1i, row2i, col2i] ，请你执行下述操作：
//
//
//
// 找出 左上角 为 (row1i, col1i) 且 右下角 为 (row2i, col2i) 的子矩阵，将子矩阵中的 每个元素 加 1 。也就是给所有满足
// row1i <= x <= row2i 和 col1i <= y <= col2i 的 mat[x][y] 加 1 。
//
//
// 返回执行完所有操作后得到的矩阵 mat 。
//
//
//
// 示例 1：
//
//
//
//
//输入：n = 3, queries = [[1,1,2,2],[0,0,1,1]]
//输出：[[1,1,0],[1,2,1],[0,1,1]]
//解释：上图所展示的分别是：初始矩阵、执行完第一个操作后的矩阵、执行完第二个操作后的矩阵。
//- 第一个操作：将左上角为 (1, 1) 且右下角为 (2, 2) 的子矩阵中的每个元素加 1 。
//- 第二个操作：将左上角为 (0, 0) 且右下角为 (1, 1) 的子矩阵中的每个元素加 1 。
//
//
// 示例 2：
//
//
//
//
//输入：n = 2, queries = [[0,0,1,1]]
//输出：[[1,1],[1,1]]
//解释：上图所展示的分别是：初始矩阵、执行完第一个操作后的矩阵。
//- 第一个操作：将矩阵中的每个元素加 1 。
//
//
//
// 提示：
//
//
// 1 <= n <= 500
// 1 <= queries.length <= 10⁴
// 0 <= row1i <= row2i < n
// 0 <= col1i <= col2i < n
//
//
// Related Topics 数组 矩阵 前缀和 👍 38 👎 0

package main

import "fmt"

func main() {
	q := [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}}
	n := 3
	queries := rangeAddQueries(n, q)
	fmt.Println(queries)
}

// leetcode submit region begin(Prohibit modification and deletion)
func rangeAddQueries(n int, queries [][]int) [][]int {
	// 差分数组
	diff := make([][]int, n+2)
	for i := range diff {
		diff[i] = make([]int, n+2) // 哨兵 * 2
	}
	for _, q := range queries { // 差分
		diff[q[0]+1][q[1]+1]++
		diff[q[0]+1][q[3]+2]--
		diff[q[2]+2][q[1]+1]--
		diff[q[2]+2][q[3]+2]++
	}
	for i := 1; i <= n; i++ { // 公式
		for j := 1; j <= n; j++ {
			diff[i][j] += diff[i-1][j] + diff[i][j-1] - diff[i-1][j-1]
		}
		diff[i-1] = diff[i-1][1 : n+1]
	}
	diff[n] = diff[n][1 : n+1]
	return diff[1 : n+1]

	// 差分
	//diff := make([][]int, n+2)
	//for i := range diff {
	//	diff[i] = make([]int, n+2)
	//}
	//for _, q := range queries {
	//	diff[q[0]+1][q[1]+1]++ // 左上加 1
	//	diff[q[0]+1][q[3]+2]-- // 右上
	//	diff[q[2]+2][q[1]+1]-- // 左下
	//	diff[q[2]+2][q[3]+2]++ // 右下角恢复 1
	//}
	//for i := 1; i <= n; i++ {
	//	for j := 1; j <= n; j++ {
	//		diff[i][j] += diff[i-1][j] + diff[i][j-1] - diff[i-1][j-1]
	//	}	// 前缀和公式
	//}
	//n++
	//diff = diff[1:n]
	//for i, d := range diff {
	//	diff[i] = d[1:n]
	//}
	//return diff

	// 迭代
	//grid := make([][]int, n)
	//for i := 0; i < n; i++ {
	//	grid[i] = make([]int, n)
	//}
	//for _, q := range queries {
	//	for i := q[0]; i <= q[2]; i++ {
	//		for j := q[1]; j <= q[3]; j++ {
	//			grid[i][j]++
	//		}
	//	}
	//}
	//return grid
}

//leetcode submit region end(Prohibit modification and deletion)
