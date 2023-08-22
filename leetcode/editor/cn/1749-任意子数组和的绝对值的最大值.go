//给你一个整数数组 nums 。一个子数组 [numsl, numsl+1, ..., numsr-1, numsr] 的 和的绝对值 为 abs(
//numsl + numsl+1 + ... + numsr-1 + numsr) 。
//
// 请你找出 nums 中 和的绝对值 最大的任意子数组（可能为空），并返回该 最大值 。
//
// abs(x) 定义如下：
//
//
// 如果 x 是负整数，那么 abs(x) = -x 。
// 如果 x 是非负整数，那么 abs(x) = x 。
//
//
//
//
// 示例 1：
//
//
//输入：nums = [1,-3,2,3,-4]
//输出：5
//解释：子数组 [2,3] 和的绝对值最大，为 abs(2+3) = abs(5) = 5 。
//
//
// 示例 2：
//
//
//输入：nums = [2,-5,1,-4,3,-2]
//输出：8
//解释：子数组 [-5,1,-4] 和的绝对值最大，为 abs(-5+1-4) = abs(-8) = 8 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
// Related Topics 数组 动态规划 👍 32 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, -3, 2, 3, -4}
	sum := maxAbsoluteSum(nums)
	fmt.Println(sum)
}

/*
思路：dp
	1.假设所求为最大子序和（不考虑绝对值），那么 dp 方程应该是
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		当 dp[i-1] > 0：dp[i] = dp[i-1]+nums[i]
		否则 dp[i] = nums[i]
	2.当考虑绝对值时，我们就要额外考虑负数的情况，也就是考虑当前最小子序和
		最后的 dp 方程为：
		dpMin[i] = min(nums[i], dpMin[i-1]+nums[i])
		dpMax[i] = max(nums[i], dpMax[i-1]+nums[i])
		MAX = max(MAX, -dpMin, dpMax)
*/
// leetcode submit region begin(Prohibit modification and deletion)
func maxAbsoluteSum(nums []int) int {
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
	dpMin, dpMax, maxV := math.MaxInt32, math.MinInt32, math.MinInt32
	for _, v := range nums {
		dpMin, dpMax = minVal(v, dpMin+v), maxVal(v, dpMax+v) // 更新最小&大子序和
		maxV = maxVal(maxVal(-dpMin, dpMax), maxV)            // 更新最大绝对子序和
	}
	return maxV
}

//leetcode submit region end(Prohibit modification and deletion)
