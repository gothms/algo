package main

import "fmt"

func main() {
	steps, arrLen := 2, 4
	steps, arrLen = 27, 7 // 127784505
	i := numWays(steps, arrLen)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numWays(steps int, arrLen int) int {
	// dp
	const mod = 1_000_000_007
	dp, temp := make([]int, arrLen+2), make([]int, arrLen+2)
	dp[1] = 1
	for i := steps; i > 0; i-- {
		for j := 0; j <= min(i, arrLen-1); j++ {
			temp[j+1] = (dp[j] + dp[j+1] + dp[j+2]) % mod
		}
		dp, temp = temp, dp
	}
	return dp[1]
}

//leetcode submit region end(Prohibit modification and deletion)
