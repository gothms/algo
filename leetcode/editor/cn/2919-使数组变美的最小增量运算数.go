package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n, k := 10, 15
	N := 20
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(N)
	}
	fmt.Println(nums, k)
	nums = []int{2, 3, 0, 0, 2} // 3
	k = 4
	//nums = []int{4, 0, 10, 2, 10, 6} // 0
	//k = 8
	//nums = []int{4, 0, 22, 41, 29, 28} // 8
	//k = 30
	//nums = []int{17, 1, 14, 8, 4, 3, 7, 4, 13, 7} // 15
	//k = 15
	operations := minIncrementOperations(nums, k)
	fmt.Println(operations)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minIncrementOperations(nums []int, k int) int64 {
	// dp
	var f0, f1, f2 int
	for _, x := range nums {
		inc := f0 + max(k-x, 0)
		f0 = min(inc, f1) // 连续不增
		f1 = min(inc, f2) // 不增
		f2 = inc          // 必增
		//fmt.Println(x, inc, f0, f1, f2)
	}
	return int64(f0)

	// 记忆化搜索

	// dp：个人
	//n := len(nums)
	//dp, same := make([]int, n), make([]int, n)
	//dp[0], dp[1], dp[2] = max(0, k-nums[0]), max(0, k-nums[1]), max(0, k-nums[2])
	//same[0], same[1], same[2] = 0, dp[0], min(dp[1], dp[0])
	//for i := 3; i < n; i++ {
	//	dp[i] = min(dp[i-1], dp[i-2], dp[i-3]) + max(0, k-nums[i]) // 增大为 k
	//	same[i] = min(dp[i-1], dp[i-2])                            // 不增大
	//}
	////fmt.Println(dp)
	////fmt.Println(same)
	//return int64(min(dp[n-1], same[n-1]))
}

//leetcode submit region end(Prohibit modification and deletion)
