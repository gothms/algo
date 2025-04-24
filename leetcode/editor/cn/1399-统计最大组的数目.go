package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countLargestGroup(n int) int {
	// 数位 dp

	// 暴力
	memo := make(map[int]int)
	mx := 0
	for i := 1; i <= n; i++ {
		v := 0
		for j := i; j > 0; j /= 10 {
			v += j % 10
		}
		memo[v]++
		mx = max(mx, memo[v])
	}
	ans := 0
	for _, v := range memo {
		if v == mx {
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
