package main

import "fmt"

func main() {
	differences := []int{3, -4, 5, 1, -2}
	lower, upper := -4, 5 // 4
	arrays := numberOfArrays(differences, lower, upper)
	fmt.Println(arrays)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfArrays(differences []int, lower int, upper int) int {
	// diff[i] = x, diff[i+1] = y
	// x 区间：[lower-d, upper-d]
	// y 区间：[lower+d, upper+d]
	// 需要考虑连续情况，记录最小和最大连续值
	ans, mn, mx, l, r := upper-lower+1, 0, 0, 0, 0
	for _, d := range differences {
		mn, mx = min(mn+d, d), max(mx+d, d)
		l, r = max(lower, lower-mn, lower-mx), min(upper, upper-mn, upper-mx)
		if l > r {
			return 0
		}
		ans = min(ans, r-l+1)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
