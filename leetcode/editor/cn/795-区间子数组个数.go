package main

import "fmt"

func main() {
	nums := []int{2, 1, 4, 3}
	left, right := 2, 3 // 3
	nums = []int{2, 9, 2, 5, 6}
	left, right = 2, 8 // 7
	nums = []int{73, 55, 36, 5, 55, 14, 9, 7, 72, 52}
	left, right = 32, 69 // 22
	nums = []int{876, 880, 482, 260, 132, 421, 732,
		703,
		795, 420, 871, 445, 400, 291, 358, 589, 617, 202, 755, 810, 227, 813, 549, 791, 418, 528, 835, 401, 526, 584, 873,
		662, 13, 314,
		988, 101, 299, 816, 833, 224, 160, 852, 179, 769,
		646, 558, 661,
		808, 651, 982, 878, 918, 406, 551, 467, 87, 139, 387, 16, 531, 307, 389, 939, 551, 613, 36, 528, 460, 404, 314, 66, 111, 458, 531, 944, 461, 951, 419, 82, 896,
		467, 353, 704,
		905,
		705,
		760, 61, 422, 395, 298, 127, 516, 153, 299, 801,
		341, 668, 598, 98, 241}
	left, right = 658, 719 // 19
	boundedMax := numSubarrayBoundedMax(nums, left, right)
	fmt.Println(boundedMax)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numSubarrayBoundedMax(nums []int, left int, right int) int {
	// 个人：hash
	// 1.计算公式：v 符合 [left,right] 且左右各有 l r 个小于等于 right 的元素，则满足的子数组有 (l+1)*(r+1)
	// 2.去重：当 l r 中包含其他符合 [left,right] 的值（如右侧最近的是 rv），会重复计算
	// 则考虑左侧 l 中所有的元素，只考虑右侧 (v,rv) 区间的元素
	//n := len(nums)
	//memoL, memoR := make(map[int]int), make(map[int]int)
	//for i := 0; i < n; i++ {
	//	if nums[i] > right {
	//		continue
	//	}
	//	j := i
	//	for ; i < n && nums[i] <= right; i++ { // 左侧所有 <= right 的元素个数
	//		if nums[i] >= left {
	//			memoL[i] = i - j
	//		}
	//	}
	//}
	//for i := n - 1; i >= 0; i-- {
	//	if nums[i] > right {
	//		continue
	//	}
	//	j := i
	//	if nums[i] >= left { // 修正
	//		j--
	//	}
	//	for i--; i >= 0 && nums[i] <= right; i-- { // 右侧仅 < left 的元素个数
	//		if nums[i] >= left {
	//			memoR[i] = j - i
	//			j = i - 1
	//		}
	//	}
	//}
	//var ans int
	//for k, v := range memoL { // 统计：(l+1)*(r+1)
	//	ans += (v + 1) * (memoR[k] + 1)
	//}
	//return ans

	// 双指针
	ans, n := 0, len(nums)
	for i := 0; i < n; i++ {
		if nums[i] > right { // fast fail
			continue
		}
		maxV, start, last := 0, i, i
		for ; i < n && nums[i] <= right; i++ { // 寻找子数组的条件
			maxV = max(maxV, nums[i])
			if maxV < left { // 全是“小”数
				continue
			}
			if nums[i] >= left { // 子数组满足的必要元素
				ans += i - start + 1 // 以 i 为核心，构建子数组
				last = i             // 用于 nums[i] < left 时，以 nums[i] 为核心构建子数组
			} else {
				ans += last - start + 1 // 必须要有 >= left 的元素
			}
		}
	}
	return ans

	// lc
	//ans, i0, i1 := 0, -1, -1
	//for i, x := range nums {
	//	if x > right {
	//		i0 = i
	//	}
	//	if x >= left {
	//		i1 = i
	//	}
	//	ans += i1 - i0
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
