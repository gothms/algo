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
	// 单调栈
	n := len(nums)
	minLeft := make([]int, n)
	maxLeft := make([]int, n)
	var minStk, maxStk []int
	for i, num := range nums {
		for len(minStk) > 0 && nums[minStk[len(minStk)-1]] > num { // 递增
			minStk = minStk[:len(minStk)-1]
		}
		if len(minStk) > 0 {
			minLeft[i] = minStk[len(minStk)-1]
		} else {
			minLeft[i] = -1
		}
		minStk = append(minStk, i)

		// 如果 nums[maxStk[len(maxStk)-1]] == num, 那么根据定义，
		// nums[maxStk[len(maxStk)-1]] 逻辑上小于 num，因为 maxStk[len(maxStk)-1] < i
		for len(maxStk) > 0 && nums[maxStk[len(maxStk)-1]] <= num { // 递减
			maxStk = maxStk[:len(maxStk)-1]
		}
		if len(maxStk) > 0 {
			maxLeft[i] = maxStk[len(maxStk)-1]
		} else {
			maxLeft[i] = -1
		}
		maxStk = append(maxStk, i)
	}

	minRight := make([]int, n)
	maxRight := make([]int, n)
	minStk = minStk[:0]
	maxStk = maxStk[:0]
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		// 如果 nums[minStk[len(minStk)-1]] == num, 那么根据定义，
		// nums[minStk[len(minStk)-1]] 逻辑上大于 num，因为 minStk[len(minStk)-1] > i
		for len(minStk) > 0 && nums[minStk[len(minStk)-1]] >= num {
			minStk = minStk[:len(minStk)-1]
		}
		if len(minStk) > 0 {
			minRight[i] = minStk[len(minStk)-1]
		} else {
			minRight[i] = n
		}
		minStk = append(minStk, i)

		for len(maxStk) > 0 && nums[maxStk[len(maxStk)-1]] < num {
			maxStk = maxStk[:len(maxStk)-1]
		}
		if len(maxStk) > 0 {
			maxRight[i] = maxStk[len(maxStk)-1]
		} else {
			maxRight[i] = n
		}
		maxStk = append(maxStk, i)
	}
	/*
		nums := []int{4, -2, -3, 4, 1}
		[-1 -1 -1 2 2]
		[-1 0 1 -1 3]
		[1 2 5 4 5]
		[3 3 3 5 5]
	*/
	//fmt.Println(minLeft)
	//fmt.Println(maxLeft)
	//fmt.Println(minRight)
	//fmt.Println(maxRight)
	var sumMax, sumMin int64
	for i, num := range nums {
		// 右侧 & 右侧以 num 为最大值的子数组的数量 = maxRight[i]-i & i-maxLeft[i]，则总数量为两者的乘积
		sumMax += int64(maxRight[i]-i) * int64(i-maxLeft[i]) * int64(num) // 最大值之和
		sumMin += int64(minRight[i]-i) * int64(i-minLeft[i]) * int64(num) // 最小值之和
	}
	return sumMax - sumMin

	// 迭代
	//sar, n := int64(0), len(nums)-1
	//for i := 0; i < n; i++ {
	//	minV, maxV := nums[i], nums[i]
	//	for j := i + 1; j <= n; j++ {
	//		if nums[j] < minV {
	//			minV = nums[j]
	//		} else if nums[j] > maxV {
	//			maxV = nums[j]
	//		}
	//		sar += int64(maxV - minV)
	//	}
	//}
	//return sar
}

//leetcode submit region end(Prohibit modification and deletion)
