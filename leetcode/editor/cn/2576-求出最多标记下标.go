//给你一个下标从 0 开始的整数数组 nums 。
//
// 一开始，所有下标都没有被标记。你可以执行以下操作任意次：
//
//
// 选择两个 互不相同且未标记 的下标 i 和 j ，满足 2 * nums[i] <= nums[j] ，标记下标 i 和 j 。
//
//
// 请你执行上述操作任意次，返回 nums 中最多可以标记的下标数目。
//
//
//
// 示例 1：
//
//
//输入：nums = [3,5,2,4]
//输出：2
//解释：第一次操作中，选择 i = 2 和 j = 1 ，操作可以执行的原因是 2 * nums[2] <= nums[1] ，标记下标 2 和 1 。
//没有其他更多可执行的操作，所以答案为 2 。
//
//
// 示例 2：
//
//
//输入：nums = [9,2,5,4]
//输出：4
//解释：第一次操作中，选择 i = 3 和 j = 0 ，操作可以执行的原因是 2 * nums[3] <= nums[0] ，标记下标 3 和 0 。
//第二次操作中，选择 i = 1 和 j = 2 ，操作可以执行的原因是 2 * nums[1] <= nums[2] ，标记下标 1 和 2 。
//没有其他更多可执行的操作，所以答案为 4 。
//
//
// 示例 3：
//
//
//输入：nums = [7,6,8]
//输出：0
//解释：没有任何可以执行的操作，所以答案为 0 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁹
//
//
// Related Topics 贪心 数组 双指针 二分查找 排序 👍 46 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 5, 2, 4}
	nums = []int{6, 8, 7}
	nums = []int{4, 8, 7, 16, 2, 3}
	nums = []int{42, 83, 48, 10, 24, 55, 9, 100, 10, 17, 17, 99, 51, 32, 16, 98, 99, 31, 28, 68, 71, 14, 64, 29, 15, 40}

	indices := maxNumOfMarkedIndices(nums)
	fmt.Println(indices)
}

/*
反证法：
	1
		如果 2 * nums ≤ nums[j]，则称 nums[i] 和 nums[j] 匹配
		如果可以匹配 kkk 对，那么也可以匹配小于 kkk 对，去掉一些数对即可做到
		如果无法匹配 kkk 对，那么也无法匹配大于 kkk 对（反证法）
		因此答案有单调性，可以二分
	2
		检测能不能匹配 kkk 对
		要让哪些数匹配呢？
		结论：从小到大排序后，如果存在 kkk 对匹配，那么一定可以让最小的 kkk 个数和最大的 kkk 个数匹配

*/
// leetcode submit region begin(Prohibit modification and deletion)
func maxNumOfMarkedIndices(nums []int) int {
	// 双指针
	sort.Ints(nums)
	j, n := 0, len(nums)
	for i := (n + 1) >> 1; i < n; i++ {
		if nums[j]<<1 <= nums[i] {
			j++
		}
	}
	return j << 1

	// lc
	//sort.Ints(nums)
	//n := len(nums)
	//return 2 * sort.Search(n/2, func(k int) bool {
	//	k++
	//	for i, x := range nums[:k] {
	//		if x*2 > nums[n-k+i] {
	//			return true
	//		}
	//	}
	//	return false
	//})
}

//leetcode submit region end(Prohibit modification and deletion)
