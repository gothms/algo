package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 857 // 247
	digits := rotatedDigits(n)
	fmt.Println(digits)
}

//leetcode submit region begin(Prohibit modification and deletion)

var memo788 []int

func init() {
	gb := [10]int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}
	n := 10000
	memo788 = make([]int, 0, 2320)
out:
	for i := 1; i <= n; i++ {
		good := false
		for v := i; v > 0; v /= 10 {
			switch gb[v%10] {
			case -1:
				continue out
			case 1:
				good = true
			}
		}
		if good {
			memo788 = append(memo788, i)
		}
	}
}
func rotatedDigits(n int) int {
	return sort.SearchInts(memo788, n+1)
}

//leetcode submit region end(Prohibit modification and deletion)
