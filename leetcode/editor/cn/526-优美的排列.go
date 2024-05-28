package main

import "math/bits"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countArrangement(n int) int {
	// 状压dp
	m := 1 << n
	dp := make([]int, m)
	dp[0] = 1 // 一个都没填
	for mask := 1; mask < m; mask++ {
		j := bits.OnesCount(uint(mask)) // 填数字 j
		for i := 1; i <= n; i++ {       // 填入的位置
			if k := 1 << (i - 1); k&mask > 0 && (i%j == 0 || j%i == 0) { // 可以填入
				dp[mask] += dp[mask^k]
			}
		}
	}
	return dp[m-1]

	// 记忆化搜索：优化
	//m := 1 << n
	//memo := make([]int, m)
	//for i := range memo {
	//	memo[i] = -1
	//}
	//var dfs func(int, int) int
	//dfs = func(mask, i int) int {
	//	if i == 0 {
	//		return 1
	//	}
	//	v := &memo[mask]
	//	if *v >= 0 {
	//		return *v
	//	}
	//	ret := 0
	//	for j := 1; j <= n; j++ {
	//		if mask>>(j-1)&1 > 0 && (i%j == 0 || j%i == 0) {
	//			ret += dfs(1<<(j-1)^mask, i-1)
	//		}
	//	}
	//	*v = ret
	//	return ret
	//}
	//return dfs(m-1, n)

	// 记忆化搜索
	//m := 1 << n
	//memo := make([][]int, m)
	//for i := range memo {
	//	memo[i] = make([]int, n+1)
	//	for j := range memo[i] {
	//		memo[i][j] = -1
	//	}
	//}
	//var dfs func(int, int) int
	//dfs = func(mask, i int) int {
	//	if i == 0 {
	//		return 1
	//	}
	//	v := &memo[mask][i]
	//	if *v >= 0 {
	//		return *v
	//	}
	//	ret := 0
	//	for j := 1; j <= n; j++ {
	//		if mask>>(j-1)&1 > 0 && (i%j == 0 || j%i == 0) {
	//			ret += dfs(1<<(j-1)^mask, i-1)
	//		}
	//	}
	//	*v = ret
	//	return ret
	//}
	//return dfs(m-1, n)
}

//leetcode submit region end(Prohibit modification and deletion)
