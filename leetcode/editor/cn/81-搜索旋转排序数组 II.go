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
