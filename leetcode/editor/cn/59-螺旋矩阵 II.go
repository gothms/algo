package main

import "fmt"

func main() {
	n := 3
	matrix := generateMatrix(n)
	fmt.Println(matrix)
}

// leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	d := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	for i, j, k, v := 0, 0, 0, 1; v <= n*n; v++ {
		ans[i][j] = v
		i += d[k][0]
		j += d[k][1]
		if j < 0 || i == n || j == n || ans[i][j] > 0 {
			i -= d[k][0]
			j -= d[k][1]
			k = (k + 1) % 4
			i += d[k][0]
			j += d[k][1]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
