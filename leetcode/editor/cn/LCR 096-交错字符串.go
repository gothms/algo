package main

import "fmt"

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	s1 = "a"
	s2 = ""
	s3 = "a"
	s1 = "aa"
	s2 = "ab"
	s3 = "abaa"
	interleave := isInterleave(s1, s2, s3)
	fmt.Println(interleave)
}

// leetcode submit region begin(Prohibit modification and deletion)
func isInterleave(s1 string, s2 string, s3 string) bool {
	// dp
	n, m, l := len(s1), len(s2), len(s3)
	if n+m != l { // fast fail
		return false
	}
	if n == 0 || m == 0 { // fast path
		return s1 == s3 || s2 == s3
	}
	dp := make([][]bool, n+1) // 也可以滚动数组
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true     // 空+空=空
	for j := range s2 { // 第 0 行
		if !dp[0][j] || s2[j] != s3[j] {
			break
		}
		dp[0][j+1] = true
	}
	for i := 0; i < n; i++ {
		if s1[i] == s3[i] { // 第 0 列
			dp[i+1][0] = dp[i][0]
		}
		for j := 0; j < m; j++ { // 状态转移方程：尾字符匹配 s3 vs s1 或 s2
			dp[i+1][j+1] = dp[i][j+1] && s1[i] == s3[i+j+1] || dp[i+1][j] && s2[j] == s3[i+j+1]
		}
	}
	return dp[n][m]
}

//leetcode submit region end(Prohibit modification and deletion)
