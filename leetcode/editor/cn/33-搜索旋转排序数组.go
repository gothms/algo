//整数数组 nums 按升序排列，数组中的值 互不相同 。
//
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[
//k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2
//,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
//
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
//
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
//
//
// 示例 2：
//
//
//输入：nums = [4,5,6,7,0,1,2], target = 3
//输出：-1
//
// 示例 3：
//
//
//输入：nums = [1], target = 0
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5000
// -10⁴ <= nums[i] <= 10⁴
// nums 中的每个值都 独一无二
// 题目数据保证 nums 在预先未知的某个下标上进行了旋转
// -10⁴ <= target <= 10⁴
//
//
// Related Topics 数组 二分查找 👍 2691 👎 0

package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	//nums = []int{0, 1, 2}
	target := 0
	nums = []int{3, 1}
	target = 1
	i := search(nums, target)
	fmt.Println(i)
}

/*
思路：二分
	1.对于旋转数组 [4, 5, 6, 7, 0, 1, 2]，讨论它的二分查找指针的缩进方向
		设中间索引为 m = (l + r) >> 1，当 nums[m] == target，表示找到元素
	2.用数组的头元素 nums[0] 来确定指针 m 所在的位置
		nums[m] >= nums[0]：表示 m 处于旋转数组的左半部分
		nums[m] < nums[0]：表示 m 处于旋转数组的右半部分
	3.示例：指针的缩进方向，假设 target = 5
		第一次二分：nums[m] = 7
		如果 5 >= nums[0] 且 5 < nums[m]，往左缩进，r=m-1
		否则 l=m+1
*/
// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	l, r, m := 0, len(nums)-1, 0
	for l <= r {
		m = (l + r) >> 1
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			if nums[0] > nums[m] && nums[0] <= target {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[0] > target && nums[0] <= nums[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1

	//l, r := 0, len(nums)-1
	//for l <= r {
	//	m := (l + r) >> 1
	//	switch {
	//	case nums[m] == target:
	//		return m
	//	case nums[m] >= nums[0]:
	//		if target < nums[m] && target >= nums[0] { // m在前半部分，且 target 往头移
	//			r = m - 1
	//		} else {
	//			l = m + 1
	//		}
	//	default:
	//		if target > nums[m] && target < nums[0] { // m在后半部分，且 target 往尾移
	//			l = m + 1
	//		} else {
	//			r = m - 1
	//		}
	//	}
	//}
	//return -1
}

//leetcode submit region end(Prohibit modification and deletion)
