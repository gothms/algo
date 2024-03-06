//给你两个长度相等下标从 0 开始的整数数组 nums1 和 nums2 。每一秒，对于所有下标 0 <= i < nums1.length ，nums1[
//i] 的值都增加 nums2[i] 。操作 完成后 ，你可以进行如下操作：
//
//
// 选择任一满足 0 <= i < nums1.length 的下标 i ，并使 nums1[i] = 0 。
//
//
// 同时给你一个整数 x 。
//
// 请你返回使 nums1 中所有元素之和 小于等于 x 所需要的 最少 时间，如果无法实现，那么返回 -1 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,2,3], nums2 = [1,2,3], x = 4
//输出：3
//解释：
//第 1 秒，我们对 i = 0 进行操作，得到 nums1 = [0,2+2,3+3] = [0,4,6] 。
//第 2 秒，我们对 i = 1 进行操作，得到 nums1 = [0+1,0,6+3] = [1,0,9] 。
//第 3 秒，我们对 i = 2 进行操作，得到 nums1 = [1+1,0+2,0] = [2,2,0] 。
//现在 nums1 的和为 4 。不存在更少次数的操作，所以我们返回 3 。
//
//
// 示例 2：
//
//
//输入：nums1 = [1,2,3], nums2 = [3,3,3], x = 4
//输出：-1
//解释：不管如何操作，nums1 的和总是会超过 x 。
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length <= 10³
// 1 <= nums1[i] <= 10³
// 0 <= nums2[i] <= 10³
// nums1.length == nums2.length
// 0 <= x <= 10⁶
//
//
// Related Topics 数组 动态规划 排序 👍 39 👎 0

package main

import (
	"fmt"
	"slices"
)

func main() {
	nums1 := []int{1, 2, 7}
	nums2 := []int{5, 2, 4}
	x := 9
	time := minimumTime(nums1, nums2, x)
	fmt.Println(time)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumTime(nums1 []int, nums2 []int, x int) int {
	// dp：详见下面 lc 注释
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	s1, s2, n := 0, 0, len(nums1)
	ids := make([]int, n)
	for i := range ids {
		s1 += nums1[i]
		s2 += nums2[i]
		ids[i] = i
	}
	slices.SortFunc(ids, func(a, b int) int {
		return nums2[a] - nums2[b]
	})
	dp := make([]int, n+1)
	for i := range ids {
		v1, v2 := nums1[ids[i]], nums2[ids[i]]
		for j := i + 1; j > 0; j-- {
			dp[j] = maxVal(dp[j], dp[j-1]+v1+j*v2) // 操作元素 j
		}
	}
	for i, v := range dp {
		if s1+s2*i-v <= x {
			return i
		}
	}
	return -1

	// lc
	//maxVal := func(a, b int) int {
	//	if b > a {
	//		return b
	//	}
	//	return a
	//}
	//n := len(nums1)
	//dp := make([][]int, n+1) // 对前 j 个元素进行 i 次操作，可以减少的最大总值，初始值为零
	//for i := 0; i <= n; i++ {
	//	dp[i] = make([]int, n+1) // 第 0 秒：nums1 的和
	//}
	//
	//nums := make([][2]int, n)
	//for i := 0; i < n; i++ {
	//	nums[i] = [2]int{nums2[i], nums1[i]}
	//}
	//sort.Slice(nums, func(i, j int) bool {
	//	return nums[i][0] < nums[j][0] // 优先操作增长速度慢的数
	//})
	//for j := 1; j <= n; j++ {
	//	b, a := nums[j-1][0], nums[j-1][1]
	//	for i := j; i > 0; i-- { // 对于第 j 个元素，我们可以选择对其进行操作或者不操作
	//		dp[j][i] = maxVal(dp[j-1][i], dp[j-1][i-1]+i*b+a)
	//	} // 状态转移方程
	//}
	//
	//s1 := 0
	//for _, v := range nums1 {
	//	s1 += v
	//}
	//s2 := 0
	//for _, v := range nums2 {
	//	s2 += v
	//}
	//for i := 0; i <= n; i++ {
	//	if s2*i+s1-dp[n][i] <= x {
	//		return i
	//	}
	//}
	//return -1
}

//leetcode submit region end(Prohibit modification and deletion)
