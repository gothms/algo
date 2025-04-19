package main

import "fmt"

func main() {
	mat := [][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	mat = [][]int{{1, 2, 3, 4},
		{5, 6, 7, 8}}
	order := findDiagonalOrder(mat)
	fmt.Println(order)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, m)
	}
	for i, row := range mat {
		for j, v := range row {
			arr[j][m-i-1] = v
		}
	}
	ans := make([]int, 0, m*n)
	m, n = len(arr), len(arr[0])
	for k := 0; k < m+n-1; k++ {
		switch k & 1 {
		case 0: // 左上 -> 右下
			i0 := max(0, k-n+1)
			j0 := max(0, n-k-1)
			for j := j0; j < min(n, j0+m-i0); j++ {
				ans = append(ans, arr[i0+j-j0][j])
			}
		case 1: // 右下 -> 左上
			i0 := min(m-1, k)
			j0 := min(n-1, m+n-k-2)
			for j := j0; j >= max(0, j0-i0); j-- {
				ans = append(ans, arr[i0-j0+j][j])
			}
		}
	}
	return ans

	// 个人
	//m, n := len(mat), len(mat[0])
	//ans := make([]int, 0, m*n)
	//for k := 0; k < m+n-1; k++ {
	//	switch k & 1 {
	//	case 0: // 左下 -> 右上
	//		i0 := min(m-1, k)
	//		j0 := max(0, k-m+1)
	//		for j := j0; j < min(n, j0+i0+1); j++ {
	//			ans = append(ans, mat[i0-j+j0][j])
	//		}
	//	case 1: // 右上 -> 左下
	//		i0 := max(0, k-n+1)
	//		j0 := min(n-1, k)
	//		for j := j0; j >= max(0, j0-m+i0+1); j-- {
	//			ans = append(ans, mat[i0+j0-j][j])
	//		}
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
