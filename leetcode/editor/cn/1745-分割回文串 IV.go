package main

import "fmt"

func main() {
	s := "abcbdd"
	s = "bbab" // true
	partitioning := checkPartitioning(s)
	fmt.Println(partitioning)
}

// leetcode submit region begin(Prohibit modification and deletion)
func checkPartitioning(s string) bool {
	// 同lc-1278

	// 个人
	n := len(s)
	memo := make([][]bool, n)
	for i := n - 1; i >= 0; i-- {
		memo[i] = make([]bool, n)
		memo[i][i] = true
		for j := i + 1; j < n; j++ {
			memo[i][j] = s[i] == s[j] && (i+1 == j || memo[i+1][j-1])
		}
	}
	dp := make([][3]bool, n)
	for i := 0; i < n; i++ {
		dp[i][0] = memo[0][i]
		for j := i; j > 0; j-- {
			if memo[j][i] {
				for k := 1; k <= 2; k++ {
					dp[i][k] = dp[i][k] || dp[j-1][k-1]
				}
			}
		}
	}
	return dp[n-1][2]
}

//leetcode submit region end(Prohibit modification and deletion)
