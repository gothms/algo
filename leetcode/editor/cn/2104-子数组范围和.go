//给你一个整数数组 nums 。nums 中，子数组的 范围 是子数组中最大元素和最小元素的差值。
//
// 返回 nums 中 所有 子数组范围的 和 。
//
// 子数组是数组中一个连续 非空 的元素序列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：4
//解释：nums 的 6 个子数组如下所示：
//[1]，范围 = 最大 - 最小 = 1 - 1 = 0
//[2]，范围 = 2 - 2 = 0
//[3]，范围 = 3 - 3 = 0
//[1,2]，范围 = 2 - 1 = 1
//[2,3]，范围 = 3 - 2 = 1
//[1,2,3]，范围 = 3 - 1 = 2
//所有范围的和是 0 + 0 + 0 + 1 + 1 + 2 = 4
//
// 示例 2：
//
//
//输入：nums = [1,3,3]
//输出：4
//解释：nums 的 6 个子数组如下所示：
//[1]，范围 = 最大 - 最小 = 1 - 1 = 0
//[3]，范围 = 3 - 3 = 0
//[3]，范围 = 3 - 3 = 0
//[1,3]，范围 = 3 - 1 = 2
//[3,3]，范围 = 3 - 3 = 0
//[1,3,3]，范围 = 3 - 1 = 2
//所有范围的和是 0 + 0 + 0 + 2 + 0 + 2 = 4
//
//
// 示例 3：
//
//
//输入：nums = [4,-2,-3,4,1]
//输出：59
//解释：nums 中所有子数组范围的和是 59
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 1000
// -10⁹ <= nums[i] <= 10⁹
//
//
//
//
// 进阶：你可以设计一种时间复杂度为 O(n) 的解决方案吗？
//
// Related Topics 栈 数组 单调栈 👍 278 👎 0

package main

import "fmt"

func main() {
	nums := []int{4, -2, -3, 4, 1}
	nums = []int{3, 1, 3}
	ranges := subArrayRanges(nums)
	fmt.Println(ranges)
}

/*
补充
    如果不考虑最大值 - 最小值，而是考虑 nums[j] - nums[i]，其中 j > i，n = nums.len - 1
    那么和 = (n - 1) * nums[-1] + (n-3) * nums[-2] + ... - (n-3) * nums[1] - (n-1) * nums[0]
*/
//leetcode submit region begin(Prohibit modification and deletion)
func subArrayRanges(nums []int) int64 {

}

//leetcode submit region end(Prohibit modification and deletion)

//func subArrayRanges(nums []int) int64 {
//	n := len(nums)
//	var (
//		minLeft  = make([]int, n)
//		maxLeft  = make([]int, n)
//		minRight = make([]int, n)
//		maxRight = make([]int, n)
//		minSt    = []int{-1}
//		maxSt    = []int{-1}
//	)
//	for i, v := range nums {
//		last := len(minSt) - 1
//		// 要包含所有的子数组，所以 == 的情况也必须有“一边”要考虑包含 v 和 nums[minSt[last]] 的子数组
//		// 约定：左边都管 =，右边都不管 =。即“左边”考虑元素相等的子数组，右边则避免
//		for last > 0 && v <= nums[minSt[last]] { // 单调非递减
//			last--
//		}
//		minLeft[i] = minSt[last] // 记录 v 作为子数组最小值的左范围，即在索引区间 (minSt[last],i] 内，v 都最小
//		minSt = append(minSt[:last+1], i)
//
//		last = len(maxSt) - 1
//		for last > 0 && v >= nums[maxSt[last]] { // 单调非递增
//			last--
//		}
//		maxLeft[i] = maxSt[last]
//		maxSt = append(maxSt[:last+1], i)
//	}
//	minSt, maxSt = minSt[:1], maxSt[:1]
//	minSt[0], maxSt[0] = n, n
//	for i := n - 1; i >= 0; i-- {
//		last := len(minSt) - 1
//		for last > 0 && nums[i] < nums[minSt[last]] {
//			last--
//		}
//		minRight[i] = minSt[last]
//		minSt = append(minSt[:last+1], i)
//
//		last = len(maxSt) - 1
//		for last > 0 && nums[i] > nums[maxSt[last]] {
//			last--
//		}
//		maxRight[i] = maxSt[last]
//		maxSt = append(maxSt[:last+1], i)
//	}
//	var ans int
//	for i, v := range nums {
//		//ans -= v * (i - minLeft[i]) * (minRight[i] - i)
//		//ans += v * (i - maxLeft[i]) * (maxRight[i] - i)
//		ans += v * ((i-maxLeft[i])*(maxRight[i]-i) - (i-minLeft[i])*(minRight[i]-i))
//		// v 作为最大值的 左范围 * 右范围 - v 作为最小值的 左范围 * 右范围
//	}
//	return int64(ans)
//
//	// 暴力法：O(n^2)
//}
