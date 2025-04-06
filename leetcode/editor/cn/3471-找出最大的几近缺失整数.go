package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{3, 9, 7, 2, 1, 7} // 3
	k := 4
	nums = []int{2, 0, 11, 6, 10, 1, 0, 10} // 2
	k = 2
	nums = []int{8, 6, 2, 8, 6} // -1
	k = 4
	nums = []int{3, 1, 7, 10, 0} // 10
	k = 1
	integer := largestInteger(nums, k)
	fmt.Println(integer)
}

// leetcode submit region begin(Prohibit modification and deletion)
func largestInteger(nums []int, k int) int {
	// 个人
	n := len(nums)
	if n == k { // 一个子数组
		return slices.Max(nums)
	}
	if k == 1 { // n 个子数组
		mx := -1
		memo := make(map[int]int)
		for _, v := range nums {
			memo[v]++
		}
		for v, c := range memo {
			if c == 1 {
				mx = max(mx, v)
			}
		}
		return mx
	}
	l, r := nums[0], nums[n-1] // 几近缺失的整数只可能是 nums[0] / nums[n-1]
	if l == r {                // 相等，则没有
		return -1
	}
	check := func(arr []int, v int) int {
		if slices.Contains(arr, v) {
			return -1
		}
		return v
	}
	return max(check(nums[1:], l), check(nums[:n-1], r))

	//for _, v := range nums[1 : n-1] { // 除去首尾，l / r 重复出现
	//	if v == l {
	//		l = -1
	//	} else if v == r {
	//		r = -1
	//	}
	//}
	//return max(l, r)
}

//leetcode submit region end(Prohibit modification and deletion)
