//矩阵对角线 是一条从矩阵最上面行或者最左侧列中的某个元素开始的对角线，沿右下方向一直到矩阵末尾的元素。例如，矩阵 mat 有 6 行 3 列，从 mat[2
//][0] 开始的 矩阵对角线 将会经过 mat[2][0]、mat[3][1] 和 mat[4][2] 。
//
// 给你一个 m * n 的整数矩阵 mat ，请你将同一条 矩阵对角线 上的元素按升序排序后，返回排好序的矩阵。
//
//
//
// 示例 1：
//
//
//
//
//输入：mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
//输出：[[1,1,1,1],[1,2,2,2],[1,2,3,3]]
//
//
// 示例 2：
//
//
//输入：mat = [[11,25,66,1,69,7],[23,55,17,45,15,52],[75,31,36,44,58,8],[22,27,33,2
//5,68,4],[84,28,14,11,5,50]]
//输出：[[5,17,4,1,52,7],[11,11,25,45,8,69],[14,23,25,44,58,15],[22,27,31,36,50,66]
//,[84,28,75,33,55,68]]
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 100
// 1 <= mat[i][j] <= 100
//
//
// Related Topics 数组 矩阵 排序 👍 93 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func diagonalSort(mat [][]int) [][]int {
	// 个人
	n, m := len(mat), len(mat[0])
	insertSort := func(s, d int) { // 插入排序
		for i := s + 1; i < n && i-d < m; i++ {
			j, v := i-1, mat[i][i-d]
			for ; j >= s && mat[j][j-d] > v; j-- {
				mat[j+1][j+1-d] = mat[j][j-d]
			}
			mat[j+1][j+1-d] = v
		}
	}
	for d := 2 - m; d <= 0; d++ { // 起点是 [0, j]
		insertSort(0, d)
	}
	for d := 1; d < n-1; d++ { // 起点是 [i, 0]
		insertSort(d, d)
	}
	return mat

	// lc
	//n := len(mat)
	//m := len(mat[0])
	//diag := make([][]int, m+n)
	//for i := 0; i < n; i++ {
	//	for j := 0; j < m; j++ {
	//		diag[i-j+m] = append(diag[i-j+m], mat[i][j])
	//	}
	//}
	//for i, d := range diag {
	//	fmt.Println(i, i, d)
	//}
	//for _, d := range diag {
	//	sort.Sort(sort.Reverse(sort.IntSlice(d)))
	//}
	//for i, d := range diag {
	//	fmt.Println(i, d)
	//}
	//for i := 0; i < n; i++ {
	//	for j := 0; j < m; j++ {
	//		mat[i][j] = diag[i-j+m][len(diag[i-j+m])-1]
	//		diag[i-j+m] = diag[i-j+m][:len(diag[i-j+m])-1]
	//	}
	//}
	//return mat
}

//leetcode submit region end(Prohibit modification and deletion)
