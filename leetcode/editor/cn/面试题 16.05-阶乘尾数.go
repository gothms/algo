package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func trailingZeroes(n int) int {
	// max(尾2, 尾5) + 后缀0
	// 思路转换：
	ans := 0
	for n > 0 {
		n /= 5
		ans += n
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
