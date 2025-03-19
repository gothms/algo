package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func largestVariance(s string) int {
	// dp

}

//leetcode submit region end(Prohibit modification and deletion)

func largestVariance_(s string) int {
	ans := 0
	for ax := 'a'; ax <= 'z'; ax++ {
		for in := 'a'; in <= 'z'; in++ {
			f0, f1 := 0, math.MinInt
			for _, c := range s {
				switch c {
				case ax:
					f0 = max(f0, 0) + 1
					f1++
				case in:
					f0 = max(f0, 0) - 1
					f1 = f0
				default:
				}
				ans = max(ans, f1)
			}
		}
	}
	return ans
}
