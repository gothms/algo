package main

import "fmt"

func main() {
	s, t := "rabbbit", "rabbit"
	s, t = "babgbag", "bag"
	s, t = "ddd", "dd" // 3
	distinct := numDistinct(s, t)
	fmt.Println(distinct)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numDistinct(s string, t string) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func numDistinct_(s string, t string) int {
	// dp
	m, n := len(s), len(t)
	dp := make([]int, n+1)
	dp[0] = 1 // 哨兵
	for i := 0; i < m; i++ {
		for j := min(i, n-1); j >= 0; j-- {
			if s[i] == t[j] {
				dp[j+1] += dp[j]
			}
		}
	}
	return dp[n]
}
