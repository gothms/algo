package main

import "fmt"

func main() {
	s := "aab"
	s = "a"
	cut := minCut(s)
	fmt.Println(cut)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minCut(s string) int {
	n := len(s)
	memo := make([][]bool, n)
	for i := n - 1; i >= 0; i-- {
		memo[i] = make([]bool, n)
		memo[i][i] = true
		for j := i + 1; j < n; j++ {
			memo[i][j] = s[i] == s[j] && (i+1 == j || memo[i+1][j-1])
		}
	}
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		if memo[0][i] {
			continue
		}
		dp[i] = n
		for j := 1; j <= i; j++ {
			if memo[j][i] {
				dp[i] = min(dp[i], dp[j-1]+1)
			}
		}
	}
	return dp[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)

func minCut(s string) int {
	n := len(s)
	memo := make([][]bool, n)
	for i := n - 1; i >= 0; i-- { // 预处理
		memo[i] = make([]bool, n)
		memo[i][i] = true
		for j := i + 1; j < n; j++ { // 记录回文串
			memo[i][j] = s[j] == s[i] && (j == i+1 || memo[i+1][j-1])
		}
	}
	dp := make([]int, n)
	for i := 1; i < n; i++ { // 处理边界问题是一个难点
		if memo[0][i] { // [0,i] 为回文串，则 dp[i]=0
			continue
		}
		dp[i] = n // 最大为 n-1
		for j := i; j > 0; j-- {
			if memo[j][i] {
				dp[i] = min(dp[i], dp[j-1]+1) // 尝试更新
			}
		}
	}
	return dp[n-1]
}
