package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 11
	n = 1000 // 3
	n = 3    // 3
	n = 10   // 1
	digit := findNthDigit(n)
	fmt.Println(digit)
}

// leetcode submit region begin(Prohibit modification and deletion)

//var counter400 [11]int
//
//func init() {
//	maxV, minV := 9, 1
//	for i := 1; i < 11; i++ {
//		counter400[i] = (maxV-minV+1)*i + counter400[i-1]
//		maxV, minV = maxV*10+9, minV*10
//	}
//	fmt.Println(counter400)
//}

// var counter400 = []int{9, 99, 999, 9999, 99999, 999999, 9999999, 99999999, 999999999, 9999999999}
var counter400 = []int{9, 189, 2889, 38889, 488889, 5888889, 68888889, 788888889, 8888888889, 98888888889}

func findNthDigit(n int) int {
	// 个人
	if n < 10 { // fast path
		return n
	}
	i := sort.SearchInts(counter400, n) // 预处理
	d := n - counter400[i-1]
	// x, y := int(math.Pow10(i))+(d-1)/(i+1), i-(d-1)%(i+1)
	x, y := (d-1)/(i+1), i-(d-1)%(i+1)
	for j := 0; j < y; j++ {
		x /= 10
	}
	if y == i {
		return x%10 + 1 // 修正最高位
	}
	return x % 10
}

//leetcode submit region end(Prohibit modification and deletion)
