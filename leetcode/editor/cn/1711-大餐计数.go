package main

import (
	"fmt"
)

func main() {
	fmt.Println(11 &^ 1)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countPairs(deliciousness []int) int {
	return 0

	// 超时
	//const mod = 1e9 + 7
	//ans := 0
	//memo := make(map[int]int)
	//for _, d := range deliciousness {
	//	for k, v := range memo {
	//		if bits.OnesCount(uint(d+k)) == 1 {
	//			ans = (ans + v) % mod
	//		}
	//	}
	//	memo[d]++
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
