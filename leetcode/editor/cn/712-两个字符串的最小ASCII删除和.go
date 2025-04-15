package main

import "fmt"

func main() {
	s1 := "sea"
	s2 := "eat"
	s1 = "delete"
	s2 = "leet"
	deleteSum := minimumDeleteSum(s1, s2)
	fmt.Println(deleteSum)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumDeleteSum(s1 string, s2 string) int {
	// 个人
	m, n := len(s1), len(s2)
	if m > n {
		s1, s2, m, n = s2, s1, n, m
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := range dp[0][1:] {
		dp[0][i+1] = dp[0][i] + int(s2[i])
	}
	for i, x := range s1 {
		dp[i+1][0] = dp[i][0] + int(s1[i])
		for j, y := range s2 {
			if x == y {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j+1]+int(x), dp[i+1][j]+int(y))
			}
		}
	}
	return dp[m][n]
}

//leetcode submit region end(Prohibit modification and deletion)
