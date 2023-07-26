//给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。
//
// 环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ， nums[i] 的前一个
//元素是 nums[(i - 1 + n) % n] 。
//
// 子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j] ，不
//存在 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,-2,3,-2]
//输出：3
//解释：从子数组 [3] 得到最大和 3
//
//
// 示例 2：
//
//
//输入：nums = [5,-3,5]
//输出：10
//解释：从子数组 [5,5] 得到最大和 5 + 5 = 10
//
//
// 示例 3：
//
//
//输入：nums = [3,-2,2,-3]
//输出：3
//解释：从子数组 [3] 和 [3,-2,2] 都可以得到最大和 3
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 3 * 10⁴
// -3 * 10⁴ <= nums[i] <= 3 * 10⁴
//
//
// Related Topics 队列 数组 分治 动态规划 单调队列 👍 511 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, -2, 3, -2}
	nums = []int{-3, -2, -3}
	circular := maxSubarraySumCircular(nums)
	fmt.Println(circular)
}

/*
思路：最小子序和
	1.根据题意，环形数组的最大子数组和
		要么是 最大子序和
		要么是数组的总和 sum - 最小子序和（最小子序和 != sum，否则取最大子序和）
	2.最大子序和
		dp[i-1] > 0：dp[i] = dp[i-1]+nums[i]
		dp[i-1] <= 0：dp[i] = nums[i]
	3.最小子序和
		思路同上
思路：dp
	1.分两种情况：最大子序和 and 最大前缀和+后缀和
	2.最大子序和
		dp[i-1] > 0：dp[i] = dp[i-1]+nums[i]
		dp[i-1] <= 0：dp[i] = nums[i]
	3.最大前缀和+后缀和
		最大前缀和：将每个位置所对应的最大前缀和记录到数组 maxL 中
		后缀和：倒序遍历 nums，累加元素之和 sumR
		最大前缀和+后缀和：sum[i] = sumR[i] + maxL[i-1]
*/
// leetcode submit region begin(Prohibit modification and deletion)
type preIdx struct {
	v   int
	idx int
}

func maxSubarraySumCircular(nums []int) int {
	// 队列+滑动窗体
	//maxVal := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//max, preSum, n, m := nums[0], nums[0], len(nums), len(nums)<<1
	//queue := []preIdx{{nums[0], 0}}
	//for i := 1; i < m; i++ {
	//	if len(queue) > 0 && queue[0].idx < i-n {
	//		queue = queue[1:]
	//	}
	//	preSum += nums[i%n]
	//	max = maxVal(max, preSum-queue[0].v)
	//	l := len(queue) - 1
	//	for l >= 0 && queue[l].v >= preSum {
	//		l--
	//	}
	//	queue = append(queue[:l+1], preIdx{preSum, i})
	//}
	//return max

	// dp：最小子序和
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	minVal := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	sum, max, min, dpMax, dpMin, n :=
		0, math.MinInt32, math.MaxInt32, 0, 0, len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
		if dpMax <= 0 { // 最大子序和
			dpMax = nums[i]
		} else {
			dpMax += nums[i]
		}
		if dpMin >= 0 { // 最小子序和
			dpMin = nums[i]
		} else {
			dpMin += nums[i]
		}
		max, min = maxVal(max, dpMax), minVal(min, dpMin)
	}
	if max < 0 { // 防止 sum-min==0
		return max
	}
	return maxVal(max, sum-min)

	// dp
	//maxVal := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//max, dp, n := nums[0], nums[0], len(nums)
	//maxL := make([]int, n)
	//maxL[0] = nums[0]
	//for i, sum := 1, nums[0]; i < n; i++ {
	//	if dp > 0 {
	//		dp += +nums[i] // 连续
	//	} else {
	//		dp = nums[i] // 噶断
	//	}
	//	max = maxVal(max, dp)            // 最大子序和
	//	sum += nums[i]                   // 前缀和
	//	maxL[i] = maxVal(maxL[i-1], sum) // 最大前缀和
	//}
	//for i, sum := n-1, 0; i > 0; i-- {
	//	sum += nums[i]                   // 后缀和
	//	max = maxVal(max, sum+maxL[i-1]) // 最大前缀和+后缀和
	//}
	//return max
}

//leetcode submit region end(Prohibit modification and deletion)
