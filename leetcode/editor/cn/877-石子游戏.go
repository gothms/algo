package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func stoneGame(piles []int) bool {
	// lc：方法三
	return true

	// lc：方法二
	// dp[i][j] 表示当剩下的石子堆为下标 i 到下标 j 时，即在下标范围 [i,j] 中，当前玩家与另一个玩家的石子数量之差的最大值
	//n := len(piles)
	//dp := make([][]int, n)
	//for i := range dp {
	//	dp[i] = make([]int, n)
	//	dp[i][i] = piles[i]
	//}
	//for i := n - 2; i >= 0; i-- {
	//	for j := i + 1; j < n; j++ {
	//		dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
	//	}
	//}
	//return dp[0][n-1] > 0

	// lc：方法一
	// 0：总是记录 [i,j] 区间首先选的最大值
	// 1：总是记录 [i,j] 区间被 0 选择 i/j 后，剩余区间的最大值
	//n := len(piles)
	//dp := make([][][2]int, n)
	//for i := 0; i < n; i++ {
	//	dp[i] = make([][2]int, n)
	//}
	//for i := n - 1; i >= 0; i-- {
	//	dp[i][i][0] = piles[i]
	//	for j := i + 1; j < n; j++ {
	//		l, r := piles[i]+dp[i+1][j][1], piles[j]+dp[i][j-1][1]
	//		if l < r {
	//			dp[i][j][0], dp[i][j][1] = r, dp[i][j-1][0]
	//		} else {
	//			dp[i][j][0], dp[i][j][1] = l, dp[i+1][j][0]
	//		}
	//	}
	//}
	//return dp[0][n-1][0] > dp[0][n-1][1]

	// 个人
	//n := len(piles)
	//pre := make([]int, n+1)
	//for i, v := range piles { // 前缀和
	//	pre[i+1] = pre[i] + v
	//}
	//
	//dp := make([][][2]int, n) // 0：Alice 1：Bob
	//for i := range dp {
	//	dp[i] = make([][2]int, n)
	//	dp[i][i][1] = piles[i]
	//}
	//for i := n - 2; i >= 0; i-- {
	//	for j := i + 1; j < n; j++ {
	//		k := (j - i + 1) & 1	// 是 Alice 还是 Bob
	//		dp[i][j][k] = max(
	//			piles[i]+pre[j+1]-pre[i+1]-dp[i+1][j][k^1],
	//			piles[j]+pre[j]-pre[i]-dp[i][j-1][k^1])
	//	}
	//}
	//return dp[0][n-1][0] >= (pre[n]+1)>>1
}

//leetcode submit region end(Prohibit modification and deletion)
