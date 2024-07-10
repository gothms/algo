package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func modifiedMatrix(matrix [][]int) [][]int {
	n := len(matrix[0])
	col := make([]int, n)
	//for i := range col {
	//	col[i] = -1
	//}
	for _, m := range matrix {
		for j, v := range m {
			col[j] = max(col[j], v)
		}
	}
	for i, m := range matrix {
		for j, v := range m {
			if v == -1 {
				matrix[i][j] = col[j]
			}
		}
	}
	return matrix
}

//leetcode submit region end(Prohibit modification and deletion)
