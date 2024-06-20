package main

import (
	"cmp"
	"fmt"
)

func main() {
	a := []int{28, 28, 4, 18, 36, -8, 39, 4}
	b := []int{-9, 9, 2, -19, -14, 31, 27, 0}
	trend := temperatureTrend(a, b)
	fmt.Println(trend)
}

// leetcode submit region begin(Prohibit modification and deletion)
func temperatureTrend(temperatureA []int, temperatureB []int) int {
	// compare
	ans, cur := 0, 0
	for i := range temperatureA[:len(temperatureA)-1] {
		if cmp.Compare(temperatureA[i], temperatureA[i+1]) == cmp.Compare(temperatureB[i], temperatureB[i+1]) {
            cur++
			ans = max(ans, cur)
		} else {
			cur = 0
		}
	}
	return ans

	// 个人
	//ans, cur, n := 0, 0, len(temperatureA)
	//for i := 1; i < n; i++ {
	//	if da, db := temperatureA[i]-temperatureA[i-1], temperatureB[i]-temperatureB[i-1]; da == db || da*db > 0 {
	//		cur++
	//	} else {
	//		cur = 0
	//	}
	//	ans = max(ans, cur)
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
