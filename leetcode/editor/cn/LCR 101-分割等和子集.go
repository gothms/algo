//给定一个非空的正整数数组 nums ，请判断能否将这些数字分成元素和相等的两部分。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,5,11,5]
//输出：true
//解释：nums 可以分割成 [1, 5, 5] 和 [11] 。
//
// 示例 2：
//
//
//输入：nums = [1,2,3,5]
//输出：false
//解释：nums 不可以分为和相等的两部分
//
//
//
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 100
//
//
//
//
//
// 注意：本题与主站 416 题相同： https://leetcode-cn.com/problems/partition-equal-subset-
//sum/
//
// Related Topics 数学 字符串 模拟 👍 86 👎 0

package main

import "fmt"

func main() {
	nums := []int{1, 2, 5}
	nums = []int{1, 5, 11, 5}
	b := canPartition(nums)
	fmt.Println(b)
}

/*
思路：dp
	1.不能均分成两部分的特殊情况，即返回 false
		nums 只有一个元素
		和为 sum，是奇数时
		nums 的最大元素 > sum/2 时
	2.什么情况下返回 true？
		2.1.如果存在均分 nums 的情况，那么一定存在某种组合的和 = sum/2
			暴力组合的情况有 2^n 种，n为nums的元素个数
		2.2.我们可以记录 [0, sum/2] 区间里，哪些组合情况是存在的
			通过追溯 sum/2 到 0 这个过程，就可以递推出 sum/2 是否存在
	3.尝试 dp
		定义 dp[i][j] 二维数组，记录 i 为 nums 的索引，j 为 [0, sum/2] 区间时
		是否存在元素组合的和为 j
	4.dp 方程
		当 nums[i]>j，肯定不能选择 nums[i]，此时 dp[i][j] = dp[i-1][j]
		当 nums[i]<=j：
			如果不选择 nums[i]，dp[i][j] = dp[i-1][j]
			如果选择 nums[i]，dp[i][j] = dp[i-1][j-nums[i]]
*/
// leetcode submit region begin(Prohibit modification and deletion)
func canPartition(nums []int) bool {
	// dp
	sum, max, n := 0, 0, len(nums)
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	m := sum >> 1 // 平分数
	if n == 1 || sum&1 > 0 || max > m {
		return false
	} // 1 个元素 / 和为奇数 / 最大数 > 平分数
	dp := make([]bool, m+1) // 滚动dp
	dp[0] = true
	for _, v := range nums {
		for i := m; i >= v; i-- { // 倒序遍历，防止脏数据
			dp[i] = dp[i] || dp[i-v]
		}
		fmt.Println(dp)
	}
	return dp[m]
}

//leetcode submit region end(Prohibit modification and deletion)
