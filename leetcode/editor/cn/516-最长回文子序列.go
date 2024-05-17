package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1)
	b := big.Int{}
	fmt.Printf("%T,%T\n", a, b)

	s := "bbbab"
	s = "aabaaba" // 6
	subseq := longestPalindromeSubseq(s)
	fmt.Println(subseq)
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindromeSubseq(s string) int {
	// dp：可滚动数组优化
	n := len(s)
	dp := make([]int, n)
	for i := range s {
		dp[i] = 1
		pre := 0
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				dp[j], pre = pre+2, dp[j]
			} else {
				pre = dp[j]
				dp[j] = max(dp[j], dp[j+1])
			}
		}
	}
	return dp[0]

	//ans, n := 1, len(s)
	//dp := make([][]int, n)
	//for i := range dp {
	//	dp[i] = make([]int, n)
	//}
	//for i := range s {
	//	dp[i][i] = 1
	//	for j := i - 1; j >= 0; j-- {
	//		//if s[i] == s[j] && (j == i-1 || dp[i-1][j+1] > 0) {
	//		if s[i] == s[j] {
	//			dp[i][j] = dp[i-1][j+1] + 2 // 状态转移
	//		} else {
	//			dp[i][j] = max(dp[i][j+1], dp[i-1][j]) // 状态转移
	//		}
	//		ans = max(ans, dp[i][j])
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
