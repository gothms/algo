package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	ans := make([][]int, r)
	i, j := 0, 0
	for a := range ans {
		ans[a] = make([]int, c)
		for b := range ans[a] {
			ans[a][b] = mat[i][j]
			if j++; j == n {
				i++
				j = 0
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
