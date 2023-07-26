//一个下标从 0 开始的数组的 交替和 定义为 偶数 下标处元素之 和 减去 奇数 下标处元素之 和 。
//
//
// 比方说，数组 [4,2,5,3] 的交替和为 (4 + 5) - (2 + 3) = 4 。
//
//
// 给你一个数组 nums ，请你返回 nums 中任意子序列的 最大交替和 （子序列的下标 重新 从 0 开始编号）。
//
//
//
//
// 一个数组的 子序列 是从原数组中删除一些元素后（也可能一个也不删除）剩余元素不改变顺序组成的数组。比方说，[2,7,4] 是 [4,2,3,7,2,1,4
//] 的一个子序列（加粗元素），但是 [2,4,2] 不是。
//
//
//
// 示例 1：
//
// 输入：nums = [4,2,5,3]
//输出：7
//解释：最优子序列为 [4,2,5] ，交替和为 (4 + 5) - 2 = 7 。
//
//
// 示例 2：
//
// 输入：nums = [5,6,7,8]
//输出：8
//解释：最优子序列为 [8] ，交替和为 8 。
//
//
// 示例 3：
//
// 输入：nums = [6,2,1,2,4,5]
//输出：10
//解释：最优子序列为 [6,1,5] ，交替和为 (6 + 5) - 1 = 10 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁵
//
//
// Related Topics 数组 动态规划 👍 54 👎 0

package main

import "fmt"

func main() {
	nums := []int{5, 6, 7, 8}
	nums = []int{6, 2, 1, 2, 4, 5}
	sum := maxAlternatingSum(nums)
	fmt.Println(sum)
}

/*
lc
	偶数：dp[i][0]
		nums[i] 被选择：dp[i][0] = dp[i-1][1] + nums[i]
		否则：dp[i][0] = dp[i-1][0]
	奇数：dp[i][1]
		nums[i] 被选择：dp[i][1] = dp[i-1][0] - nums[i]
		否则：dp[i][1] = dp[i-1][1]
*/
/*
思路：dp
	1.最优子序列的元素选择上，对于任意 nums[i] 存在选与不选两种情况
		那么什么时候选择，什么时候不选择呢？
	2.当前子序列的元素个数为奇数/偶数时：
		偶数：由于 1 <= nums[i] <= 10^5，所以只有在接连选择两个数 x y，且 x<y 时
			对当前子序列的交替和的贡献值为正
		奇数：选择任意数 x，对当前子序列的交替和的贡献值都为正
	3.动态规划
		3.1.经过上面分析，当偶数时，仍然需要维护 偶数-nums[i] 的情况，以备后续 x<y 的情况
		3.2.需要维护子序列的两种状态，分别是偶数/奇数个元素
			偶数个元素的交替和为 max0：max0 = maxVal(max0, max1-nums[i])
			奇数个元素的交替和为 max1：max1 = maxVal(max1, max0+nums[i])
*/
// leetcode submit region begin(Prohibit modification and deletion)
func maxAlternatingSum(nums []int) int64 {
	// dp
	maxVal := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}
	var (
		max0 int64 // 子序列元素为偶数个
		max1 int64 // 子序列元素为奇数个
		n    = len(nums)
	)
	for i := 0; i < n; i++ {
		max0, max1 = maxVal(max0, max1-int64(nums[i])), maxVal(max1, max0+int64(nums[i]))
	}
	return maxVal(max0, max1)
}

//leetcode submit region end(Prohibit modification and deletion)
