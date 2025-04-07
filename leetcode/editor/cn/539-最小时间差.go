package main

import (
	"sort"
	"strconv"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findMinDifference(timePoints []string) int {
	// lc
	//getMunite := func(t string) int {
	//	return (int(t[0]-'0')*10+int(t[1]-'0'))*60 + int(t[3]-'0')*10 + int(t[4]-'0')
	//}

	const minuteMax = 1440
	n := len(timePoints)
	if n > minuteMax { // 鸽巢原理
		return 0
	}
	sort.Slice(timePoints, func(i, j int) bool {
		return timePoints[i] < timePoints[j]
	})
	d := func(x, y string) int {
		xh, _ := strconv.Atoi(x[:2])
		xm, _ := strconv.Atoi(x[3:])
		yh, _ := strconv.Atoi(y[:2])
		ym, _ := strconv.Atoi(y[3:])
		return (yh-xh)*60 + ym - xm
	}
	ans := d(timePoints[n-1], "24:00") + d("00:00", timePoints[0])
	for i := 1; i < n; i++ {
		ans = min(ans, d(timePoints[i-1], timePoints[i]))
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
