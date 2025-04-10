package main

import (
	"math"
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func smallestDifference(a []int, b []int) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	sort.Ints(a)
	sort.Ints(b)
	ans := math.MaxInt
	for i, j := 0, 0; i < len(a) && j < len(b); {
		ans = min(ans, abs(a[i]-b[j]))
		if a[i] == b[j] {
			break
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
