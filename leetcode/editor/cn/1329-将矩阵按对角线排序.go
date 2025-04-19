package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func diagonalSort(mat [][]int) [][]int {
	// 捺：左上 -> 右下，k 从 [0,n-1] 开始
	//i0 = max(0, k-n+1)
	//j0 = max(0, n-k-1)
	//cur = 0
	//for j := j0; j < min(n, j0+m-i0); j++ {
	//	//fmt.Println(i0, j0, j0+m-i0, j, i0+j0-j)
	//	cur = cur*10 + mat[i0+j-j0][j]
	//	memo[cur]++
	//}
	//// 捺：右下 -> 左上，k 从 [0,n-1] 开始
	//i0 = min(m-1, k)
	//j0 = min(n-1, m+n-k-2)
	//cur = 0
	//for j := j0; j >= max(0, j0-i0); j-- {
	//	//fmt.Println(i0, j0, j, i0-j0+j)
	//	cur = cur*10 + mat[i0-j0+j][j]
	//	memo[cur]++
	//}

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
