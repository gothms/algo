package main

import (
	"fmt"
)

func main() {
	nums := []int{12, 15, 1, 15, 18}   // 34
	nums = []int{10, 3, 8, 10, 18, 12} // 36
	cost := minCost(nums)
	fmt.Println(cost)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minCost(nums []int) int {
	// 记忆化搜索
	// 状态定义：f(i,j)，独立元素为 j，区间 [i,n) 的最小成本，其中 j<i
	// 状态转移方程：f(i,j) = min(
	// f(i+2,j) + max(nums[i],nums[i+1),
	// f(i+2,i) + max(nums[j],nums[i+1),
	// f(i+2,i+1) + max(nums[j],nums[i),
	// )
	// 边界条件：j==n-1，j==n-2
	n := len(nums)
	memo := make([][]int, n-1)
	for i := range memo {
		memo[i] = make([]int, i) // j<i
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == n {
			return nums[j]
		}
		if i == n-1 {
			return max(nums[i], nums[j])
		}
		v := &memo[i][j]
		if *v == 0 {
			*v = min(
				dfs(i+2, j)+max(nums[i], nums[i+1]),
				dfs(i+2, i)+max(nums[j], nums[i+1]),
				dfs(i+2, i+1)+max(nums[j], nums[i]))
		}
		return *v
	}
	dfs(1, 0)
	return dfs(1, 0)
}

//leetcode submit region end(Prohibit modification and deletion)
