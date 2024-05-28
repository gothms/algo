package main

import (
	"fmt"
)

func main() {
	s := "ccc"
	s = "a"
	s = "bananas"     // anana
	s = "dabbac"      // abba
	s = "cbbd"        // bb
	s = "babad"       // bab
	s = "aacabdkacaa" // aca
	palindrome := longestPalindrome(s)
	fmt.Println(palindrome)
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	// manacher

}

//leetcode submit region end(Prohibit modification and deletion)

func longestPalindrome_(s string) string {
	// manacher：E:\gothmslee\algo\string\manacher.go

	// dp
	maxL, idx, n := 1, 0, len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 1; i < n; i++ {
		dp[i][i] = 1
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] && (j == i-1 || dp[i-1][j+1] > 0) {
				dp[i][j] = dp[i-1][j+1] + 2
			}
			if dp[i][j] > maxL {
				maxL, idx = dp[i][j], j
			}
		}
	}
	return s[idx : idx+maxL]

	// 中心扩散算法
}
