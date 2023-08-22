//给你一个正方形矩阵 mat，请你返回矩阵对角线元素的和。
//
// 请你返回在矩阵主对角线上的元素和副对角线上且不在主对角线上元素的和。
//
//
//
// 示例 1：
//
//
//
//
//输入：mat = [[1,2,3],
//            [4,5,6],
//            [7,8,9]]
//输出：25
//解释：对角线的和为：1 + 5 + 9 + 3 + 7 = 25
//请注意，元素 mat[1][1] = 5 只会被计算一次。
//
//
// 示例 2：
//
//
//输入：mat = [[1,1,1,1],
//            [1,1,1,1],
//            [1,1,1,1],
//            [1,1,1,1]]
//输出：8
//
//
// 示例 3：
//
//
//输入：mat = [[5]]
//输出：5
//
//
//
//
// 提示：
//
//
// n == mat.length == mat[i].length
// 1 <= n <= 100
// 1 <= mat[i][j] <= 100
//
//
// Related Topics 数组 矩阵 👍 74 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func diagonalSum(mat [][]int) int {
	// 优化
	sum, i, j := 0, 0, len(mat)-1
	for ; i < j; i, j = i+1, j-1 {
		sum += mat[i][i] + mat[j][j] + mat[i][j] + mat[j][i]
	}
	if i == j {
		sum += mat[i][i]
	}
	return sum

	//sum, n := 0, len(mat)
	//for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
	//	sum += mat[i][i] + mat[i][j]
	//}
	//if n&1 > 0 {
	//	return sum - mat[n>>1][n>>1]
	//}
	//return sum
}

//leetcode submit region end(Prohibit modification and deletion)
