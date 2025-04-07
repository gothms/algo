package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, col := make([]bool, m), make([]bool, n)
	for i, mat := range matrix {
		for j, v := range mat {
			if v == 0 {
				row[i], col[j] = true, true
			}
		}
	}
	for i, mat := range matrix {
		for j := range mat {
			if row[i] || col[j] {
				mat[j] = 0
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
