package main

import "fmt"

func main() {
	et := 10 // 6
	start := []int{0, 3, 7, 9}
	end := []int{1, 4, 8, 10}
	time := maxFreeTime(et, start, end)
	fmt.Println(time)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
	memo := [4][2]int{{}, {-1, -1}, {-1, -1}, {-1, -1}} // 长度 & 索引
	startTime = append(startTime, eventTime)            // 后面增加 eventTime
	endTime = append([]int{0}, endTime...)              // 前面增加 0
	for i, s := range startTime {                       // 选出前三的空闲时间段
		memo[0] = [2]int{s - endTime[i], i}
		for j := 1; j < 4; j++ {
			if memo[j][0] < memo[j-1][0] {
				memo[j], memo[j-1] = memo[j-1], memo[j]
			}
		}
	}
	ans := 0
	for i, s := range startTime[:len(startTime)-1] {
		l, r := s-endTime[i], startTime[i+1]-endTime[i+1]
		d := endTime[i+1] - s
		ans = max(ans, l+r)      // 左右滑动
		for j := 3; j > 0; j-- { // 调动到前三的空闲时间段的其中之一
			if i != memo[j][1] && i+1 != memo[j][1] && d <= memo[j][0] {
				ans = max(ans, l+r+d)
				break
			}
		}
	}
	return ans

	// lc
	//n := len(startTime)
	//a, b, c := 0, -1, -1
	//get := func(i int) (d int) {
	//	if i == 0 {
	//		d = startTime[0]
	//	} else if i == n {
	//		d = eventTime - endTime[n-1]
	//	} else {
	//		d = startTime[i] - endTime[i-1]
	//	}
	//	return
	//}
	//for i := 1; i <= n; i++ {
	//	d := get(i)
	//	if d > get(a) {
	//		a, b, c = i, a, b
	//	} else if b < 0 || d > get(b) {
	//		b, c = i, b
	//	} else if c < 0 || d > get(c) {
	//		c = i
	//	}
	//}
	//ans := 0
	//for i := 0; i < n; i++ {
	//	size := endTime[i] - startTime[i]
	//	d := get(i) + get(i+1)
	//	ans = max(ans, d)
	//	if i != a && i+1 != a && size <= get(a) ||
	//		i != b && i+1 != b && size <= get(b) ||
	//		size <= get(c) {
	//		ans = max(ans, size+d)
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
