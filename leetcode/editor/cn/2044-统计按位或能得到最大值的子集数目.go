package main

import "fmt"

func main() {
	nums := []int{3, 1}
	//nums = []int{2, 2, 2}
	//nums = []int{3, 2, 1, 5}
	orSubsets := countMaxOrSubsets(nums)
	fmt.Println(orSubsets)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countMaxOrSubsets(nums []int) int {
	// dfs

	// 位运算

	// 个人
	//mx := 0
	//for _, v := range nums {
	//	mx |= v
	//}
	//dp := make([]int, mx+1)
	//dp[0] = 1
	//for _, v := range nums {
	//	temp := make([]int, mx+1)
	//	for j, w := range dp {
	//		if w > 0 {
	//			temp[v|j] += w
	//		}
	//	}
	//	for i, v := range temp {
	//		dp[i] += v
	//	}
	//}
	//return dp[mx]
}

//leetcode submit region end(Prohibit modification and deletion)
