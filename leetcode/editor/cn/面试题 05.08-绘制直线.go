package main

import (
	"fmt"
)

func main() {
	length, w := 3, 96
	x1, x2, y := 0, 95, 0
	//length, w = 15, 96
	//x1, x2, y = 81, 95, 1 // [0,0,0,0,0,32767,0,0,0,0,0,0,0,0,0]
	line := drawLine(length, w, x1, x2, y)
	fmt.Println(line)
}

// leetcode submit region begin(Prohibit modification and deletion)
func drawLine(length int, w int, x1 int, x2 int, y int) []int {
	col := w / 32
	l, r, lv, rv := x1/32+y*col, x2/32+y*col, x1&31, x2&31
	ans := make([]int, length)
	if l == r {
		ans[l] = int(int32(1<<(rv-lv+1)-1) << (31 - rv)) // 正确的 int32 移位
	} else {
		ans[l], ans[r] = int(int32(1<<(31-lv+1))-1), int(int32(1<<(rv+1)-1)<<(31-rv))
		for i := l + 1; i < r; i++ {
			ans[i] = -1
		}
	}
	return ans

	// 个人：问题点，1<<32 时，1 为 int 则 (1<<32) - 1 = 4294967295
	//col := w / 32
	//l, r, lv, rv := x1/32+y*col, x2/32+y*col, x1&31, x2&31
	//v := int32(1)
	//temp := make([]int32, length)
	//if l == r {
	//	temp[l] = (v<<(rv-lv+1) - 1) << (31 - rv)
	//} else {
	//	temp[l], temp[r] = v<<(31-lv+1)-1, (v<<(rv+1)-1)<<(31-rv)
	//	for i := l + 1; i < r; i++ {
	//		temp[i] = -1
	//	}
	//}
	//ans := make([]int, length)
	//for i, val := range temp {
	//	ans[i] = int(val)
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
