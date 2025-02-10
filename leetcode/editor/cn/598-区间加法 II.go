package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxCount(m int, n int, ops [][]int) int {
	x, y := m, n
	for _, o := range ops {
		x = min(x, o[0])
		y = min(y, o[1])
	}
	return x*y
}

//leetcode submit region end(Prohibit modification and deletion)
