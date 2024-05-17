package dp

/*
滚动数组

lc
	516
*/

// 模板：滚动数组
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = 1
		pre := 0 // dp[i+1][i]
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				pre, dp[j] = dp[j], pre+2
			} else {
				pre = dp[j]
				dp[j] = max(dp[j], dp[j-1])
			}
		}
	}
	return dp[n-1]
}
