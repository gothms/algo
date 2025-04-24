package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	nums := []int{11, 13, 11, 14, 11}
	i := countPairs(nums)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countPairs(nums []int) int {
	mx := slices.Max(nums)
	k := int(math.Log10(float64(mx))) + 1
	memo := make(map[string]int)
	ans := 0
	for _, v := range nums {
		s := fmt.Sprintf("%0*d", k, v)
		ans += memo[s]
		buf := []byte(s)
		for i := range slices.Backward(buf) {
			for j := range buf[:i] {
				if buf[i] != buf[j] { // 交换
					buf[i], buf[j] = buf[j], buf[i]
					ans += memo[string(buf)]
					buf[i], buf[j] = buf[j], buf[i]
				}
			}
		}
		memo[s]++
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
