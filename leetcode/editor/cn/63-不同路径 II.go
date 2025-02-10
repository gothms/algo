package main

import "fmt"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	_, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, n+1)
	dp[1] = 1
	for _, r := range obstacleGrid {
		for i, v := range r {
			if v == 0 {
				dp[i+1] += dp[i]
			} else {
				dp[i+1] = 0
			}
		}
	}
	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
