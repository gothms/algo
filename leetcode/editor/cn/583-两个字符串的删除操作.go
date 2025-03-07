package main

import "fmt"

func main() {
	w1 := "leetcode"
	w2 := "etco"
	//w1 = "a"
	//w2 = "a"
	distance := minDistance(w1, w2)
	fmt.Println(distance)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {
	// 优化
	m := len(word2)
	dp := make([]int, m+1)
	for i := range dp {
		dp[i] = i
	}
	for _, c1 := range word1 {
		pre := dp[0]
		dp[0]++ // 重点 1
		for j, c2 := range word2 {
			if c1 == c2 {
				dp[j+1], pre = pre, dp[j+1] // 重点 2
			} else {
				dp[j+1], pre = min(dp[j], dp[j+1])+1, dp[j+1] // 重点 3
			}
		}
	}
	return dp[m]

	// 个人
	//n, m := len(word1), len(word2)
	//dp := make([][]int, n+1)
	//for i := range dp {
	//	dp[i] = make([]int, m+1)
	//}
	//for i, c1 := range word1 { // 计算最长公共子序列
	//	for j, c2 := range word2 {
	//		if c1 == c2 {
	//			dp[i+1][j+1] = dp[i][j] + 1
	//		} else {
	//			dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
	//		}
	//	}
	//}
	//return n + m - dp[n][m]<<1
}

//leetcode submit region end(Prohibit modification and deletion)
