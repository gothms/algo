package array

/*
两条对角线，四种遍历顺序

lc
	498
*/

// 撇：右上 -> 左下
// i0 j0 初始位置为：﹁
func pieFromTopRight(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	ans := make([][]int, m+n-1)
	for k := range ans {
		i0 := max(0, k-n+1)
		j0 := min(n-1, k)
		//ans[k] = make([]int, 0, j0-max(0, j0-m+i0+1)+1)
		for j := j0; j >= max(0, j0-m+i0+1); j-- { // 更直观的写法：分别记录 i j，判断 i<m && j>=0
			ans[k] = append(ans[k], mat[i0+j0-j][j])
		}
	}
	return ans
}

// 撇：左下 -> 右上
// i0 j0 初始位置为：﹂
func pieFromBottomLeft(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	ans := make([][]int, m+n-1)
	for k := range ans {
		i0 := min(m-1, k)
		j0 := max(0, k-m+1)
		//ans[k] = make([]int, 0, min(n, j0+i0+1)-j0)
		for j := j0; j < min(n, j0+i0+1); j++ {
			ans[k] = append(ans[k], mat[i0-j+j0][j])
		}
	}
	return ans
}

// 捺：左上 -> 右下
// i0 j0 初始位置为：「
func naFromTopLeft(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	ans := make([][]int, m+n-1)
	for k := range ans {
		i0 := max(0, k-n+1)
		j0 := max(0, n-k-1)
		//ans[k] = make([]int, 0, min(n, j0+m-i0)-j0)
		for j := j0; j < min(n, j0+m-i0); j++ {
			ans[k] = append(ans[k], mat[i0+j-j0][j])
		}
	}
	return ans
}

// 捺：右下 -> 左上
// i0 j0 初始位置为：」
func naFromBottomRight(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	ans := make([][]int, m+n-1)
	for k := range ans {
		i0 := min(m-1, k)
		j0 := min(n-1, m+n-k-2)
		//ans[k] = make([]int, 0, j0-max(0, j0-i0)+1)
		for j := j0; j >= max(0, j0-i0); j-- {
			ans[k] = append(ans[k], mat[i0-j0+j][j])
		}
	}
	return ans
}
