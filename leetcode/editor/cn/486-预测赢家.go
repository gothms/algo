package main

import "fmt"

func main() {
	nums := []int{1, 5, 2} // false
	//nums = []int{1, 5, 233, 7} // true
	winner := predictTheWinner(nums)
	fmt.Println(winner)
}

// leetcode submit region begin(Prohibit modification and deletion)
func predictTheWinner(nums []int) bool {
	// 方法二：dp[i][j] 表示选择 i/j 后，两个玩家的最大差值
	//n := len(nums)
	//dp := make([][]int, n)
	//for i := n - 1; i >= 0; i-- {
	//	dp[i] = make([]int, n)
	//	dp[i][i] = nums[i]
	//	for j := i + 1; j < n; j++ {
	//		dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
	//	}
	//}
	//return dp[0][n-1] >= 0

	// 方法一：参见 877
	n := len(nums)
	dp := make([][][2]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = make([][2]int, n)
		dp[i][i][0] = nums[i]
		for j := i + 1; j < n; j++ {
			l, r := nums[i]+dp[i+1][j][1], nums[j]+dp[i][j-1][1]
			if l > r {
				dp[i][j][0], dp[i][j][1] = l, dp[i+1][j][0]
			} else {
				dp[i][j][0], dp[i][j][1] = r, dp[i][j-1][0]
			}
		}
	}
	return dp[0][n-1][0] >= dp[0][n-1][1]
}

//leetcode submit region end(Prohibit modification and deletion)
