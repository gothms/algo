package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func monkeyMove(n int) int {
	// 快速幂
	const mod = 1_000_000_007
	return (func(a int) int {
		res, base := 1, 2
		for ; a != 0; a >>= 1 {
			if a&1 != 0 {
				res = res * base % mod
			}
			base = base * base % mod
		}
		return res
	}(n) - 2 + mod) % mod
}

//leetcode submit region end(Prohibit modification and deletion)
