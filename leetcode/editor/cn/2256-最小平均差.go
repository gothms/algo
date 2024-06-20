package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{4, 2, 0}
	difference := minimumAverageDifference(nums)
	fmt.Println(difference)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumAverageDifference(nums []int) int {
	sl, sr, n := 0, 0, len(nums)
	if n == 1 {
		return 0
	}
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}
	for _, v := range nums {
		sr += v
	}
	d, ans := math.MaxInt32, 0
	for i, v := range nums[:n-1] {
		sl += v
		sr -= v
		if dif := abs(sl/(i+1) - sr/(n-i-1)); dif < d {
			d, ans = dif, i
		}
	}
	if (sl+nums[n-1])/n < d {
		return n - 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
