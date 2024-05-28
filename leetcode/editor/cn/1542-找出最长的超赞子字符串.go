package main

import (
	"fmt"
)

func main() {
	s := "213123" // 6
	s = "940884"  // 5
	awesome := longestAwesome(s)
	fmt.Println(awesome)
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestAwesome(s string) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func longestAwesome_(s string) int {
	// 前缀异或和
	const n = 10
	m := len(s)
	dp := [1 << n]int{}
	for i := range dp {
		dp[i] = m
	}
	dp[0] = -1
	ans, b := 0, 0
	for i, c := range s {
		b ^= 1 << int(c-'0')
		for j := 0; j < n; j++ {
			ans = max(ans, i-dp[b^1<<j])
		}
		ans = max(ans, i-dp[b])
		if dp[b] == m {
			dp[b] = i
		}
	}
	return ans

	// 超时
	//ans := 0
	//for i := range s {
	//	bit := 0
	//	for j := i; j < len(s); j++ {
	//		bit ^= 1 << (s[j] - '0')
	//		if bit&(bit-1) == 0 {
	//			ans = max(ans, j-i)
	//		}
	//	}
	//}
	//return ans + 1

	// 暴力：超时
	//ans := 0
	//for i := range s {
	//	memo := make(map[uint8]int)
	//	cnt := 0
	//	for j := i; j < len(s); j++ {
	//		memo[s[j]]++
	//		if memo[s[j]]&1 == 0 {
	//			cnt--
	//		} else {
	//			cnt++
	//		}
	//		if cnt <= 1 {
	//			ans = max(ans, j-i+1)
	//		}
	//	}
	//}
	//return ans
}
