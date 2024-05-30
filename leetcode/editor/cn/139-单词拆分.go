package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func wordBreak(s string, wordDict []string) bool {
	// dp
	n := len(s)
	dp, memo := make([]int, 1, n), make(map[string]struct{}, len(wordDict))
	for _, w := range wordDict {
		memo[w] = struct{}{}
	}
	for i := 1; i <= n; i++ {
		for j := len(dp) - 1; j >= 0; j-- {
			if _, ok := memo[s[dp[j]:i]]; ok {
				dp = append(dp, i)
				break
			}
		}
	}
	return dp[len(dp)-1] == n

	//n := len(s)
	//memo := make(map[string]struct{}, len(wordDict))
	//for _, w := range wordDict {
	//	memo[w] = struct{}{}
	//}
	//vis := make([]bool, n)
	//var dfs func(int) bool
	//dfs = func(i int) bool {
	//	if i == n {
	//		return true
	//	}
	//	if vis[i] {
	//		return false
	//	}
	//	for j := i; j < n; j++ {
	//		if _, ok := memo[s[i:j+1]]; ok && dfs(j+1) {
	//			return true
	//		}
	//	}
	//	vis[i] = true
	//	return false
	//}
	//return dfs(0)
}

//leetcode submit region end(Prohibit modification and deletion)
