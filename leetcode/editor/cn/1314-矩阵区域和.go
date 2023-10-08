//给你一个 m x n 的矩阵 mat 和一个整数 k ，请你返回一个矩阵 answer ，其中每个 answer[i][j] 是所有满足下述条件的元素
//mat[r][c] 的和：
//
//
// i - k <= r <= i + k,
// j - k <= c <= j + k 且
// (r, c) 在矩阵内。
//
//
//
//
// 示例 1：
//
//
//输入：mat = [[1,2,3],[4,5,6],[7,8,9]], k = 1
//输出：[[12,21,16],[27,45,33],[24,39,28]]
//
//
// 示例 2：
//
//
//输入：mat = [[1,2,3],[4,5,6],[7,8,9]], k = 2
//输出：[[45,45,45],[45,45,45],[45,45,45]]
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n, k <= 100
// 1 <= mat[i][j] <= 100
//
//
// Related Topics 数组 矩阵 前缀和 👍 182 👎 0

package main

import "fmt"

func main() {
	mat := [][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	k := 2
	blockSum := matrixBlockSum(mat, k)
	fmt.Println(blockSum)
}

// leetcode submit region begin(Prohibit modification and deletion)
func matrixBlockSum(mat [][]int, k int) [][]int {
	// dp：前缀和
	// TODO
	return nil

	// 遍历
	/*minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	m, n := len(mat), len(mat[0])
	ret := make([][]int, m)
	for i := 0; i < m; i++ {
		ret[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			mat[i][j] += mat[i][j-1]
		}
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] += mat[i-1][j]
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			xu, yu := minVal(m-1, i+k), minVal(n-1, j+k)
			xd, yd := i-k-1, j-k-1
			ret[i][j] = mat[xu][yu] // i,j 为中心
			if xd >= 0 && yd >= 0 { // 加左上
				ret[i][j] += mat[xd][yd]
			}
			if xd >= 0 { // 减上
				ret[i][j] -= mat[xd][yu]
			}
			if yd >= 0 { // 减左
				ret[i][j] -= mat[xu][yd]
			}
		}
	}
	return ret*/
}

//leetcode submit region end(Prohibit modification and deletion)
