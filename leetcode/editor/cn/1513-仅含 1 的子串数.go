package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numSub(s string) int {
	const mod = 1_000_000_007
	ans, cnt := 0, 0
	for _, c := range s {
		if c == '1' {
			cnt++
			ans = (ans + cnt) % mod
		} else {
			cnt = 0
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
