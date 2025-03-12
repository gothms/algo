package main

import (
	"math"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func sumOfBeauties(nums []int) int {
	n := len(nums)
	l, r := make([]int, n), make([]int, n)
	for i, lv, rv := 0, math.MinInt, math.MaxInt; i < n; i++ {
		lv, rv = max(lv, nums[i]), min(rv, nums[n-i-1])
		l[i], r[n-i-1] = lv, rv
	}
	ans := 0
	for i := 1; i < len(nums)-1; i++ {
		if l[i-1] < nums[i] && nums[i] < r[i+1] {
			ans += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
