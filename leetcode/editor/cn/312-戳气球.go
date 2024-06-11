package main

import "fmt"

func main() {
	nums := []int{3, 1, 5, 8}
	coins := maxCoins(nums)
	fmt.Println(coins)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxCoins(nums []int) int {
	// 记忆化搜索
}

//leetcode submit region end(Prohibit modification and deletion)

func maxCoins_(nums []int) int {
	temp := append(append([]int{1}, nums...), 1)
	n := len(temp)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 左闭右开
	for i := n - 2; i > 0; i-- { // 左边界
		for j := i + 1; j < n; j++ { // 右边界
			for k := i; k < j; k++ { // 中间值
				dp[i][j] = max(dp[i][j], temp[k]*temp[i-1]*temp[j]+dp[i][k]+dp[k+1][j])
			}
		}
	}
	return dp[1][n-1]

	// 记忆化搜索
}
