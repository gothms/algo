//峰值元素是指其值严格大于左右相邻值的元素。
//
// 给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
//
// 你可以假设 nums[-1] = nums[n] = -∞ 。
//
// 你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3,1]
//输出：2
//解释：3 是峰值元素，你的函数应该返回其索引 2。
//
// 示例 2：
//
//
//输入：nums = [1,2,1,3,5,6,4]
//输出：1 或 5
//解释：你的函数可以返回索引 1，其峰值元素为 2；
//     或者返回索引 5， 其峰值元素为 6。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 1000
// -2³¹ <= nums[i] <= 2³¹ - 1
// 对于所有有效的 i 都有 nums[i] != nums[i + 1]
//
//
// Related Topics 数组 二分查找 👍 1225 👎 0

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	//nums = []int{1, 2, 1, 3, 5, 6, 4}
	//nums = []int{1, 2}
	//nums = []int{1}
	element := findPeakElement(nums)
	fmt.Println(element)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findPeakElement(nums []int) int {
	// 二分查找
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) >> 1        // 保证了 m+1 不会越界
		if nums[m] < nums[m+1] { // 与 nums[0] 或 nums[r] 比较的方案都不行
			l = m + 1
		} else {
			r = m
		}
	}
	return l

	// 迭代
	//last, n := nums[0], len(nums)
	//for i := 1; i < n; i++ {
	//	if nums[i] < last  {
	//		return i-1
	//	}
	//	last = nums[i]
	//}
	//return n-1
}

//leetcode submit region end(Prohibit modification and deletion)
