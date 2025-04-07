package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n>>1; i++ {
		for j := i; j < n-1-i; j++ {
			i2, j2 := j, n-1-i
			i3, j3 := n-1-i, n-1-j
			i4, j4 := j3, i
			//i4, j4 := n-1-j, i // 另一种写法
			matrix[i][j], matrix[i2][j2], matrix[i3][j3], matrix[i4][j4] =
				matrix[i4][j4], matrix[i][j], matrix[i2][j2], matrix[i3][j3]
		}
	}

	// 方法二：先左右对调，再按“撇”对角线翻转。或者先水平翻转，再“捺”对角线翻转
}

//leetcode submit region end(Prohibit modification and deletion)
