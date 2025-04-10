package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type NumMatrix struct {
	s [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	s := make([][]int, len(matrix)+1)
	n := len(matrix[0]) + 1
	s[0] = make([]int, n)
	for i, m := range matrix {
		s[i+1] = make([]int, n)
		for j, v := range m {
			s[i+1][j+1] = s[i+1][j] + s[i][j+1] - s[i][j] + v
		}
	}
	return NumMatrix{s}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.s[row2+1][col2+1] -
		this.s[row1][col2+1] -
		this.s[row2+1][col1] +
		this.s[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
//leetcode submit region end(Prohibit modification and deletion)
