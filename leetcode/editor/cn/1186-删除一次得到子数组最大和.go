package main

import "fmt"

func main() {
	arr := []int{1, -2, 0, 3}
	arr = []int{-1, -1, -2, -1}
	arr = []int{1, -2, -2, 3}
	arr = []int{1, -2, -2, 3, 6, -4, 2, 5}
	sum := maximumSum(arr)
	fmt.Println(sum)
}

/*
思路：dp
	1.对于任意 arr[i]，当前最大和 = max(不删除元素的前缀和, 删除一个元数的前缀和)
	2.分情况讨论状态转移方程
		2.1.当前最大和 = 不删除元素的前缀和：
			dp[i] = arr[i]，dp[i-1]<0
			dp[i] = dp[i-1] + arr[i]，dp[i-1]>=0
		2.2.当前最大和 = 删除一个元数的前缀和：
			dp[i] = dp[i-1]，删除 arr[i]
			dp[i] = dp[i-1] + arr[i]，含删除某个元素
*/
// leetcode submit region begin(Prohibit modification and deletion)
func maximumSum(arr []int) int {
	n := len(arr)
	ans, sum, dp := arr[0], arr[0], 0
	for i := 1; i < n; i++ {
		dp = max(sum, dp+arr[i])
		if sum < 0 {
			sum = arr[i]
		} else {
			sum += arr[i]
		}
		dp = max(dp, sum)
		ans = max(ans, dp)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func maximumSum_(arr []int) int {
	// dp
	// 初始化，防止 arr[0]<0
	ans, sum, dp, n := arr[0], arr[0], 0, len(arr)
	for i := 1; i < n; i++ {
		dp = max(sum, dp+arr[i]) // 初始化当前最大和
		if sum < 0 {             // 当前最大前缀和
			sum = arr[i]
		} else {
			sum += arr[i]
		}
		dp = max(dp, sum)  // 当前最大和
		ans = max(ans, dp) // 当前最大最大和
	}
	return ans
}
