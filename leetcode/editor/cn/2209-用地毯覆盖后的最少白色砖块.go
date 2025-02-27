package main

import "fmt"

func main() {
	floor := "10110101"
	numCarpets, carpetLen := 2, 2
	floor = "11111"
	numCarpets, carpetLen = 2, 3
	tiles := minimumWhiteTiles(floor, numCarpets, carpetLen)
	fmt.Println(tiles)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	// 记忆化搜索
	//n := len(floor)
	//memo := make([][]int, numCarpets+1)
	//for i := range memo {
	//	memo[i] = make([]int, n)
	//	for j := range memo[i] {
	//		memo[i][j] = -1 // 初始化
	//	}
	//}
	//var dfs func(int, int) int
	//dfs = func(k, i int) int {
	//	//if i < 0 {
	//	if i < k*carpetLen { // 剪枝
	//		return 0
	//	}
	//	v := &memo[k][i]
	//	if *v != -1 {
	//		return *v
	//	}
	//	ans := dfs(k, i-1) + int(floor[i]-'0')
	//	if k > 0 { // 必须有剩余的地毯
	//		ans = min(ans, dfs(k-1, i-carpetLen))
	//	}
	//	*v = ans // 记忆化
	//	return ans
	//}
	//return dfs(numCarpets, n-1)

	// dp

	// 滚动数组：终版
	//n := len(floor)
	//dp := make([]int, n)
	//dp[0] = int(floor[0] - '0')
	//for i := 1; i < n; i++ {
	//	dp[i] += dp[i-1] + int(floor[i]-'0') // 初始化不覆盖的情况
	//}
	//d := make([]int, n)
	//for i := 1; i <= numCarpets; i++ {
	//	d[min(n-1, i*carpetLen-1)] = 0 // 防止越界
	//	for j := i * carpetLen; j < n; j++ {
	//		d[j] = min(d[j-1]+int(floor[j]-'0'), dp[j-carpetLen])
	//	}
	//	dp, d = d, dp // 滚动数组
	//}
	//return dp[n-1]

	// dp
	//n := len(floor)
	//dp := make([][]int, numCarpets+1)
	//for i := range dp {
	//	dp[i] = make([]int, n+1)
	//}
	//for i, c := range floor {
	//	dp[0][i+1] += dp[0][i] + int(c-'0')
	//}
	//for i := 1; i <= numCarpets; i++ {
	//	for j, c := range floor {
	//		v := int(c - '0')
	//		dp[i][j+1] = min(dp[i-1][max(0, j+1-carpetLen)], dp[i][j]+v)
	//	}
	//}
	//return dp[numCarpets][n]
}

//leetcode submit region end(Prohibit modification and deletion)
