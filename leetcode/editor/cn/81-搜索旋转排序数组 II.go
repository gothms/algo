//已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。
//
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，使数组变为 [nums[k],
//nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,
//2,4,4,4,5,6,6,7] 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。
//
// 给你 旋转后 的数组 nums 和一个整数 target ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 nums 中存在这个目标值
//target ，则返回 true ，否则返回 false 。
//
// 你必须尽可能减少整个操作步骤。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,5,6,0,0,1,2], target = 0
//输出：true
//
//
// 示例 2：
//
//
//输入：nums = [2,5,6,0,0,1,2], target = 3
//输出：false
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5000
// -10⁴ <= nums[i] <= 10⁴
// 题目数据保证 nums 在预先未知的某个下标上进行了旋转
// -10⁴ <= target <= 10⁴
//
//
//
//
// 进阶：
//
//
// 这是 搜索旋转排序数组 的延伸题目，本题中的 nums 可能包含重复元素。
// 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
//
//
//
//
// Related Topics 数组 二分查找 👍 726 👎 0

package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}
	target := 2
	b := search(nums, target)
	fmt.Println(b)
}

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) bool {
	// 二分
	//l, r := 0, len(nums)-1
	//for r > 0 && nums[0] == nums[r] {
	//	r--
	//}
	//for l <= r {
	//	m := (l + r) >> 1
	//	if target == nums[m] {
	//		return true
	//	}
	//	if nums[m] >= nums[0] {
	//		if target < nums[m] && target >= nums[0] {
	//			r = m - 1
	//		} else {
	//			l = m + 1
	//		}
	//	} else {
	//		if target > nums[m] && target < nums[0] {
	//			l = m + 1
	//		} else {
	//			r = m - 1
	//		}
	//	}
	//}
	//return false

	// lc
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) >> 1
		if target == nums[m] {
			return true
		}
		if nums[m] == nums[l] && nums[m] == nums[r] { // 首尾相同，m 可能在前半部分，也可能在后
			l++
			r--
		} else if nums[m] >= nums[l] { // 逻辑和 lc-33 一样
			if target < nums[m] && target >= nums[l] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target > nums[m] && target < nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
