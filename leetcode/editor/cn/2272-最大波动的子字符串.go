package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func largestVariance(s string) (ans int) {
	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			if b == a {
				continue
			}
			f0, f1 := 0, math.MinInt
			for _, ch := range s {
				if ch == a {
					f0 = max(f0, 0) + 1
					f1++
				} else if ch == b {
					f1, f0 = max(f0, 0)-1, max(f0, 0)-1
				} // else { f0 = max(f0, 0) } 可以留到 ch 等于 a 或者 b 的时候计算，f1 不变
				ans = max(ans, f1)
			}
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
