package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findMaxForm(strs []string, m int, n int) int {
	// dp
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		// bits.OnesCount / strings.Count(s, "0")
		var oz [2]int
		for _, c := range s {
			oz[c-'0']++
		}
		for i := m; i >= oz[0]; i-- {
			for j := n; j >= oz[1]; j-- {
				dp[i][j] = max(dp[i][j], dp[i-oz[0]][j-oz[1]]+1)
			}
		}
	}
	return dp[m][n]
}

//leetcode submit region end(Prohibit modification and deletion)
