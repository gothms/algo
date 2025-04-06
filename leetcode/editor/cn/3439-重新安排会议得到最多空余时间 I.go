package main

import "fmt"

func main() {
	et, k := 34, 2 // 18
	start := []int{0, 17}
	end := []int{14, 19}
	time := maxFreeTime(et, k, start, end)
	fmt.Println(time)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	// 滑动窗体
	startTime = append(startTime, eventTime)
	endTime = append([]int{0}, endTime...)
	ans, sum := 0, 0
	for i, s := range startTime {
		sum += s - endTime[i]
		if i >= k {
			ans = max(ans, sum)
			sum -= startTime[i-k] - endTime[i-k]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
