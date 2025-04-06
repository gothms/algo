package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countOdds(low int, high int) int {
	// lc
	cnt := func(n int) int {
		return (n + 1) >> 1
	}
	return cnt(high) - cnt(low-1)

	// 个人
	//return (high-low-low&1-high&1)>>1 + low&1 + high&1
}

//leetcode submit region end(Prohibit modification and deletion)
