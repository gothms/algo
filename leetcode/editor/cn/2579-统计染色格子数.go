package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func coloredCells(n int) int64 {
	// lc：设答案为 f(n)，那么 f(n) 相当于在 f(n−1) 的基础上多了 4 组 n−1 的格子
	// f(n) = f(n-1) + 4(n-1)
	// f(n)=1+4(1+2+⋯n−1)=1+2n(n−1)

	// 个人
	b := int64(n<<1 - 1)
	return b*b - int64((n*(n-1)))<<1
}

//leetcode submit region end(Prohibit modification and deletion)
