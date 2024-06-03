package main

import "fmt"

func main() {
	s := "aba"  // 6
	s = "aaa"   // 3
	s = "adba"  // 14
	s = "adbae" // 29
	s = "lee"   // 5
	ii := distinctSubseqII(s)
	fmt.Println(ii)
}

// leetcode submit region begin(Prohibit modification and deletion)
func distinctSubseqII(s string) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func distinctSubseqII_(s string) int {
	// dp + 前缀和
	//const mod = 1_000_000_007
	//memo := [26]int{}
	//for i := range memo {
	//	memo[i] = -1
	//}
	//dp := make([]int, len(s)+1)
	//for i, c := range s {
	//	idx := c - 'a'
	//	if memo[idx] == -1 {
	//		dp[i+1] = (dp[i]<<1 + 1) % mod
	//	} else {
	//		dp[i+1] = (dp[i]<<1 - dp[memo[idx]] + mod) % mod
	//	}
	//	memo[idx] = i
	//}
	//return dp[len(s)]

	// lc
	const mod = 1_000_000_007
	memo := [26]int{}
	ans := 0
	for _, c := range s {
		i := int(c - 'a')
		ans, memo[i] = (ans<<1-memo[i]+1+mod)%mod, ans+1 // memo[i]=ans+1：前缀 + c，再加上1个 c
	}
	return ans
}
